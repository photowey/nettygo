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
	"encoding/binary"
)

const (
	IntBytes    = 4
	UInt16Bytes = 2
	UInt32Bytes = 4
	Int64Bytes  = 8
	UInt64Bytes = 8
)

func ByteToInt(bytez []byte) int {
	return (int(bytez[0])&0xff)<<24 |
		(int(bytez[1])&0xff)<<16 |
		(int(bytez[2])&0xff)<<8 |
		(int(bytez[3]) & 0xff)
}

func ByteToInt64(data []byte) int64 {
	return (int64(data[0])&0xff)<<56 |
		(int64(data[1])&0xff)<<48 |
		(int64(data[2])&0xff)<<40 |
		(int64(data[3])&0xff)<<32 |
		(int64(data[4])&0xff)<<24 |
		(int64(data[5])&0xff)<<16 |
		(int64(data[6])&0xff)<<8 |
		(int64(data[7]) & 0xff)
}

func ByteToUInt16(bytez []byte) uint16 {
	return binary.BigEndian.Uint16(bytez)
}

func ByteToUInt32(bytez []byte) uint32 {
	return binary.BigEndian.Uint32(bytez)
}

func ByteToUInt64(bytez []byte) uint64 {
	return binary.BigEndian.Uint64(bytez)
}

func IntToBytez(v int, bytez []byte) {
	bytez[0] = byte(v >> 24)
	bytez[1] = byte(v >> 16)
	bytez[2] = byte(v >> 8)
	bytez[3] = byte(v)
}

func IntToBytes(v int) []byte {
	bytez := make([]byte, IntBytes)
	IntToBytez(v, bytez)

	return bytez
}

func Int64ToBytez(v int64, bytez []byte) {
	bytez[0] = byte(v >> 56)
	bytez[1] = byte(v >> 48)
	bytez[2] = byte(v >> 40)
	bytez[3] = byte(v >> 32)
	bytez[4] = byte(v >> 24)
	bytez[5] = byte(v >> 16)
	bytez[6] = byte(v >> 8)
	bytez[7] = byte(v)
}

func Int64ToBytes(v int64) []byte {
	bytez := make([]byte, Int64Bytes)
	Int64ToBytez(v, bytez)

	return bytez
}

func Uint64ToBytez(v uint64, bytez []byte) {
	binary.BigEndian.PutUint64(bytez, v)
}

func Uint64ToBytes(v uint64) []byte {
	bytez := make([]byte, UInt64Bytes)
	Uint64ToBytez(v, bytez)

	return bytez
}

func Uint32ToBytez(v uint32, bytez []byte) {
	binary.BigEndian.PutUint32(bytez, v)
}

func Uint32ToBytes(v uint32) []byte {
	bytez := make([]byte, UInt32Bytes)
	Uint32ToBytez(v, bytez)

	return bytez
}

func Uint16ToBytez(v uint16, bytez []byte) {
	binary.BigEndian.PutUint16(bytez, v)
}

func Uint16ToBytes(v uint16) []byte {
	bytez := make([]byte, UInt16Bytes)
	Uint16ToBytez(v, bytez)

	return bytez
}
