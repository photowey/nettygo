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
	"math/rand"
	"net"
	"os"
	"sort"
	"strings"
	"sync/atomic"
	"time"

	"github.com/photowey/nettygo/interal/buf"
)

const (
	MerchantIdLength = 8
	ProcessIdLength  = 4
	SequenceLength   = 4
	TimestampLength  = 8
	RandomLength     = 4

	delta = 1

	maxRandom = 100_000_000
)

var (
	_        Id = (*channelId)(nil)
	sequence uint64
)

type Id interface {
	ShortText() string
	LongText() string
}

type channelId struct {
	data       [28]byte // 8 + 4 + 4 + 8 + 4
	shortValue string
	longValue  string
	sequence   uint64
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func NewChannelId() Id {
	data := [28]byte{}

	idx := 0
	merchantId := findMaxMacUInt64()
	processId := os.Getpid()
	timestamp := time.Now().UnixNano() / int64(time.Millisecond)
	random := rand.Intn(maxRandom)

	fmt.Println(idx)
	fmt.Println(merchantId)
	fmt.Println(processId)
	fmt.Println(timestamp)
	fmt.Println(random)

	sequenceNo := atomic.AddUint64(&sequence, delta)
	return &channelId{
		data:     data,
		sequence: sequenceNo,
	}
}

func (chId *channelId) ShortText() string {
	v := chId.data[len(chId.data)-RandomLength:]
	random := dumpHexString(v)

	return random
}

/*
	// 8 + 4 + 4 + 8 + 4
	// 005056fffec00001-00002a78-00000000-1886a0c1389cc682-d2d0b2af
	// merchantId-processId-sequence-timestamp-random
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

func macAddress() (macs []string, err error) {
	netInterfaces, err := net.Interfaces()
	if err != nil {
		return macs, err
	}

	for _, netInterface := range netInterfaces {
		macAddr := netInterface.HardwareAddr.String()
		if len(macAddr) == 0 {
			continue
		}

		macs = append(macs, macAddr)
	}

	return macs, nil
}

func macToInt64(mac string) uint64 {
	macBytes := make([]byte, 8)
	hexNum, _ := hex.DecodeString(strings.ReplaceAll(mac, ":", ""))
	for i, hexVal := range hexNum {
		macBytes[i+2] = hexVal
	}

	return buf.ByteToUInt64(macBytes)
}

func findMaxMacUInt64() uint64 {
	macAddrs, _ := macAddress() // TODO ignore err now.
	if len(macAddrs) == 0 {
		return 0
	}

	macUInt64s := make([]uint64, 0)
	for _, mac := range macAddrs {
		macUInt64 := macToInt64(mac)
		macUInt64s = append(macUInt64s, macUInt64)
	}
	sort.Slice(macUInt64s, func(i, j int) bool {
		return macUInt64s[i] > macUInt64s[j]
	})

	return macUInt64s[0]
}
