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
)

// -------------------------------------------------------------------------------- make

func BenchmarkMake_Alloc_64(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		var mem []byte
		for pb.Next() {
			// Alloc
			mem = make([]byte, 64)
		}

		// Free
		mem = mem[:0]
	})
}

func BenchmarkMake_Alloc_128(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		var mem []byte
		for pb.Next() {
			// Alloc
			mem = make([]byte, 128)
		}

		// Free
		mem = mem[:0]
	})
}

func BenchmarkMake_Alloc_256(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		var mem []byte
		for pb.Next() {
			// Alloc
			mem = make([]byte, 256)
		}

		// Free
		mem = mem[:0]
	})
}

func BenchmarkMake_Alloc_512(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		var mem []byte
		for pb.Next() {
			// Alloc
			mem = make([]byte, 512)
		}

		// Free
		mem = mem[:0]
	})
}
