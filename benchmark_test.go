// Copyright (c) 2015 ZionSoft. All rights reserved.
// Use of this source code is governed by a BSD-style license
// that can be found in the LICENSE file.

package lrucache

import (
	"strconv"
	"testing"
)

type bytes []byte

func (b bytes) Size() uint64 {
	return uint64(cap(b))
}

func BenchmarkGet(b *testing.B) {
	c := New(64 * 1024 * 1024)
	value := make(bytes, 1024)

	for i := 0; i < 1024; i++ {
		c.Set(strconv.Itoa(i), value)
	}

	for i := 0; i < b.N; i++ {
		v, _ := c.Get("512")
		if v == nil {
			panic("error")
		}
		_ = v
	}
}
