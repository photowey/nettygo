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
	"fmt"

	"github.com/photowey/nettygo/interal/byteunit"
)

func Format(bytes uint64) string {
	switch {
	case bytes < byteunit.KB:
		return fmt.Sprintf("%dB", bytes)
	case bytes < byteunit.MB:
		return fmt.Sprintf("%.2fK", float64(bytes)/byteunit.KB)
	case bytes < byteunit.GB:
		return fmt.Sprintf("%.2fM", float64(bytes)/byteunit.MB)
	case bytes < byteunit.TB:
		return fmt.Sprintf("%.2fG", float64(bytes)/byteunit.GB)
	case bytes < byteunit.PB:
		return fmt.Sprintf("%.2fT", float64(bytes)/byteunit.TB)
	default:
		return fmt.Sprintf("%.2fP", float64(bytes)/byteunit.PB)
	}
}
