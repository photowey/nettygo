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

type bufPool struct {
	classes     []BufferedPool
	classesSize sizes
	min         uint
	max         uint
}

func (pool *bufPool) Alloc(expectSize uint) []byte {
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
				_ = pool.classes[i].Release(NewBufferByBuf(mem))

				return
			}
		}
	}
}

func NewBufPool(min, max, factor uint) Pool {
	block := 0
	for chunkSize := min; chunkSize <= max; chunkSize *= factor {
		block++
	}

	pool := &bufPool{
		make([]BufferedPool, block),
		make(sizes, block),
		min,
		max,
	}

	block = 0
	for chunkSize := min; chunkSize <= max; chunkSize *= factor {
		pool.classesSize[block] = chunkSize
		buf := NewBufferBySize(chunkSize)
		pool.classes[block].bufCh <- buf
		block++
	}

	return pool
}
