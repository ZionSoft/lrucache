// Copyright (c) 2015 ZionSoft. All rights reserved.
// Use of this source code is governed by a BSD-style license
// that can be found in the LICENSE file.

package lrucache

import (
	"testing"
)

type cacheValue struct {
	size uint64
}

func (cv *cacheValue) Size() uint64 {
	return cv.size
}

func TestInitialState(t *testing.T) {
	var capacity uint64 = 100
	var size uint64 = 0

	c := New(capacity)
	if c.Capacity() != capacity {
		t.Errorf("capacity = %v, want %v", c.Capacity(), capacity)
	}
	if c.Size() != size {
		t.Errorf("size = %v, want %v", c.Size(), size)
	}
}

func TestSimpleCache(t *testing.T) {
	var size uint64 = 20
	key := "key"
	value := &cacheValue{size}

	c := New(100)
	if v := c.Get(key); v != nil {
		t.Errorf("LRUCache has incorrect value for key(%v): %v", key, v)
	}

	c.Set(key, value)
	if c.Size() != size {
		t.Errorf("size = %v, want %v", c.Size(), size)
	}
	if v := c.Get(key); v.(*cacheValue) != value {
		t.Errorf("%v = %v, want %v", key, v, value)
	}

	var size2 uint64 = 40
	value2 := &cacheValue{size2}
	c.Set(key, value2)
	if c.Size() != size2 {
		t.Errorf("size = %v, want %v", c.Size(), size2)
	}
	if v := c.Get(key); v.(*cacheValue) != value2 {
		t.Errorf("%v = %v, want %v", key, v, value2)
	}
}

func TestCapacity(t *testing.T) {
	var capacity uint64 = 3
	c := New(capacity)
	value := &cacheValue{1}

	c.Set("key1", value)
	c.Set("key2", value)
	c.Set("key3", value)
	if c.Size() != capacity {
		t.Errorf("size = %v, want %v", c.Size(), capacity)
	}

	c.Set("key4", value)
	if c.Size() != capacity {
		t.Errorf("size = %v, want %v", c.Size(), capacity)
	}
	if v := c.Get("key1"); v != nil {
		t.Errorf("key1 is not evicted")
	}
}

func TestDelete(t *testing.T) {
	var size uint64 = 20
	key := "key"
	value := &cacheValue{size}

	c := New(100)
	c.Set(key, value)
	if v := c.Get(key); v.(*cacheValue) != value {
		t.Errorf("%v = %v, want %v", key, v, value)
	}

	c.Delete("key2")
	if v := c.Get(key); v.(*cacheValue) != value {
		t.Errorf("%v = %v, want %v", key, v, value)
	}
	if c.Size() != size {
		t.Errorf("size = %v, want %v", c.Size(), size)
	}

	c.Delete(key)
	if v := c.Get(key); v != nil {
		t.Errorf("failed to delete %v", key)
	}
	if c.Size() != 0 {
		t.Errorf("size = %v, want %v", c.Size(), 0)
	}
}

func TestClear(t *testing.T) {
	key := "key"
	value := &cacheValue{20}

	c := New(100)
	c.Set(key, value)
	if v := c.Get(key); v.(*cacheValue) != value {
		t.Errorf("%v = %v, want %v", key, v, value)
	}

	c.Clear()
	if v := c.Get(key); v != nil {
		t.Errorf("failed to delete %v", key)
	}
	if c.Size() != 0 {
		t.Errorf("size = %v, want %v", c.Size(), 0)
	}
}
