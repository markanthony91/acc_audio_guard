package audio

import (
	"context"
	"errors"
	"time"

	"github.com/marcelo/acc_audio_guard/internal/config"
)

// NoiseSuppressor is the minimal contract for the RNNoise wrapper.
type NoiseSuppressor interface {
	ProcessFrame(frame []float32) ([]float32, error)
	Close() error
}

// RuntimeStats summarizes runtime processing behavior.
type RuntimeStats struct {
	FramesProcessed int64
	TotalProcess    time.Duration
	AvgProcess      time.Duration
	MaxProcess      time.Duration
}

// Pipeline coordinates capture, suppression and output paths.
type Pipeline struct {
	cfg        config.Config
	suppressor NoiseSuppressor
	source     FrameSource
}

// NewPipeline creates a processing pipeline. Suppressor may be nil in dry-run mode.
func NewPipeline(cfg config.Config, suppressor NoiseSuppressor) *Pipeline {
	return &Pipeline{
		cfg:        cfg,
		suppressor: suppressor,
		source:     NewSyntheticSource(cfg.SampleRate),
	}
}

// Run starts the processing loop until context cancellation.
func (p *Pipeline) Run(ctx context.Context) (RuntimeStats, error) {
	if p.cfg.FrameMs <= 0 {
		return RuntimeStats{}, errors.New("frame-ms must be greater than zero")
	}

	frameSize := config.FrameSizeSamples(p.cfg.SampleRate, p.cfg.FrameMs)
	if frameSize <= 0 {
		return RuntimeStats{}, errors.New("invalid frame size calculated")
	}

	ticker := time.NewTicker(time.Duration(p.cfg.FrameMs) * time.Millisecond)
	defer ticker.Stop()

	stats := RuntimeStats{}

	for {
		select {
		case <-ctx.Done():
			if p.suppressor != nil {
				_ = p.suppressor.Close()
			}
			if stats.FramesProcessed > 0 {
				stats.AvgProcess = time.Duration(int64(stats.TotalProcess) / stats.FramesProcessed)
			}
			return stats, nil
		case <-ticker.C:
			frame, err := p.source.NextFrame(frameSize)
			if err != nil {
				return stats, err
			}

			start := time.Now()
			if p.suppressor != nil {
				frame, err = p.suppressor.ProcessFrame(frame)
				if err != nil {
					return stats, err
				}
			}

			_ = frame
			elapsed := time.Since(start)
			stats.FramesProcessed++
			stats.TotalProcess += elapsed
			if elapsed > stats.MaxProcess {
				stats.MaxProcess = elapsed
			}
		}
	}
}
