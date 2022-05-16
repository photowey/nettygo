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
	"runtime"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSyncPool(t *testing.T) {
	proxy := assert.New(t)

	pool := &sync.Pool{
		New: func() interface{} {
			t.Log("create a new object")
			return 1024
		},
	}
	v := pool.Get().(int)
	proxy.Equal(1024, v)

	pool.Put(9527)
	v = pool.Get().(int)
	proxy.Equal(9527, v)

	runtime.GC()

	v = pool.Get().(int)
	proxy.Equal(1024, v)
}

func Test_syncPool_Alloc_Smaller(t *testing.T) {
	proxy := assert.New(t)

	pool, _ := NewSyncPool(64, 1024, 2)

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

func Test_syncPool_Alloc_Smaller_GC(t *testing.T) {
	proxy := assert.New(t)

	pool, _ := NewSyncPool(64, 1024, 2)

	mem1 := pool.Alloc(32)
	proxy.Equal(32, len(mem1))
	proxy.Equal(64, cap(mem1))

	mem2 := pool.Alloc(32)
	proxy.Equal(32, len(mem2))
	proxy.Equal(64, cap(mem2))

	runtime.GC()

	mem3 := pool.Alloc(64)
	proxy.Equal(64, len(mem3))
	proxy.Equal(64, cap(mem3))

	pool.Free(mem1)
	pool.Free(mem2)
	pool.Free(mem3)
}

func Test_syncPool_Alloc_Larger(t *testing.T) {
	proxy := assert.New(t)

	pool, _ := NewSyncPool(64, 1024, 2)

	mem := pool.Alloc(2048) // 2048 > 1024
	proxy.Equal(2048, len(mem))
	proxy.Equal(2048, cap(mem))
	pool.Free(mem)
}

func BenchmarkSyncPool_Alloc_64(b *testing.B) {
	pool, _ := NewSyncPool(64, 1024, 2)

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			pool.Free(pool.Alloc(64))
		}
	})
}

func BenchmarkSyncPool_Alloc_128(b *testing.B) {
	pool, _ := NewSyncPool(64, 1024, 2)

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			pool.Free(pool.Alloc(128))
		}
	})
}

func BenchmarkSyncPool_Alloc_256(b *testing.B) {
	pool, _ := NewSyncPool(64, 1024, 2)

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			pool.Free(pool.Alloc(256))
		}
	})
}

func BenchmarkSyncPool_Alloc_512(b *testing.B) {
	pool, _ := NewSyncPool(64, 1024, 2)

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			pool.Free(pool.Alloc(512))
		}
	})
}
