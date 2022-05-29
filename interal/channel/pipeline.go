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

var _ Pipeline = (*pipeline)(nil)

type Pipeline interface {
	AddFirst(name string, handler Handler) Pipeline
	AddLast(name string, handler Handler) Pipeline
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
}

type pipeline struct {
	channel Channel
}

func (pl *pipeline) AddFirst(name string, handler Handler) Pipeline {
	return nil
}

func (pl *pipeline) AddLast(name string, handler Handler) Pipeline {
	return nil
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

func NewPipeline(channel Channel) Pipeline {
	return &pipeline{
		channel: channel,
	}
}
