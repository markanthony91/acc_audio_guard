//go:build !rnnoise

package rnnoise

// New creates an RNNoise engine instance.
func New(sampleRate, frameMs int) (Engine, error) {
	if sampleRate != RequiredSampleRate || frameMs != RequiredFrameMs {
		return nil, ErrInvalidAudioShape
	}
	return nil, ErrBuildTagDisabled
}
