package config

import "testing"

func TestDefaultConstants(t *testing.T) {
	if DefaultSampleRate != 48000 {
		t.Fatalf("unexpected default sample rate: %d", DefaultSampleRate)
	}

	if DefaultFrameMs != 10 {
		t.Fatalf("unexpected default frame ms: %d", DefaultFrameMs)
	}
}

func TestFrameSizeSamples(t *testing.T) {
	if got := FrameSizeSamples(48000, 10); got != 480 {
		t.Fatalf("unexpected frame size: got=%d want=480", got)
	}

	if got := FrameSizeSamples(0, 10); got != 0 {
		t.Fatalf("unexpected frame size for invalid input: %d", got)
	}
}
