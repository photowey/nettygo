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
	"github.com/photowey/nettygo/interal/exception"
)

var _ Pipeline = (*pipeline)(nil)

type Pipeline interface {
	AddFirst(name string, handler Handler) Pipeline
	AddFirstGroup(group concurrent.EventExecutorGroup, name string, handler Handler) Pipeline
	AddLast(name string, handler Handler) Pipeline
	AddLastGroup(group concurrent.EventExecutorGroup, name string, handler Handler) Pipeline
	AddBefore(base, name string, handler Handler) Pipeline
	AddAfter(base, name string, handler Handler) Pipeline
	Remove(name string) Handler
	RemoveFirst(name string) Handler
	RemoveLast(name string) Handler
	Replace(old, new string, handler Handler) Handler
	First() Handler
	Last() Handler
	Get(name string) Handler
	Context(handler Handler) HandlerContext // by handler
	ContextN(name string) HandlerContext    // by name
	Channel() Channel
	Names() []string
	ToMap() map[string]Handler
	Flush() Pipeline
	FireChannelRegistered() Pipeline
	FireChannelUnregistered() Pipeline
	FireChannelActive() Pipeline
	FireChannelInactive() Pipeline
	FireExceptionCaught(ex exception.Exception) Pipeline
	FireUserEventTriggered(event Event) Pipeline
	FireChannelRead(message Message) Pipeline
	FireChannelReadComplete() Pipeline
	FireChannelWritabilityChanged() Pipeline
}

type pipeline struct {
	head            *defaultHandlerContext
	tail            *defaultHandlerContext
	channel         Channel
	succeededFuture Future
	childExecutors  map[string]concurrent.EventExecutor
	size            int
}

func (pl *pipeline) AddFirst(name string, handler Handler) Pipeline {
	return pl.AddFirstGroup(nil, name, handler)
}

func (pl *pipeline) AddFirstGroup(group concurrent.EventExecutorGroup, name string, handler Handler) Pipeline {
	name = pl.filterName(name)
	ctx := newContext(group, name, handler)
	pl.addFirst0(ctx)

	return pl
}

func (pl *pipeline) AddLast(name string, handler Handler) Pipeline {
	return pl.AddLastGroup(nil, name, handler)
}

func (pl *pipeline) AddLastGroup(group concurrent.EventExecutorGroup, name string, handler Handler) Pipeline {
	name = pl.filterName(name)
	ctx := newContext(group, name, handler)
	pl.addFirst0(ctx)

	return pl
}

func (pl *pipeline) AddBefore(base, name string, handler Handler) Pipeline {
	return nil
}

func (pl *pipeline) AddAfter(base, name string, handler Handler) Pipeline {
	return nil
}

func (pl *pipeline) Remove(name string) Handler {
	return nil
}

func (pl *pipeline) RemoveFirst(name string) Handler {
	return nil
}

func (pl *pipeline) RemoveLast(name string) Handler {
	return nil
}

func (pl *pipeline) Replace(old, new string, handler Handler) Handler {
	return nil
}

func (pl *pipeline) First() Handler {
	return nil
}

func (pl *pipeline) Last() Handler {
	return nil
}

func (pl *pipeline) Get(name string) Handler {
	return nil
}

func (pl *pipeline) Context(handler Handler) HandlerContext {
	return nil
}

func (pl *pipeline) ContextN(name string) HandlerContext {
	return nil
}

func (pl *pipeline) Channel() Channel {
	return nil
}

func (pl *pipeline) Names() []string {
	return nil
}

func (pl *pipeline) ToMap() map[string]Handler {
	return nil
}

func (pl *pipeline) Flush() Pipeline {
	return nil
}

func (pl *pipeline) FireChannelRegistered() Pipeline {
	return nil
}

func (pl *pipeline) FireChannelUnregistered() Pipeline {
	return nil
}

func (pl *pipeline) FireChannelActive() Pipeline {
	return nil
}

func (pl *pipeline) FireChannelInactive() Pipeline {
	return nil
}

func (pl *pipeline) FireExceptionCaught(ex exception.Exception) Pipeline {
	return nil
}

func (pl *pipeline) FireUserEventTriggered(event Event) Pipeline {
	return nil
}

func (pl *pipeline) FireChannelRead(message Message) Pipeline {
	return nil
}

func (pl *pipeline) FireChannelReadComplete() Pipeline {
	return nil
}

func (pl *pipeline) FireChannelWritabilityChanged() Pipeline {
	return nil
}

func (pl *pipeline) addFirst0(newCtx *defaultHandlerContext) {
	nextCtx := pl.head.next
	newCtx.prev = pl.head
	newCtx.next = nextCtx
	pl.head.next = newCtx
	nextCtx.prev = newCtx

	pl.size = pl.size + 1
}

func (pl *pipeline) addLast0(newCtx *defaultHandlerContext) {
	prev := pl.tail.prev
	newCtx.prev = prev
	newCtx.next = pl.tail
	prev.next = newCtx
	pl.tail.prev = newCtx

	pl.size = pl.size + 1
}

func (pl *pipeline) filterName(name string) string {
	return name
}

func NewPipeline(channel Channel) Pipeline {
	pl := &pipeline{
		channel:         channel,
		succeededFuture: NewSucceededFuture(channel),
		childExecutors:  make(map[string]concurrent.EventExecutor),
	}

	pl.head = newHeadContext(pl)
	pl.tail = newTailContext(pl)
	pl.head.next = pl.tail
	pl.tail.prev = pl.head

	pl.size = 2

	return pl
}
