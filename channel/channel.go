package channel

import (
	`github.com/photowey/nettygo/pipeline`
)

type Config struct{}

type Id interface {
	ShortText() string
	LongText() string
}

type Channel interface {
	Id() Id
	Parent() Channel
	Config() Config
	Read() Channel
	Flush() Channel
	Pipeline() pipeline.ChannelPipeline
}

type channel struct {
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

func (c *channel) Pipeline() pipeline.ChannelPipeline {
	return c
}
