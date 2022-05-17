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

func Test_bufPool_Alloc(t *testing.T) {
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
