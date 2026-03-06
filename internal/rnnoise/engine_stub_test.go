//go:build !rnnoise

package rnnoise

import "testing"

func TestNewStubValidShapeReturnsBuildTagError(t *testing.T) {
	_, err := New(RequiredSampleRate, RequiredFrameMs)
	if err == nil {
		t.Fatal("expected error")
	}
	if err != ErrBuildTagDisabled {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestNewStubInvalidShape(t *testing.T) {
	_, err := New(16000, 20)
	if err == nil {
		t.Fatal("expected error")
	}
	if err != ErrInvalidAudioShape {
		t.Fatalf("unexpected error: %v", err)
	}
}
