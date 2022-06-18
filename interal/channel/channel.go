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
	"context"
)

var _ Channel = (*channel)(nil)

type Channel interface {
	Id() Id
	Parent() Channel
	Config() Config
	Read() Channel
	Flush() Channel
	Pipeline() Pipeline
	Close(err error)
	IsActive() bool
	State() int32
	Attachment() Attachment
	SetAttachment(attr Attachment)
	Context() context.Context
}

type channel struct {
	id         Id
	ctx        context.Context
	cancel     context.CancelFunc
	pipeline   Pipeline
	attachment Attachment
	state      int32 // high-low 0xxx0000 00000000 00000000 00000000 ?
}

func (ch *channel) Id() Id {
	return ch.id
}

func (ch *channel) Parent() Channel {
	return ch
}

func (ch *channel) Config() Config {
	return Config{}
}

func (ch *channel) Read() Channel {
	return ch
}

func (ch *channel) Flush() Channel {
	return ch
}

func (ch *channel) Pipeline() Pipeline {
	return ch.pipeline
}

func (ch *channel) Close(err error) { // Exception ?
	// TODO
}

func (ch *channel) IsActive() bool {
	// TODO
	return false
}

func (ch *channel) State() int32 {
	return ch.state
}

func (ch *channel) Attachment() Attachment {
	return ch.attachment
}

func (ch *channel) SetAttachment(attr Attachment) {
	ch.attachment = attr
}

func (ch *channel) Context() context.Context {
	return ch.ctx
}

func NewChannel() Channel {
	ch := &channel{
		id: NewChannelId(),
	}
	ch.pipeline = NewPipeline(ch)

	return ch
}
