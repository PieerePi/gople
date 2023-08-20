// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 362.
//
// The version of this program that appeared in the first and second
// printings did not comply with the proposed rules for passing
// pointers between Go and C, described here:
// https://github.com/golang/proposal/blob/master/design/12416-cgo-pointers.md
//
// The rules forbid a C function like bz2compress from storing 'in'
// and 'out' (pointers to variables allocated by Go) into the Go
// variable 's', even temporarily.
//
// The version below, which appears in the third printing, has been
// corrected.  To comply with the rules, the bz_stream variable must
// be allocated by C code.  We have introduced two C functions,
// bz2alloc and bz2free, to allocate and free instances of the
// bz_stream type.  Also, we have changed bz2compress so that before
// it returns, it clears the fields of the bz_stream that contain
// pointers to Go variables.

//!+

// Package bzip provides a writer that uses bzip2 compression (bzip.org).
package bzip

import (
	"io"
	"os/exec"
	"sync"
)

type writer struct {
	w     io.Writer // underlying output stream
	cmd   *exec.Cmd
	stdin io.WriteCloser
	sync.Mutex
}

// NewWriter returns a writer for bzip2-compressed streams.
func NewWriter(out io.Writer) (io.WriteCloser, error) {
	var err error
	w := &writer{w: out}
	w.cmd = exec.Command("/bin/bzip2", "-9")
	w.cmd.Stdout = out
	if w.stdin, err = w.cmd.StdinPipe(); err != nil {
		return nil, err
	}
	if err = w.cmd.Start(); err != nil {
		return nil, err
	}
	return w, nil
}

//!-

// !+write
func (w *writer) Write(data []byte) (int, error) {
	w.Lock()
	defer w.Unlock()
	if w.cmd == nil {
		panic("closed")
	}
	return w.stdin.Write(data)
}

//!-write

// !+close
// Close flushes the compressed data and closes the stream.
// It does not close the underlying io.Writer.
func (w *writer) Close() error {
	w.Lock()
	defer w.Unlock()
	if w.cmd == nil {
		panic("closed")
	}
	pipeErr := w.stdin.Close()
	cmdErr := w.cmd.Wait()
	w.cmd = nil
	if pipeErr != nil {
		return pipeErr
	}
	if cmdErr != nil {
		return cmdErr
	}
	return nil
}

//!-close
