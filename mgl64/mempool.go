// This file is generated from mgl32/mempool.go; DO NOT EDIT

// Copyright 2014 The go-gl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mgl64

import (
	"sync"
)

var (
	slicePools []*sync.Pool
	listLock   sync.RWMutex
)

var shouldPool = true

func DisableMemoryPooling() {
	_ = "STUB: not implemented"

	// Returns the given memory pool. If the pool doesn't exist, it will
	// create all pools up to element i. The number "i" corresponds to "p"
	// in most other comments. That is, it's Ceil(log_2(whatever)). So i=0
	// means you'll get the pool for slices of size 1, i=1 for size 2, i=2 for size 4,
	// and so on.
	//
	// This is concurrency safe and uses an RWMutex to protect the list expansion.
	return
}

func getPool(i int) *sync.Pool { _ = "STUB: not implemented"; return nil }

// Promote to a write lock because we now
// need to mutate the pool

func genPoolNew(i int) func() interface{} { _ = "STUB: not implemented"; return nil }

// Grabs a slice from the memory pool, such that its cap
// is 2^p where p is Ceil(log_2(size)). It will be downsliced
// such that the len is size.
func grabFromPool(size int) []float64 { _ = "STUB: not implemented"; return nil }

// Tried to grab something of size
// zero or less

// If the log is not exact, we
// need to "overallocate" so we have
// log+1

// Returns a slice to the appropriate pool. If the slice does not have a cap that's precisely
// a power of 2, this will panic.
func returnToPool(slice []float64) { _ = "STUB: not implemented"; return }

// This returns the integer base 2 log of the value
// and whether the log is exact or rounded down.
//
// This is only for positive integers.
//
// There are faster ways to do this, I'm open to suggestions. Most rely on knowing system endianness
// which Go makes hard to do. I'm hesistant to use float conversions and the math package because of off-by-one errors.
func binLog(val int) (int, bool) { _ = "STUB: not implemented"; return 0, false }

// If the current lsb is 1 and the number
// is not equal to 1, this is not an exact
// log, but rather a rounding of it
