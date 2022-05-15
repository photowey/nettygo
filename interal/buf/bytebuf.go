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

var _ ByteBuf = (*bytebuf)(nil)

type bytez = []byte

// ByteBuf a buf with byte arrays
/*
 * <pre>
 *      +-------------------+------------------+------------------+
 *      | discardable bytes |  readable bytes  |  writable bytes  |
 *      |                   |     (CONTENT)    |                  |
 *      +-------------------+------------------+------------------+
 *      |                   |                  |                  |
 *      0      <=      readerIndex   <=   writerIndex    <=    capacity
 * </pre>
 */
type ByteBuf interface {
	Read(bytes []byte) (n int, err error)
	ReadInt() (int, error)

	Write(bytes []byte) (int, error)
	WriteInt(v int) (n int, err error)

	Capacity() uint
	// Readable Returns the number of readable bytes which is equal to
	//
	// buf.writerIndex - buf.readerIndex
	Readable() uint
	// Writeable Returns the number of writable bytes which is equal to
	//
	// this.capacity - this.writerIndex
	Writeable() uint

	Release()
	Resume()
}

type bytebuf struct {
	capacity    uint
	pool        Pool
	buf         bytez
	readerIndex uint
	writerIndex uint
	markedIndex uint
}

func (buf *bytebuf) Read(bytes []byte) (n int, err error) {
	// TODO
	return 0, nil
}

func (buf *bytebuf) ReadInt() (n int, err error) {
	// TODO
	return 0, nil
}

func (buf *bytebuf) Write(bytes []byte) (n int, err error) {
	// TODO
	return 0, nil
}

func (buf *bytebuf) WriteInt(v int) (n int, err error) {
	// TODO
	return 0, nil
}

func (buf *bytebuf) Capacity() uint {
	return uint(len(buf.buf))
}

func (buf *bytebuf) Readable() uint {
	return buf.writerIndex - buf.readerIndex
}

func (buf *bytebuf) Writeable() uint {
	return buf.Capacity() - buf.writerIndex
}

func (buf *bytebuf) Release() {
	buf.pool.Free(buf.buf)
}

func (buf *bytebuf) Resume() {
	buf.buf = buf.pool.Alloc(buf.capacity)
}

func NewByteBuf(capacity uint, pools ...Pool) ByteBuf {
	buff := &bytebuf{
		capacity:    capacity,
		readerIndex: 0,
		writerIndex: 0,
	}

	if len(pools) == 0 {
		buff.pool = GetDefaultPoolInstance()
	} else {
		buff.pool = pools[0]
	}
	buff.buf = buff.pool.Alloc(capacity)

	return buff
}
