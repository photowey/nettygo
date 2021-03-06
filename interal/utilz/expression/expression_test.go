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

package expression

import (
	"reflect"
	"testing"
)

func TestTrinaryOperation(t *testing.T) {
	type args struct {
		expression bool
		standard   any
		standBy    any
	}
	tests := []struct {
		name string
		args args
		want any
	}{
		{
			name: "Test TrinaryOperation-any-false",
			args: args{
				expression: 1 > 2,
				standard:   1,
				standBy:    2,
			},
			want: 2,
		},
		{
			name: "Test TrinaryOperation-any-false",
			args: args{
				expression: 1 > 2,
				standard:   1,
				standBy:    "hello",
			},
			want: "hello",
		},
		{
			name: "Test TrinaryOperation-any-true",
			args: args{
				expression: 1 < 2,
				standard:   1,
				standBy:    2,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TrinaryOperation(tt.args.expression, tt.args.standard, tt.args.standBy); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TrinaryOperation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTrinaryOperationFloat32(t *testing.T) {
	type args struct {
		expression bool
		standard   float32
		standBy    float32
	}
	tests := []struct {
		name string
		args args
		want float32
	}{
		{
			name: "Test TrinaryOperation-float32-true",
			args: args{
				expression: 1 < 2,
				standard:   1.000009,
				standBy:    2.000001,
			},
			want: 1.000009,
		},
		{
			name: "Test TrinaryOperation-float32-false",
			args: args{
				expression: 1 > 2,
				standard:   1.000009,
				standBy:    2.000001,
			},
			want: 2.000001,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TrinaryOperationFloat32(tt.args.expression, tt.args.standard, tt.args.standBy); got != tt.want {
				t.Errorf("TrinaryOperationFloat32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTrinaryOperationFloat64(t *testing.T) {
	type args struct {
		expression bool
		standard   float64
		standBy    float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "Test TrinaryOperation-float64-true",
			args: args{
				expression: 1 < 2,
				standard:   1.000009,
				standBy:    2.000001,
			},
			want: 1.000009,
		},
		{
			name: "Test TrinaryOperation-float64-false",
			args: args{
				expression: 1 > 2,
				standard:   1.000009,
				standBy:    2.000001,
			},
			want: 2.000001,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TrinaryOperationFloat64(tt.args.expression, tt.args.standard, tt.args.standBy); got != tt.want {
				t.Errorf("TrinaryOperationFloat64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTrinaryOperationInt(t *testing.T) {
	type args struct {
		expression bool
		standard   int
		standBy    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test TrinaryOperation-int-true",
			args: args{
				expression: 1 < 2,
				standard:   1,
				standBy:    2,
			},
			want: 1,
		},
		{
			name: "Test TrinaryOperation-int-false",
			args: args{
				expression: 1 > 2,
				standard:   1,
				standBy:    2,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TrinaryOperationInt(tt.args.expression, tt.args.standard, tt.args.standBy); got != tt.want {
				t.Errorf("TrinaryOperationInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTrinaryOperationInt16(t *testing.T) {
	type args struct {
		expression bool
		standard   int16
		standBy    int16
	}
	tests := []struct {
		name string
		args args
		want int16
	}{
		{
			name: "Test TrinaryOperation-int16-true",
			args: args{
				expression: 1 < 2,
				standard:   1,
				standBy:    2,
			},
			want: 1,
		},
		{
			name: "Test TrinaryOperation-int16-false",
			args: args{
				expression: 1 > 2,
				standard:   1,
				standBy:    2,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TrinaryOperationInt16(tt.args.expression, tt.args.standard, tt.args.standBy); got != tt.want {
				t.Errorf("TrinaryOperationInt16() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTrinaryOperationInt32(t *testing.T) {
	type args struct {
		expression bool
		standard   int32
		standBy    int32
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		{
			name: "Test TrinaryOperation-int32-true",
			args: args{
				expression: 1 < 2,
				standard:   1,
				standBy:    2,
			},
			want: 1,
		},
		{
			name: "Test TrinaryOperation-int32-false",
			args: args{
				expression: 1 > 2,
				standard:   1,
				standBy:    2,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TrinaryOperationInt32(tt.args.expression, tt.args.standard, tt.args.standBy); got != tt.want {
				t.Errorf("TrinaryOperationInt32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTrinaryOperationInt64(t *testing.T) {
	type args struct {
		expression bool
		standard   int64
		standBy    int64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "Test TrinaryOperation-int64-true",
			args: args{
				expression: 1 < 2,
				standard:   1,
				standBy:    2,
			},
			want: 1,
		},
		{
			name: "Test TrinaryOperation-int64-false",
			args: args{
				expression: 1 > 2,
				standard:   1,
				standBy:    2,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TrinaryOperationInt64(tt.args.expression, tt.args.standard, tt.args.standBy); got != tt.want {
				t.Errorf("TrinaryOperationInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTrinaryOperationInt8(t *testing.T) {
	type args struct {
		expression bool
		standard   int8
		standBy    int8
	}
	tests := []struct {
		name string
		args args
		want int8
	}{
		{
			name: "Test TrinaryOperation-int8-true",
			args: args{
				expression: 1 < 2,
				standard:   1,
				standBy:    2,
			},
			want: 1,
		},
		{
			name: "Test TrinaryOperation-int8-false",
			args: args{
				expression: 1 > 2,
				standard:   1,
				standBy:    2,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TrinaryOperationInt8(tt.args.expression, tt.args.standard, tt.args.standBy); got != tt.want {
				t.Errorf("TrinaryOperationInt8() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTrinaryOperationString(t *testing.T) {
	type args struct {
		expression bool
		standard   string
		standBy    string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test TrinaryOperation-any-false",
			args: args{
				expression: 1 > 2,
				standard:   "standard",
				standBy:    "standBy",
			},
			want: "standBy",
		},
		{
			name: "Test TrinaryOperation-string-true",
			args: args{
				expression: 1 < 2,
				standard:   "standard",
				standBy:    "standBy",
			},
			want: "standard",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TrinaryOperationString(tt.args.expression, tt.args.standard, tt.args.standBy); got != tt.want {
				t.Errorf("TrinaryOperationString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTrinaryOperationUInt(t *testing.T) {
	type args struct {
		expression bool
		standard   uint
		standBy    uint
	}
	tests := []struct {
		name string
		args args
		want uint
	}{
		{
			name: "Test TrinaryOperation-uint-true",
			args: args{
				expression: 1 < 2,
				standard:   1,
				standBy:    2,
			},
			want: 1,
		},
		{
			name: "Test TrinaryOperation-uint-false",
			args: args{
				expression: 1 > 2,
				standard:   1,
				standBy:    2,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TrinaryOperationUInt(tt.args.expression, tt.args.standard, tt.args.standBy); got != tt.want {
				t.Errorf("TrinaryOperationUInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTrinaryOperationUInt16(t *testing.T) {
	type args struct {
		expression bool
		standard   uint16
		standBy    uint16
	}
	tests := []struct {
		name string
		args args
		want uint16
	}{
		{
			name: "Test TrinaryOperation-uint16-true",
			args: args{
				expression: 1 < 2,
				standard:   1,
				standBy:    2,
			},
			want: 1,
		},
		{
			name: "Test TrinaryOperation-uint16-false",
			args: args{
				expression: 1 > 2,
				standard:   1,
				standBy:    2,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TrinaryOperationUInt16(tt.args.expression, tt.args.standard, tt.args.standBy); got != tt.want {
				t.Errorf("TrinaryOperationUInt16() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTrinaryOperationUInt32(t *testing.T) {
	type args struct {
		expression bool
		standard   uint32
		standBy    uint32
	}
	tests := []struct {
		name string
		args args
		want uint32
	}{
		{
			name: "Test TrinaryOperation-uint32-true",
			args: args{
				expression: 1 < 2,
				standard:   1,
				standBy:    2,
			},
			want: 1,
		},
		{
			name: "Test TrinaryOperation-uint32-false",
			args: args{
				expression: 1 > 2,
				standard:   1,
				standBy:    2,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TrinaryOperationUInt32(tt.args.expression, tt.args.standard, tt.args.standBy); got != tt.want {
				t.Errorf("TrinaryOperationUInt32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTrinaryOperationUInt64(t *testing.T) {
	type args struct {
		expression bool
		standard   uint64
		standBy    uint64
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		{
			name: "Test TrinaryOperation-uint64-true",
			args: args{
				expression: 1 < 2,
				standard:   1,
				standBy:    2,
			},
			want: 1,
		},
		{
			name: "Test TrinaryOperation-uint64-false",
			args: args{
				expression: 1 > 2,
				standard:   1,
				standBy:    2,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TrinaryOperationUInt64(tt.args.expression, tt.args.standard, tt.args.standBy); got != tt.want {
				t.Errorf("TrinaryOperationUInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTrinaryOperationUInt8(t *testing.T) {
	type args struct {
		expression bool
		standard   uint8
		standBy    uint8
	}
	tests := []struct {
		name string
		args args
		want uint8
	}{
		{
			name: "Test TrinaryOperation-uint8-true",
			args: args{
				expression: 1 < 2,
				standard:   1,
				standBy:    2,
			},
			want: 1,
		},
		{
			name: "Test TrinaryOperation-uint8-false",
			args: args{
				expression: 1 > 2,
				standard:   1,
				standBy:    2,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TrinaryOperationUInt8(tt.args.expression, tt.args.standard, tt.args.standBy); got != tt.want {
				t.Errorf("TrinaryOperationUInt8() = %v, want %v", got, tt.want)
			}
		})
	}
}
