package app

import (
	"context"
	"fmt"
	"time"

	"github.com/marcelo/acc_audio_guard/internal/audio"
	"github.com/marcelo/acc_audio_guard/internal/config"
	"github.com/marcelo/acc_audio_guard/internal/logx"
	"github.com/marcelo/acc_audio_guard/internal/rnnoise"
)

// App coordinates dependencies and lifecycle.
type App struct {
	cfg    config.Config
	logger *logx.Logger
}

// New validates and creates an app instance.
func New(cfg config.Config) (*App, error) {
	if cfg.SampleRate <= 0 {
		return nil, fmt.Errorf("invalid sample-rate: %d", cfg.SampleRate)
	}
	if cfg.FrameMs <= 0 {
		return nil, fmt.Errorf("invalid frame-ms: %d", cfg.FrameMs)
	}
	if config.FrameSizeSamples(cfg.SampleRate, cfg.FrameMs) <= 0 {
		return nil, fmt.Errorf("invalid audio shape sample-rate=%d frame-ms=%d", cfg.SampleRate, cfg.FrameMs)
	}

	return &App{
		cfg:    cfg,
		logger: logx.New("AudioGuard"),
	}, nil
}

// Run starts pipeline execution.
func (a *App) Run(ctx context.Context) error {
	a.logger.Info("startup input=%s output=%s sample_rate=%d frame_ms=%d dry_run=%t duration_sec=%d",
		a.cfg.InputDevice,
		a.cfg.OutputDevice,
		a.cfg.SampleRate,
		a.cfg.FrameMs,
		a.cfg.DryRun,
		a.cfg.DurationSec,
	)

	runCtx := ctx
	cancel := func() {}
	if a.cfg.DurationSec > 0 {
		runCtx, cancel = context.WithTimeout(ctx, time.Duration(a.cfg.DurationSec)*time.Second)
	}
	defer cancel()

	var suppressor audio.NoiseSuppressor
	if !a.cfg.DryRun {
		engine, err := rnnoise.New(a.cfg.SampleRate, a.cfg.FrameMs)
		if err != nil {
			return fmt.Errorf("rnnoise init failed: %w", err)
		}
		suppressor = engine
	}

	pipeline := audio.NewPipeline(a.cfg, suppressor)
	stats, err := pipeline.Run(runCtx)
	if err != nil {
		a.logger.Error("pipeline_stopped err=%v", err)
		return err
	}

	a.logger.Info("runtime_summary frames=%d avg_process=%s max_process=%s total_process=%s",
		stats.FramesProcessed,
		stats.AvgProcess,
		stats.MaxProcess,
		stats.TotalProcess,
	)
	a.logger.Info("shutdown complete")
	return nil
}
