package audio

import (
	"context"
	"strings"
	"testing"
	"time"

	"github.com/marcelo/acc_audio_guard/internal/config"
)

func TestPipelineRejectsInvalidFrameMs(t *testing.T) {
	cfg := config.Config{SampleRate: 48000, FrameMs: 0}
	p := NewPipeline(cfg, nil)

	_, err := p.Run(context.Background())
	if err == nil {
		t.Fatal("expected error for invalid frame-ms")
	}

	if !strings.Contains(err.Error(), "frame-ms") {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestPipelineStopsOnContextCancel(t *testing.T) {
	cfg := config.Config{SampleRate: 48000, FrameMs: 10}
	p := NewPipeline(cfg, nil)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	done := make(chan error, 1)
	go func() {
		_, err := p.Run(ctx)
		done <- err
	}()

	time.Sleep(20 * time.Millisecond)
	cancel()

	select {
	case err := <-done:
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	case <-time.After(time.Second):
		t.Fatal("pipeline did not stop after cancel")
	}
}
