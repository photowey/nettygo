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
			name: "Test byte to int-1",
			args: args{
				bytez: []byte{0x00, 0x00, 0x00, 0x01},
			},
			want: 1,
		},
		{
			name: "Test byte to int-17",
			args: args{
				bytez: []byte{0x00, 0x00, 0x01, 0x01},
			},
			want: 257,
		},
		{
			name: "Test byte to int-65793",
			args: args{
				bytez: []byte{0x00, 0x01, 0x01, 0x01},
			},
			want: 65_793, // 1<<16 + 1<<8 + 1<<0
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
		bytez []byte
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "Test byte to int64-1",
			args: args{
				bytez: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01},
			},
			want: 1,
		},
		{
			name: "Test byte to int64-257",
			args: args{
				bytez: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x01},
			},
			want: 257,
		},
		{
			name: "Test byte to int64-65793",
			args: args{
				bytez: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x01, 0x01},
			},
			want: 65_793,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ByteToInt64(tt.args.bytez); got != tt.want {
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
		{
			name: "Test byte to uint16-1",
			args: args{
				bytez: []byte{0x00, 0x01},
			},
			want: 1,
		},
		{
			name: "Test byte to uint16-257",
			args: args{
				bytez: []byte{0x01, 0x01},
			},
			want: 257,
		},
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
		{
			name: "Test byte to uint32-1",
			args: args{
				bytez: []byte{0x00, 0x00, 0x00, 0x01},
			},
			want: 1,
		},
		{
			name: "Test byte to uint32-257",
			args: args{
				bytez: []byte{0x00, 0x00, 0x01, 0x01},
			},
			want: 257,
		},
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
		{
			name: "Test byte to uint64-1",
			args: args{
				bytez: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01},
			},
			want: 1,
		},
		{
			name: "Test byte to uint64-257",
			args: args{
				bytez: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x01},
			},
			want: 257,
		},
		{
			name: "Test byte to uint64-65793",
			args: args{
				bytez: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x01, 0x01},
			},
			want: 65_793,
		},
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
		{
			name: "Test int64 to byte-1",
			args: args{
				v: 1,
			},
			want: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01},
		},
		{
			name: "Test int64 to byte-257",
			args: args{
				v: 257,
			},
			want: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x01},
		},
		{
			name: "Test int64 to byte-65793",
			args: args{
				v: 65_793,
			},
			want: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x01, 0x01},
		},
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
		{
			name: "Test int64 to byte-1",
			args: args{
				v:     1,
				bytez: make([]byte, 8),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Int64ToBytez(tt.args.v, tt.args.bytez)
			if got := ByteToInt64(tt.args.bytez); got != tt.args.v {
				t.Errorf("ByteToInt64() = %v, want %v", got, tt.args.v)
			}
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
		{
			name: "Test int to bytes-1",
			args: args{
				v: 1,
			},
			want: []byte{0x00, 0x00, 0x00, 0x01},
		},
		{
			name: "Test int to bytes-257",
			args: args{
				v: 257,
			},
			want: []byte{0x00, 0x00, 0x01, 0x01},
		},
		{
			name: "Test int to bytes-65793",
			args: args{
				v: 65_793,
			},
			want: []byte{0x00, 0x01, 0x01, 0x01},
		},
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
		{
			name: "Test int to bytez-65793",
			args: args{
				v:     65_793,
				bytez: make([]byte, 4),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			IntToBytez(tt.args.v, tt.args.bytez)
			if got := ByteToInt(tt.args.bytez); got != tt.args.v {
				t.Errorf("ByteToInt() = %v, want %v", got, tt.args.v)
			}
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
		{
			name: "Test uint16 to bytes-257",
			args: args{
				v: 257,
			},
			want: []byte{0x01, 0x01},
		},
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
		{
			name: "Test uint16 to bytes-257",
			args: args{
				v:     257,
				bytez: make([]byte, 2),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Uint16ToBytez(tt.args.v, tt.args.bytez)
			if got := ByteToUInt16(tt.args.bytez); got != tt.args.v {
				t.Errorf("ByteToUInt16() = %v, want %v", got, tt.args.v)
			}
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
		{
			name: "Test uint16 to bytes-65793",
			args: args{
				v: 65_793,
			},
			want: []byte{0x00, 0x01, 0x01, 0x01},
		},
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
		{
			name: "Test uint32 to bytes-257",
			args: args{
				v:     257,
				bytez: make([]byte, 4),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Uint32ToBytez(tt.args.v, tt.args.bytez)
			if got := ByteToUInt32(tt.args.bytez); got != tt.args.v {
				t.Errorf("ByteToUInt32() = %v, want %v", got, tt.args.v)
			}
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
		{
			name: "Test uint64 to bytes-65793",
			args: args{
				v: 65_793,
			},
			want: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x01, 0x01},
		},
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
		{
			name: "Test uint64 to bytes-65793",
			args: args{
				v:     65_793,
				bytez: make([]byte, 8),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Uint64ToBytez(tt.args.v, tt.args.bytez)
			if got := ByteToUInt64(tt.args.bytez); got != tt.args.v {
				t.Errorf("ByteToUInt64() = %v, want %v", got, tt.args.v)
			}
		})
	}
}

func Test_bytebuf_Capacity(t *testing.T) {
	type fields struct {
		capacity    uint
		pool        Pool
		buf         bytez
		readerIndex uint
		writerIndex uint
		markedIndex uint
	}

	tests := []struct {
		name   string
		fields fields
		want   uint
	}{
		{
			name: "Test bytebuf capacity",
			fields: fields{
				capacity: 8,
			},
			want: 8,
		},
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
		capacity    uint
		pool        Pool
		buf         bytez
		readerIndex uint
		writerIndex uint
		markedIndex uint
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
		capacity    uint
		pool        Pool
		buf         bytez
		readerIndex uint
		writerIndex uint
		markedIndex uint
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
		capacity    uint
		pool        Pool
		buf         bytez
		readerIndex uint
		writerIndex uint
		markedIndex uint
	}
	tests := []struct {
		name   string
		fields fields
		want   uint
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
		capacity    uint
		pool        Pool
		buf         bytez
		readerIndex uint
		writerIndex uint
		markedIndex uint
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
		capacity    uint
		pool        Pool
		buf         bytez
		readerIndex uint
		writerIndex uint
		markedIndex uint
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
		capacity    uint
		pool        Pool
		buf         bytez
		readerIndex uint
		writerIndex uint
		markedIndex uint
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
		capacity    uint
		pool        Pool
		buf         bytez
		readerIndex uint
		writerIndex uint
		markedIndex uint
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
		capacity    uint
		pool        Pool
		buf         bytez
		readerIndex uint
		writerIndex uint
		markedIndex uint
	}
	tests := []struct {
		name   string
		fields fields
		want   uint
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
