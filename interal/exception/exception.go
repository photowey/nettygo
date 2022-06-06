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

package exception

import (
	"io"
	"runtime"
)

const (
	defaultSkip = 3  // default skip depth
	depth       = 32 // default stack depth
	single      = 1  // [x]
)

type Stack []uintptr

type PrintWriter struct {
	writer io.Writer
}

type Exception interface {
	Unwrap() error
	Error() string
	Stack() Stack
	PrintStackTrace(pw PrintWriter, message ...string)
}

type exception struct {
	error error
	stack Stack
}

func (ex exception) Unwrap() error {
	return ex.error
}

func (ex exception) Error() string {
	return ex.error.Error()
}

func (ex exception) Stack() Stack {
	return ex.stack
}

func (ex exception) PrintStackTrace(pw PrintWriter, message ...string) {
	// TODO
}

func parseStack(skipd ...int) Stack {
	skip := defaultSkip
	switch len(skipd) {
	case single:
		skip = skipd[0]
	}
	var pcs [depth]uintptr
	length := runtime.Callers(skip, pcs[:])

	stack := pcs[:length]

	return stack
}
