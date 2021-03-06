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
	"testing"

	"github.com/stretchr/testify/assert"
)

// -------------------------------------------------------------------------------- smaller

func Test_bufPool_Alloc_Smaller(t *testing.T) {
	proxy := assert.New(t)

	pool, _ := NewBufPool(64, 1024, 2)

	mem := pool.Alloc(64)
	proxy.Equal(64, len(mem))
	proxy.Equal(64, cap(mem))
	pool.Free(mem)

	mem = pool.Alloc(32)
	proxy.Equal(32, len(mem))
	proxy.Equal(64, cap(mem))
	pool.Free(mem)

	mem = pool.Alloc(128)
	proxy.Equal(128, len(mem))
	proxy.Equal(128, cap(mem))
	pool.Free(mem)
}

// -------------------------------------------------------------------------------- larger

func Test_bufPool_Alloc_Larger(t *testing.T) {
	proxy := assert.New(t)

	pool, _ := NewBufPool(64, 1024, 2)

	mem := pool.Alloc(2048) // 2048 > 1024
	proxy.Equal(2048, len(mem))
	proxy.Equal(2048, cap(mem))
	pool.Free(mem)
}

// -------------------------------------------------------------------------------- buf_pool

func BenchmarkBufPool_Alloc_64(b *testing.B) {
	pool, _ := NewBufPool(64, 1024, 2)

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			pool.Free(pool.Alloc(64))
		}
	})
}

func BenchmarkBufPool_Alloc_128(b *testing.B) {
	pool, _ := NewBufPool(64, 1024, 2)

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			pool.Free(pool.Alloc(128))
		}
	})
}

func BenchmarkBufPool_Alloc_256(b *testing.B) {
	pool, _ := NewBufPool(64, 1024, 2)

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			pool.Free(pool.Alloc(256))
		}
	})
}

func BenchmarkBufPool_Alloc_512(b *testing.B) {
	pool, _ := NewBufPool(64, 1024, 2)

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			pool.Free(pool.Alloc(512))
		}
	})
}
