// Copyright (c) 2015 ZionSoft. All rights reserved.
// Use of this source code is governed by a BSD-style license
// that can be found in the LICENSE file.

// A simple LRU (least recently used) cache.
package lrucache

import (
	"container/list"
	"sync"
)

// Value is the interface that a type must satisfy to work with the cache.
type Value interface {
	Size() uint64
}

type item struct {
	key   string
	value Value
}

// A LRUCache is a thread-safe LRU cache that holds a limited number of values.
type LRUCache struct {
	mu       sync.RWMutex
	items    map[string]*list.Element
	list     *list.List
	capacity uint64
	size     uint64
}

// New returns an initialized LRUCache with the specified capacity in bytes.
func New(capacity uint64) *LRUCache {
	return &LRUCache{
		items:    make(map[string]*list.Element),
		list:     list.New(),
		capacity: capacity,
	}
}

// Get returns the value for key or nil.
func (lru *LRUCache) Get(key string) Value {
	lru.mu.RLock()
	defer lru.mu.RUnlock()

	e := lru.items[key]
	if e == nil {
		return nil
	}
	lru.list.MoveToFront(e)
	return e.Value.(*item).value
}

// Set caches value for key.
func (lru *LRUCache) Set(key string, value Value) {
	lru.mu.Lock()
	defer lru.mu.Unlock()

	if e := lru.items[key]; e != nil {
		lru.list.MoveToFront(e)
		lru.size += (value.Size() - e.Value.(*item).value.Size())
		e.Value.(*item).value = value
	} else {
		item := &item{
			key:   key,
			value: value,
		}
		e := lru.list.PushFront(item)
		lru.items[key] = e
		lru.size += value.Size()
	}

	for lru.size > lru.capacity {
		e := lru.list.Back()
		delete(lru.items, e.Value.(*item).key)
		lru.list.Remove(e)
		lru.size -= e.Value.(*item).value.Size()
	}
}

// Delete removes the value for key.
func (lru *LRUCache) Delete(key string) {
	lru.mu.Lock()
	defer lru.mu.Unlock()

	e := lru.items[key]
	if e == nil {
		return
	}

	delete(lru.items, key)
	lru.list.Remove(e)
	lru.size -= e.Value.(*item).value.Size()
}

// Clear removes all values.
func (lru *LRUCache) Clear() {
	lru.mu.Lock()
	defer lru.mu.Unlock()

	lru.items = make(map[string]*list.Element)
	lru.list.Init()
	lru.size = 0
}

// Capacity returns the capacity in bytes for the cache.
func (lru *LRUCache) Capacity() uint64 {
	lru.mu.RLock()
	defer lru.mu.RUnlock()

	return lru.capacity
}

// Size returns the current size in bytes for the cache.
func (lru *LRUCache) Size() uint64 {
	lru.mu.RLock()
	defer lru.mu.RUnlock()

	return lru.size
}
