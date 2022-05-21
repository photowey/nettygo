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

// ---------------------------------------------------------------- any

func TrinaryOperation(expression bool, standard, standBy any) any {
	if expression {
		return standard
	}
	return standBy
}

// ---------------------------------------------------------------- string

func TrinaryOperationString(expression bool, standard, standBy string) string {
	if expression {
		return standard
	}
	return standBy
}

// ---------------------------------------------------------------- int

func TrinaryOperationInt64(expression bool, standard, standBy int64) int64 {
	if expression {
		return standard
	}
	return standBy
}

func TrinaryOperationInt32(expression bool, standard, standBy int32) int32 {
	if expression {
		return standard
	}
	return standBy
}

func TrinaryOperationInt(expression bool, standard, standBy int) int {
	if expression {
		return standard
	}
	return standBy
}

func TrinaryOperationInt16(expression bool, standard, standBy int16) int16 {
	if expression {
		return standard
	}
	return standBy
}

func TrinaryOperationInt8(expression bool, standard, standBy int8) int8 {
	if expression {
		return standard
	}
	return standBy
}

// ----------------------------------------------------------------

func TrinaryOperationUInt64(expression bool, standard, standBy uint64) uint64 {
	if expression {
		return standard
	}
	return standBy
}

func TrinaryOperationUInt32(expression bool, standard, standBy uint32) uint32 {
	if expression {
		return standard
	}
	return standBy
}

func TrinaryOperationUInt(expression bool, standard, standBy uint) uint {
	if expression {
		return standard
	}
	return standBy
}

func TrinaryOperationUInt16(expression bool, standard, standBy uint16) uint16 {
	if expression {
		return standard
	}
	return standBy
}

func TrinaryOperationUInt8(expression bool, standard, standBy uint8) uint8 {
	if expression {
		return standard
	}
	return standBy
}

// ---------------------------------------------------------------- float

func TrinaryOperationFloat64(expression bool, standard, standBy float64) float64 {
	if expression {
		return standard
	}
	return standBy
}

func TrinaryOperationFloat32(expression bool, standard, standBy float32) float32 {
	if expression {
		return standard
	}
	return standBy
}
