/*
 * Copyright © 2022 photowey (photowey@gmail.com)
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package buf

import (
	"fmt"
	"sync"
)

const (
	DefaultMinBufferSize = 256 * Byte
	DefaultMaxBufferSize = 64 * MB
	DefaultFactor        = 2
)

var (
	_     Pool = (*syncPool)(nil)
	_lock sync.Mutex

	_defaultPool Pool
)

type (
	syncPools = []sync.Pool
	sizes     = []uint
)

// syncPool - wrap of sync.Pool
//
// runtime.GC
type syncPool struct {
	classes     syncPools // slab class
	classesSize sizes
	min         uint
	max         uint
}

func (pool *syncPool) Alloc(expectSize uint) []byte {
	if expectSize <= pool.max {
		for i := 0; i < len(pool.classesSize); i++ {
			if pool.classesSize[i] >= expectSize {
				mem := pool.classes[i].Get().(*[]byte)

				return (*mem)[:expectSize]
			}
		}
	}

	return make(bytez, expectSize) // direct alloc
}

func (pool *syncPool) Free(mem bytez) {
	if size := uint(cap(mem)); size <= pool.max {
		for i := 0; i < len(pool.classesSize); i++ {
			if pool.classesSize[i] >= size {
				pool.classes[i].Put(&mem)

				return
			}
		}
	}
}

// ---------------------------------------------------------------- export

func NewSyncPool(min, max, factor uint) (Pool, error) {
	if min%2 != 0 {
		return nil, fmt.Errorf("the min value must be an integer multiple of 2:%d", min)
	}
	if max%2 != 0 {
		return nil, fmt.Errorf("the max value must be an integer multiple of 2:%d", max)
	}
	if max%min != 0 {
		return nil, fmt.Errorf("the max value must be an integer multiple of min:%d", min)
	}

	block := 0
	for chunkSize := min; chunkSize <= max; chunkSize *= factor {
		block++
	}

	pool := &syncPool{
		make(syncPools, block),
		make(sizes, block),
		min,
		max,
	}

	block = 0
	for chunkSize := min; chunkSize <= max; chunkSize *= factor {
		pool.classesSize[block] = chunkSize
		pool.classes[block].New = func(size uint) func() any {
			return func() any {
				buf := make(bytez, size)

				return &buf
			}
		}(chunkSize)
		block++
	}

	return pool, nil
}

func GetDefaultPoolInstance() Pool {
	if _defaultPool == nil {
		_lock.Lock()
		defer _lock.Unlock()
		if _defaultPool == nil {
			_defaultPool = newDefaultPool()
		}
	}

	return _defaultPool
}

// ---------------------------------------------------------------- private

func newDefaultPool() Pool {
	// 256B-64MB-2
	pool, _ := NewSyncPool(DefaultMinBufferSize, DefaultMaxBufferSize, DefaultFactor)

	return pool
}
