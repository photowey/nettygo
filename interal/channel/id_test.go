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

package channel

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_channelId_LongText(t *testing.T) {
	type fields struct {
		data       [28]byte
		shortValue string
		longValue  string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Test shortText()",
			fields: fields{
				data: [28]byte{
					0, 80, 86, 255, 254, 192, 0, 1, // merchantId
					0, 0, 42, 120, // processId
					0, 0, 0, 0, // sequence
					24, 134, 160, 193, 56, 156, 198, 130, // timestamp
					210, 208, 178, 175, // random
				},
			},
			want: "005056fffec00001-00002a78-00000000-1886a0c1389cc682-d2d0b2af",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			chId := channelId{
				data:       tt.fields.data,
				shortValue: tt.fields.shortValue,
				longValue:  tt.fields.longValue,
			}
			if got := chId.LongText(); got != tt.want {
				t.Errorf("LongText() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_channelId_ShortText(t *testing.T) {
	type fields struct {
		data       [28]byte
		shortValue string
		longValue  string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Test shortText()",
			fields: fields{
				data: [28]byte{
					0, 80, 86, 255, 254, 192, 0, 1, // merchantId
					0, 0, 42, 120, // processId
					0, 0, 0, 0, // sequence
					24, 134, 160, 193, 56, 156, 198, 130, // timestamp
					210, 208, 178, 175, // random
				},
			},
			want: "d2d0b2af",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			chId := &channelId{
				data:       tt.fields.data,
				shortValue: tt.fields.shortValue,
				longValue:  tt.fields.longValue,
			}
			if got := chId.ShortText(); got != tt.want {
				t.Errorf("ShortText() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewChannelId(t *testing.T) {
	tests := []struct {
		name string
		want Id
	}{
		{
			name: "Test new channel id",
			want: &channelId{
				sequence: 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewChannelId()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewChannelId() = %v, want %v", got, tt.want)
			}

			macAddrs, _ := macAddress()
			fmt.Println(macAddrs)
			for _, mac := range macAddrs {
				fmt.Println(macToInt64(mac))
			}
		})
	}
}
