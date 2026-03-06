package audio

import "math"

// FrameSource provides audio frames to the pipeline.
type FrameSource interface {
	NextFrame(frameSize int) ([]float32, error)
}

// SyntheticSource simulates a microphone stream for local pipeline validation.
type SyntheticSource struct {
	sampleRate int
	phase      float64
}

// NewSyntheticSource creates a deterministic synthetic source.
func NewSyntheticSource(sampleRate int) *SyntheticSource {
	return &SyntheticSource{sampleRate: sampleRate}
}

// NextFrame returns a frame with a voice-like tone plus low-level noise.
func (s *SyntheticSource) NextFrame(frameSize int) ([]float32, error) {
	frame := make([]float32, frameSize)
	step := 2 * math.Pi * 220.0 / float64(s.sampleRate)

	for i := 0; i < frameSize; i++ {
		tone := 0.2 * math.Sin(s.phase)
		noise := 0.02 * math.Sin(3*s.phase)
		frame[i] = float32(tone + noise)
		s.phase += step
	}

	return frame, nil
}
