package bzip

import (
	"C"
	"io"
)

type writer struct {
	w io.Writer
	stream *C.bz_stream
	outbuf [64 * 1024]byte
}

func NewWriter(out io.Writer) io.WriteCloser {
	const blockSize = 9
	const verbosity = 0
	const workFactor = 30
	w := &writer{w: out, stream: C.bz2alloc()}
}
