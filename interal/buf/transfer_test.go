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
	"reflect"
	"testing"
)

func TestByteToInt(t *testing.T) {
	type args struct {
		bytez []byte
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test byte to int",
			args: args{
				bytez: []byte{0x00, 0x00, 0x00, 0x00, 0x01},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ByteToInt(tt.args.bytez); got != tt.want {
				t.Errorf("ByteToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestByteToInt64(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ByteToInt64(tt.args.data); got != tt.want {
				t.Errorf("ByteToInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestByteToUInt16(t *testing.T) {
	type args struct {
		bytez []byte
	}
	tests := []struct {
		name string
		args args
		want uint16
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ByteToUInt16(tt.args.bytez); got != tt.want {
				t.Errorf("ByteToUInt16() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestByteToUInt32(t *testing.T) {
	type args struct {
		bytez []byte
	}
	tests := []struct {
		name string
		args args
		want uint32
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ByteToUInt32(tt.args.bytez); got != tt.want {
				t.Errorf("ByteToUInt32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestByteToUInt64(t *testing.T) {
	type args struct {
		bytez []byte
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ByteToUInt64(tt.args.bytez); got != tt.want {
				t.Errorf("ByteToUInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt64ToBytes(t *testing.T) {
	type args struct {
		v int64
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int64ToBytes(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Int64ToBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt64ToBytez(t *testing.T) {
	type args struct {
		v     int64
		bytez []byte
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Int64ToBytez(tt.args.v, tt.args.bytez)
		})
	}
}

func TestIntToBytes(t *testing.T) {
	type args struct {
		v int
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IntToBytes(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IntToBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIntToBytez(t *testing.T) {
	type args struct {
		v     int
		bytez []byte
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			IntToBytez(tt.args.v, tt.args.bytez)
		})
	}
}

func TestUint16ToBytes(t *testing.T) {
	type args struct {
		v uint16
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Uint16ToBytes(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Uint16ToBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint16ToBytez(t *testing.T) {
	type args struct {
		v     uint16
		bytez []byte
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Uint16ToBytez(tt.args.v, tt.args.bytez)
		})
	}
}

func TestUint32ToBytes(t *testing.T) {
	type args struct {
		v uint32
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Uint32ToBytes(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Uint32ToBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint32ToBytez(t *testing.T) {
	type args struct {
		v     uint32
		bytez []byte
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Uint32ToBytez(tt.args.v, tt.args.bytez)
		})
	}
}

func TestUint64ToBytes(t *testing.T) {
	type args struct {
		v uint64
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Uint64ToBytes(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Uint64ToBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint64ToBytez(t *testing.T) {
	type args struct {
		v     uint64
		bytez []byte
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Uint64ToBytez(tt.args.v, tt.args.bytez)
		})
	}
}

func Test_bytebuf_Capacity(t *testing.T) {
	type fields struct {
		capacity    int
		pool        Pool
		buf         *bytes.Buffer
		readerIndex int
		writerIndex int
		markedIndex int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf := &bytebuf{
				capacity:    tt.fields.capacity,
				pool:        tt.fields.pool,
				buf:         tt.fields.buf,
				readerIndex: tt.fields.readerIndex,
				writerIndex: tt.fields.writerIndex,
				markedIndex: tt.fields.markedIndex,
			}
			if got := buf.Capacity(); got != tt.want {
				t.Errorf("Capacity() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_bytebuf_Read(t *testing.T) {
	type fields struct {
		capacity    int
		pool        Pool
		buf         *bytes.Buffer
		readerIndex int
		writerIndex int
		markedIndex int
	}
	type args struct {
		bytes []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantN   int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf := &bytebuf{
				capacity:    tt.fields.capacity,
				pool:        tt.fields.pool,
				buf:         tt.fields.buf,
				readerIndex: tt.fields.readerIndex,
				writerIndex: tt.fields.writerIndex,
				markedIndex: tt.fields.markedIndex,
			}
			gotN, err := buf.Read(tt.args.bytes)
			if (err != nil) != tt.wantErr {
				t.Errorf("Read() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotN != tt.wantN {
				t.Errorf("Read() gotN = %v, want %v", gotN, tt.wantN)
			}
		})
	}
}

func Test_bytebuf_ReadInt(t *testing.T) {
	type fields struct {
		capacity    int
		pool        Pool
		buf         *bytes.Buffer
		readerIndex int
		writerIndex int
		markedIndex int
	}
	tests := []struct {
		name    string
		fields  fields
		wantN   int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf := &bytebuf{
				capacity:    tt.fields.capacity,
				pool:        tt.fields.pool,
				buf:         tt.fields.buf,
				readerIndex: tt.fields.readerIndex,
				writerIndex: tt.fields.writerIndex,
				markedIndex: tt.fields.markedIndex,
			}
			gotN, err := buf.ReadInt()
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadInt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotN != tt.wantN {
				t.Errorf("ReadInt() gotN = %v, want %v", gotN, tt.wantN)
			}
		})
	}
}

func Test_bytebuf_Readable(t *testing.T) {
	type fields struct {
		capacity    int
		pool        Pool
		buf         *bytes.Buffer
		readerIndex int
		writerIndex int
		markedIndex int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf := &bytebuf{
				capacity:    tt.fields.capacity,
				pool:        tt.fields.pool,
				buf:         tt.fields.buf,
				readerIndex: tt.fields.readerIndex,
				writerIndex: tt.fields.writerIndex,
				markedIndex: tt.fields.markedIndex,
			}
			if got := buf.Readable(); got != tt.want {
				t.Errorf("Readable() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_bytebuf_Release(t *testing.T) {
	type fields struct {
		capacity    int
		pool        Pool
		buf         *bytes.Buffer
		readerIndex int
		writerIndex int
		markedIndex int
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf := &bytebuf{
				capacity:    tt.fields.capacity,
				pool:        tt.fields.pool,
				buf:         tt.fields.buf,
				readerIndex: tt.fields.readerIndex,
				writerIndex: tt.fields.writerIndex,
				markedIndex: tt.fields.markedIndex,
			}
			buf.Release()
		})
	}
}

func Test_bytebuf_Resume(t *testing.T) {
	type fields struct {
		capacity    int
		pool        Pool
		buf         *bytes.Buffer
		readerIndex int
		writerIndex int
		markedIndex int
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf := &bytebuf{
				capacity:    tt.fields.capacity,
				pool:        tt.fields.pool,
				buf:         tt.fields.buf,
				readerIndex: tt.fields.readerIndex,
				writerIndex: tt.fields.writerIndex,
				markedIndex: tt.fields.markedIndex,
			}
			buf.Resume()
		})
	}
}

func Test_bytebuf_Write(t *testing.T) {
	type fields struct {
		capacity    int
		pool        Pool
		buf         *bytes.Buffer
		readerIndex int
		writerIndex int
		markedIndex int
	}
	type args struct {
		bytes []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantN   int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf := &bytebuf{
				capacity:    tt.fields.capacity,
				pool:        tt.fields.pool,
				buf:         tt.fields.buf,
				readerIndex: tt.fields.readerIndex,
				writerIndex: tt.fields.writerIndex,
				markedIndex: tt.fields.markedIndex,
			}
			gotN, err := buf.Write(tt.args.bytes)
			if (err != nil) != tt.wantErr {
				t.Errorf("Write() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotN != tt.wantN {
				t.Errorf("Write() gotN = %v, want %v", gotN, tt.wantN)
			}
		})
	}
}

func Test_bytebuf_WriteInt(t *testing.T) {
	type fields struct {
		capacity    int
		pool        Pool
		buf         *bytes.Buffer
		readerIndex int
		writerIndex int
		markedIndex int
	}
	type args struct {
		v int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantN   int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf := &bytebuf{
				capacity:    tt.fields.capacity,
				pool:        tt.fields.pool,
				buf:         tt.fields.buf,
				readerIndex: tt.fields.readerIndex,
				writerIndex: tt.fields.writerIndex,
				markedIndex: tt.fields.markedIndex,
			}
			gotN, err := buf.WriteInt(tt.args.v)
			if (err != nil) != tt.wantErr {
				t.Errorf("WriteInt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotN != tt.wantN {
				t.Errorf("WriteInt() gotN = %v, want %v", gotN, tt.wantN)
			}
		})
	}
}

func Test_bytebuf_Writeable(t *testing.T) {
	type fields struct {
		capacity    int
		pool        Pool
		buf         *bytes.Buffer
		readerIndex int
		writerIndex int
		markedIndex int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf := &bytebuf{
				capacity:    tt.fields.capacity,
				pool:        tt.fields.pool,
				buf:         tt.fields.buf,
				readerIndex: tt.fields.readerIndex,
				writerIndex: tt.fields.writerIndex,
				markedIndex: tt.fields.markedIndex,
			}
			if got := buf.Writeable(); got != tt.want {
				t.Errorf("Writeable() = %v, want %v", got, tt.want)
			}
		})
	}
}
