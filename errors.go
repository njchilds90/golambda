package golambda

import "errors"

// ErrInvalidChunkSize is returned when Chunk is called with size < 1.
var ErrInvalidChunkSize = errors.New("golambda: chunk size must be >= 1")
