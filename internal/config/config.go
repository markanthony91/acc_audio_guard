package config

import "flag"

const (
	DefaultSampleRate = 48000
	DefaultFrameMs    = 10
)

// Config contains runtime options for the ORFEU processing loop.
type Config struct {
	InputDevice       string
	OutputDevice      string
	SampleRate        int
	FrameMs           int
	DryRun            bool
	DurationSec       int
	ReportEveryFrames int
}

// FromFlags parses command line flags into a Config.
func FromFlags() Config {
	cfg := Config{}

	flag.StringVar(&cfg.InputDevice, "input", "default", "input audio device name")
	flag.StringVar(&cfg.OutputDevice, "output", "default", "output audio device name")
	flag.IntVar(&cfg.SampleRate, "sample-rate", DefaultSampleRate, "audio sample rate in Hz")
	flag.IntVar(&cfg.FrameMs, "frame-ms", DefaultFrameMs, "frame size in milliseconds")
	flag.BoolVar(&cfg.DryRun, "dry-run", true, "run processing loop without real RNNoise I/O")
	flag.IntVar(&cfg.DurationSec, "duration-sec", 0, "run duration in seconds (0 keeps running until signal)")
	flag.IntVar(&cfg.ReportEveryFrames, "report-every-frames", 100, "emit progress report every N frames")

	flag.Parse()
	return cfg
}

// FrameSizeSamples calculates samples per frame for the selected sample rate and frame duration.
func FrameSizeSamples(sampleRate, frameMs int) int {
	if sampleRate <= 0 || frameMs <= 0 {
		return 0
	}
	return (sampleRate * frameMs) / 1000
}
