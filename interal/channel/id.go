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
	"encoding/hex"
	"fmt"
)

const (
	MerchantIdLength = 8
	ProcessIdLength  = 4
	SequenceLength   = 4
	TimestampLength  = 8
	RandomLength     = 4
)

var _ Id = (*channelId)(nil)

type Id interface {
	ShortText() string
	LongText() string
}

type channelId struct {
	data       [28]byte // 8 + 4 + 4 + 8 + 4
	shortValue string
	longValue  string
}

func (chId *channelId) ShortText() string {
	v := chId.data[len(chId.data)-RandomLength:]
	random := dumpHexString(v)

	return random
}

/*
	// 8 + 4 + 4 + 8 + 4
	// 005056fffec00001-00002a78-00000000-1886a0c1389cc682-d2d0b2af
	// d2d0b2af
	//
	// Java byte
	// [0 80 86 -1 -2 -64 0 1 0 0 42 120 0 0 0 0 24 -122 -96 -63 56 -100 -58 -126 -46 -48 -78 -81]
	// Go byte
	// [0 80 86 255 254 192 0 1 0 0 42 120 0 0 0 0 24 134 160 193 56 156 198 130 210 208 178 175]
*/
func (chId *channelId) LongText() string {
	idx := 0

	v1 := chId.data[idx : idx+MerchantIdLength]
	merchantId := dumpHexString(v1)
	idx += MerchantIdLength

	v2 := chId.data[idx : +idx+ProcessIdLength]
	processId := dumpHexString(v2)
	idx += ProcessIdLength

	v3 := chId.data[idx : idx+SequenceLength]
	sequence := dumpHexString(v3)
	idx += SequenceLength

	v4 := chId.data[idx : idx+TimestampLength]
	timestamp := dumpHexString(v4)
	idx += TimestampLength

	v5 := chId.data[idx : idx+RandomLength]
	random := dumpHexString(v5)
	idx += RandomLength

	lv := fmt.Sprintf("%s-%s-%s-%s-%s", merchantId, processId, sequence, timestamp, random)

	return lv
}

func dumpHexString(hexByte []byte) string {
	hexStr := hex.EncodeToString(hexByte)

	return hexStr
}
