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

var _ Channel = (*channel)(nil)

type Channel interface {
	Id() Id
	Parent() Channel
	Config() Config
	Read() Channel
	Flush() Channel
	Pipeline() Pipeline
}

type channel struct {
	pipeline Pipeline
}

func (c *channel) Id() Id {
	return nil
}

func (c *channel) Parent() Channel {
	return nil
}

func (c *channel) Config() Config {
	return Config{}
}

func (c *channel) Read() Channel {
	return c
}

func (c *channel) Flush() Channel {
	return c
}

func (c *channel) Pipeline() Pipeline {
	return c.pipeline
}
