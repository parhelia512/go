// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sysrand

import "internal/syscall/windows"

// RtlGenRandom only returns 1<<32-1 bytes at a time. We only read at
// most 1<<31-1 bytes at a time so that  this works the same on 32-bit
// and 64-bit systems.
const maxSize = 1<<31-1

func read(b []byte) error {
	for len(b) > 0 {
		size := len(b)
		if size > maxSize {
			size = maxSize
		}
		err := windows.RtlGenRandom(b[:size])
		if err != nil {
			return err
		}
		b = b[size:]
	}
	return nil
}
