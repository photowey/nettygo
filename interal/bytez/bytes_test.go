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

package bytez

import (
	"testing"
)

func TestFormat(t *testing.T) {
	type args struct {
		bytes uint64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test byte format-B",
			args: args{
				bytes: 418,
			},
			want: "418B",
		},
		{
			name: "Test byte format-KB",
			args: args{
				bytes: 6379,
			},
			want: "6.23K",
		},
		{
			name: "Test byte format-MB",
			args: args{
				bytes: 6379876,
			},
			want: "6.08M",
		},
		{
			name: "Test byte format-GB",
			args: args{
				bytes: 1234567890,
			},
			want: "1.15G",
		},
		{
			name: "Test byte format-TB",
			args: args{
				bytes: 12345678901234,
			},
			want: "11.23T",
		},
		{
			name: "Test byte format-PB",
			args: args{
				bytes: 12345678901234321,
			},
			want: "10.97P",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Format(tt.args.bytes); got != tt.want {
				t.Errorf("Format() = %v, want %v", got, tt.want)
			}
		})
	}
}
