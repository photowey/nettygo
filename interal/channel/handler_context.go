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
}

type HeadContext struct {
	DefaultHandlerContext
}

type TailContext struct {
	DefaultHandlerContext
}

type DefaultHandlerContext struct {
	prev            *DefaultHandlerContext
	next            *DefaultHandlerContext
	name            string
	handler         Handler
	executor        concurrent.EventExecutor
	pipeline        Pipeline
	succeededFuture Future
	handlerState    int
}

func (hc *DefaultHandlerContext) Channel() Channel {
	return hc.pipeline.Channel()
}

func (hc *DefaultHandlerContext) Handler() Handler {
	return nil
}

func (hc *DefaultHandlerContext) Executor() concurrent.EventExecutor {
	return nil
}

func newHeadContext(pl Pipeline) *DefaultHandlerContext {
	// TODO
	return &DefaultHandlerContext{
		pipeline:     pl,
		handlerState: Init,
	}
}

func newTailContext(pl Pipeline) *DefaultHandlerContext {
	// TODO
	return &DefaultHandlerContext{
		pipeline:     pl,
		handlerState: Init,
	}
}

func newContext(group concurrent.EventExecutorGroup, name string, handler Handler) *DefaultHandlerContext {
	executor := childExecutor(group)

	ctx := &DefaultHandlerContext{
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
