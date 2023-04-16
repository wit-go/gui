package main

import (
	"bytes"
	"io"
)

type FakeFile struct {
	buffer bytes.Buffer
	offset int64
}

func (f *FakeFile) Read(p []byte) (n int, err error) {
	n, err = f.buffer.ReadAt(p, f.offset)
	f.offset += int64(n)
	return n, err
}

func (f *FakeFile) Write(p []byte) (n int, err error) {
	n, err = f.buffer.WriteAt(p, f.offset)
	f.offset += int64(n)
	return n, err
}

func (f *FakeFile) Seek(offset int64, whence int) (int64, error) {
	newOffset := f.offset

	switch whence {
	case io.SeekStart:
		newOffset = offset
	case io.SeekCurrent:
		newOffset += offset
	case io.SeekEnd:
		newOffset = int64(f.buffer.Len()) + offset
	default:
		return 0, io.ErrInvalidWhence
	}

	if newOffset < 0 {
		return 0, io.ErrInvalidWhence
	}

	f.offset = newOffset
	return f.offset, nil
}
