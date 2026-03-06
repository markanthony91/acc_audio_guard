package device

import "errors"

// ErrNotImplemented marks placeholder hardware integration points.
var ErrNotImplemented = errors.New("device integration not implemented")

// InputDevice is a logical representation of a capture device.
type InputDevice struct {
	Name string
}

// OutputDevice is a logical representation of a playback/virtual cable device.
type OutputDevice struct {
	Name string
}

// DiscoverInputDevices should return available input devices from audio backend.
func DiscoverInputDevices() ([]InputDevice, error) {
	return nil, ErrNotImplemented
}

// DiscoverOutputDevices should return available output devices from audio backend.
func DiscoverOutputDevices() ([]OutputDevice, error) {
	return nil, ErrNotImplemented
}
