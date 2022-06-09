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

const (
	Init = iota
	AddPending
	AddComplete
	RemoveComplete
)

type HandlerContext interface {
	Channel() Channel
	Handler() Handler
	Executor() concurrent.EventExecutor
	Close(err error)
	Attachment() Attachment
	SetAttachment(attr Attachment)
	Inbound() int8
}

type HeadContext interface {
	HandlerContext
}

type ExceptionContext interface {
	HandlerContext
}

type InboundContext interface {
	HandlerContext
}
type OutboundContext interface {
	HandlerContext
}

type ActiveContext interface {
	HandlerContext
}

type InactiveContext interface {
	HandlerContext
}

type EventContext interface {
	HandlerContext
	HandleEvent(event Event)
}

type TailContext interface {
	HandlerContext
}

type defaultHandlerContext struct {
	prev            *defaultHandlerContext
	next            *defaultHandlerContext
	name            string
	handler         Handler
	executor        concurrent.EventExecutor
	pipeline        Pipeline
	succeededFuture Future
	handlerState    int8
}

func (hctx *defaultHandlerContext) Channel() Channel {
	return hctx.pipeline.Channel()
}

func (hctx *defaultHandlerContext) Handler() Handler {
	return hctx.handler
}

func (hctx *defaultHandlerContext) Executor() concurrent.EventExecutor {
	return hctx.executor
}

func (hctx *defaultHandlerContext) Close(err error) {
	hctx.Channel().Close(err)
}

func (hctx *defaultHandlerContext) Attachment() Attachment {
	return hctx.Channel().Attachment
}

func (hctx *defaultHandlerContext) SetAttachment(attr Attachment) {
	hctx.Channel().SetAttachment(attr)
}

func (hctx *defaultHandlerContext) Inbound() int8 {
	return hctx.handler.Inbound() // ?
}

func (hctx *defaultHandlerContext) HandleEvent(event Event) {
	next := hctx
	for {
		if next = next.nextCtx(); next == nil {
			break
		}
		if handler, ok := next.Handler().(EventHandler); ok {
			handler.HandleEvent(next, event)
			break
		}
	}
}

func (hctx *defaultHandlerContext) prevCtx() *defaultHandlerContext {
	return hctx.prev
}

func (hctx *defaultHandlerContext) nextCtx() *defaultHandlerContext {
	return hctx.next
}

func newHeadContext(pl Pipeline) *defaultHandlerContext {
	// TODO
	return &defaultHandlerContext{
		pipeline:     pl,
		handlerState: Init,
	}
}

func newTailContext(pl Pipeline) *defaultHandlerContext {
	// TODO
	return &defaultHandlerContext{
		pipeline:     pl,
		handlerState: Init,
	}
}

func newContext(group concurrent.EventExecutorGroup, name string, handler Handler) *defaultHandlerContext {
	executor := childExecutor(group)

	ctx := &defaultHandlerContext{
		handlerState: Init,
	}
	ctx.name = name
	ctx.handler = handler
	ctx.executor = executor

	return ctx
}

func childExecutor(group concurrent.EventExecutorGroup) concurrent.EventExecutor {
	if group == nil {
		return nil
	}

	return nil
}
