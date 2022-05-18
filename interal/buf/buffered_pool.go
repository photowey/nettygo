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

type Buffer struct {
	buf []byte
}

type BufferedPool struct {
	bufCh chan *Buffer
}

func (pool *BufferedPool) Acquire() *Buffer {
	select {
	case rvt := <-pool.bufCh:

		return rvt
	}
}

func (pool *BufferedPool) Release(mem []byte) error {
	buf := NewBufferByBuf(mem)
	select {
	case pool.bufCh <- buf:
		return nil
	default:
		// return errors.New("buffered channel overflow")
		// overflow -> discard this mem

		buf = nil // Help GC?

		return nil
	}
}

func NewBuffer() *Buffer {
	return &Buffer{
		buf: make([]byte, 0),
	}
}

func NewBufferBySize(size uint) *Buffer {
	return &Buffer{
		buf: make([]byte, size),
	}
}

func NewBufferByBuf(buf []byte) *Buffer {
	return &Buffer{
		buf: buf,
	}
}

func NewBufferedPool(chSize uint, bufSize uint) *BufferedPool {
	pool := &BufferedPool{
		bufCh: make(chan *Buffer, chSize),
	}
	for i := uint(0); i < chSize; i++ {
		pool.bufCh <- NewBufferBySize(bufSize)
	}

	return pool
}
