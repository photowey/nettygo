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
	"github.com/photowey/nettygo/interal/exception"
)

type Base interface {
	HandlerAdded(ctx HandlerContext)
	HandlerRemoved(ctx HandlerContext)
	ExceptionCaught(ctx HandlerContext, ex exception.Exception)
	Inbound() int // -1: in 0:all 1: out
}

type Handler interface {
	Base
	HeadHandler
	ActiveHandler
	InboundHandler
	OutboundHandler
	ExceptionHandler
	InactiveHandler
	TailHandler
}

type HeadHandler interface {
	Head(ctx HeadContext)
}

type InboundHandler interface {
	HandleRead(ctx InboundContext, message Message)
}

type OutboundHandler interface {
	HandleWrite(ctx OutboundContext, message Message)
}

type ActiveHandler interface {
	HandleActive(ctx ActiveContext)
}

type InactiveHandler interface {
	HandleInactive(ctx InactiveContext, ex exception.Exception)
}

type ExceptionHandler interface {
	HandleException(ctx ExceptionContext, ex exception.Exception)
}

type TailHandler interface {
	Tail(ctx TailContext)
}
