package rnnoise

import "errors"

const (
	// RequiredSampleRate is the sample rate expected by RNNoise in its default model.
	RequiredSampleRate = 48000
	// RequiredFrameMs is the frame duration expected by RNNoise.
	RequiredFrameMs = 10
	// RequiredFrameSize is the number of float32 samples per frame for RNNoise.
	RequiredFrameSize = 480
)

var (
	ErrBuildTagDisabled  = errors.New("rnnoise integration disabled: build with -tags rnnoise")
	ErrInvalidAudioShape = errors.New("rnnoise requires sample-rate=48000 and frame-ms=10")
)

// Engine defines operations expected from the RNNoise wrapper.
type Engine interface {
	ProcessFrame(frame []float32) ([]float32, error)
	Close() error
}
