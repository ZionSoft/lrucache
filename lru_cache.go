// Copyright (c) 2015 ZionSoft. All rights reserved.
// Use of this source code is governed by a BSD-style license
// that can be found in the LICENSE file.

// A simple LRU (least recently used) cache.
package lrucache

import (
	"sync"
)

// Value is the interface that a type must satisfy to work with the cache.
type Value interface {
	Size() uint64
}

// A LRUCache is a thread-safe LRU cache that holds a limited number of values.
type LRUCache struct {
	mu       sync.RWMutex
	capacity uint64
	size     uint64
}

// New returns an initialized LRUCache with the specified capacity in bytes.
func New(capacity uint64) *LRUCache {
	return &LRUCache{
		capacity: capacity,
	}
}

// Get returns the value for key or nil.
func (lru *LRUCache) Get(key string) Value {
	lru.mu.RLock()
	defer lru.mu.RUnlock()

	return nil
}

// Set caches value for key.
func (lru *LRUCache) Set(key string, value Value) {
	lru.mu.Lock()
	defer lru.mu.Unlock()
}

// Delete removes the value for key.
func (lru *LRUCache) Delete(key string) {
	lru.mu.Lock()
	defer lru.mu.Unlock()
}

// Clear removes all values.
func (lru *LRUCache) Clear() {
	lru.mu.Lock()
	defer lru.mu.Unlock()
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
