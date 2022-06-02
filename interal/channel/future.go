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
	"github.com/photowey/nettygo/interal/concurrent"
)

var _ Future = (*SucceededFuture)(nil)

type Future interface {
	concurrent.Future
	Channel() Channel
}

type SucceededFuture struct {
	channel Channel
}

func (fu SucceededFuture) IsSuccess() bool {
	return false
}

func (fu SucceededFuture) IsCancellable() bool {
	return false
}

func (fu SucceededFuture) Cause() error {
	return nil
}

func (fu *SucceededFuture) Sync() any {
	return fu
}

func (fu *SucceededFuture) Await() any {
	return fu
}

func (fu SucceededFuture) Channel() Channel {
	return nil
}

func NewSucceededFuture(ch Channel) Future {
	return &SucceededFuture{
		channel: ch,
	}
}
