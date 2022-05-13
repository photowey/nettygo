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
	"bytes"
)

var (
	_ ByteBuf = (*bytebuf)(nil)
)

type ByteBuf interface {
	Read(bytes []byte) (n int, err error)
	ReadInt() (int, error)
	Write(bytes []byte) (int, error)
	WriteInt(v int) (n int, err error)
	Capacity() int
	Release()
	Resume()
}

type Pool interface {
	Alloc(int) []byte
	Free([]byte)
}

type bytebuf struct {
	capacity    int
	pool        Pool
	buf         *bytes.Buffer
	readerIndex int
	writerIndex int
	markedIndex int
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

func (buf *bytebuf) Capacity() int {
	return len(buf.buf.Bytes())
}

func (buf *bytebuf) Release() {
	buf.pool.Free(buf.buf.Bytes())
}

func (buf *bytebuf) Resume() {
	buf.buf = bytes.NewBuffer(buf.pool.Alloc(buf.capacity))
}