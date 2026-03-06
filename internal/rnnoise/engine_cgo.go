//go:build cgo && rnnoise

package rnnoise

/*
#cgo pkg-config: rnnoise
#include <stdlib.h>
#include <rnnoise.h>
*/
import "C"

import (
	"fmt"
	"sync"
	"unsafe"
)

type engine struct {
	mu     sync.Mutex
	state  *C.DenoiseState
	closed bool
}

// New creates an RNNoise engine instance.
func New(sampleRate, frameMs int) (Engine, error) {
	if sampleRate != RequiredSampleRate || frameMs != RequiredFrameMs {
		return nil, ErrInvalidAudioShape
	}

	state := C.rnnoise_create(nil)
	if state == nil {
		return nil, fmt.Errorf("failed to create rnnoise state")
	}

	return &engine{state: state}, nil
}

func (e *engine) ProcessFrame(frame []float32) ([]float32, error) {
	if len(frame) != RequiredFrameSize {
		return nil, fmt.Errorf("invalid frame size: got=%d want=%d", len(frame), RequiredFrameSize)
	}

	e.mu.Lock()
	defer e.mu.Unlock()

	if e.closed {
		return nil, fmt.Errorf("rnnoise engine already closed")
	}

	out := make([]float32, RequiredFrameSize)
	C.rnnoise_process_frame(
		e.state,
		(*C.float)(unsafe.Pointer(&out[0])),
		(*C.float)(unsafe.Pointer(&frame[0])),
	)

	return out, nil
}

func (e *engine) Close() error {
	e.mu.Lock()
	defer e.mu.Unlock()

	if e.closed {
		return nil
	}

	C.rnnoise_destroy(e.state)
	e.closed = true
	e.state = nil
	return nil
}
