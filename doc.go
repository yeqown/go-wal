// Package wal implements a write-ahead log for storing data that needs to be
// persisted to disk. The WAL is append-only, and is safe for concurrent
// access. The WAL is optimized for writing, and is not optimized for reading.
//
// Similar to the https://github.com/tidwall/wal package, but with a few differences.
package wal

import "errors"

var (
	// ErrSegmentNotFound is returned when a segment is not found.
	ErrSegmentNotFound = errors.New("segment not found")

	// ErrSegmentCorrupted is returned when a segment is corrupted.
	ErrSegmentCorrupted = errors.New("segment corrupted")
)