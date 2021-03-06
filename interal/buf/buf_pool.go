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
)

type bufPool struct {
	classes     []*BufferedPool
	classesSize sizes
	min         uint
	max         uint
}

func (pool *bufPool) Alloc(expectSize uint) bytez {
	if expectSize <= pool.max {
		for i := 0; i < len(pool.classesSize); i++ {
			if pool.classesSize[i] >= expectSize {
				mem := pool.classes[i].Acquire()
				if len(mem.buf) == 0 {
					// the mem maybe the default value by New BufferedPool
					return make(bytez, expectSize)
				} // make new

				return (mem.buf)[:expectSize]
			} // make new
		}
	} // make new

	return make(bytez, expectSize)
}

func (pool *bufPool) Free(mem bytez) {
	if size := uint(cap(mem)); size <= pool.max {
		for i := 0; i < len(pool.classesSize); i++ {
			if pool.classesSize[i] >= size {
				_ = pool.classes[i].Release(mem)

				return
			}
		}
	}
}

func NewBufPool(min, max, factor uint) (Pool, error) {
	if min%2 != 0 {
		return nil, fmt.Errorf("the min value must be an integer multiple of 2:%d", min)
	}
	if max%2 != 0 {
		return nil, fmt.Errorf("the max value must be an integer multiple of 2:%d", max)
	}
	if max%min != 0 {
		return nil, fmt.Errorf("the max value must be an integer multiple of min:%d", min)
	}

	block := uint(0)
	for chunkSize := min; chunkSize <= max; chunkSize *= factor {
		block++
	}

	pool := &bufPool{
		classes:     make([]*BufferedPool, block),
		classesSize: make(sizes, block),
		min:         min,
		max:         max,
	}

	index := 0
	for chunkSize := min; chunkSize <= max; chunkSize *= factor {
		newPool := NewBufferedPool(block, chunkSize)
		pool.classes[index] = newPool

		pool.classesSize[index] = chunkSize
		buf := NewBufferBySize(chunkSize)

		go func(idx int, buff *Buffer) {
			pool.classes[idx].bufCh <- buff
		}(index, buf)

		index++
	}

	return pool, nil
}
