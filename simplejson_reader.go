package simplejson

import (
	"io"
)

// Implements the json.Unmarshaler interface.
func (j *Json) UnmarshalJSON(p []byte) error { _ = "STUB: not implemented"; return nil }

// NewFromReader returns a *Json by decoding from an io.Reader
func NewFromReader(r io.Reader) (*Json, error) { _ = "STUB: not implemented"; return nil, nil }

// Float64 coerces into a float64
func (j *Json) Float64() (float64, error) { _ = "STUB: not implemented"; return 0, nil }

// Int coerces into an int
func (j *Json) Int() (int, error) { _ = "STUB: not implemented"; return 0, nil }

// Int64 coerces into an int64
func (j *Json) Int64() (int64, error) { _ = "STUB: not implemented"; return 0, nil }

// Uint64 coerces into an uint64
func (j *Json) Uint64() (uint64, error) { _ = "STUB: not implemented"; return 0, nil }
