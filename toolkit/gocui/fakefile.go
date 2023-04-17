package main

import (
	"bytes"
	"io"
	"errors"
)

type FakeFile struct {
	reader *bytes.Reader
	buffer *bytes.Buffer
	offset int64
}

func (f *FakeFile) Read(p []byte) (n int, err error) {
	n, err = f.reader.ReadAt(p, f.offset)
	f.offset += int64(n)
	return n, err
}

func (f *FakeFile) Write(p []byte) (n int, err error) {
	n, err = f.buffer.Write(p)
	f.offset += int64(n)
	f.reader.Reset(f.buffer.Bytes())
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
		return 0, errors.New("Seek: whence not at start,current or end")
	}
	// never can get here right?

	if newOffset < 0 {
		return 0, errors.New("Seek: offset < 0")
	}

	f.offset = newOffset
	return f.offset, nil
}

func NewFakeFile() *FakeFile {
	buf := &bytes.Buffer{}
	return &FakeFile{
		reader: bytes.NewReader(buf.Bytes()),
		buffer: buf,
		offset: 0,
	}
}
