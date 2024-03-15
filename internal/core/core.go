package core

import (
	"context"

	"github.com/VladimirKovalevus/go-music/internal/core/events"
	"github.com/VladimirKovalevus/go-music/internal/core/playback"
	"github.com/faiface/beep"
)

type Core struct {
	eventloop   *events.EventLoop
	mainContext context.Context
	cancel      context.CancelFunc
}

func NewCore(e *events.EventLoop) *Core {
	return &Core{eventloop: e}
}

func (c *Core) Play(playlist playback.Playlist, index int) {
	c.cancel()
	c.mainContext, c.cancel = context.WithCancel(context.Background())
	go c.play(playlist, index)
}
func (c *Core) play(playlist playback.Playlist, index int) {
	queue := playlist.TracksFromIndex(index)
	play := make(chan struct{})
	for i := 0; i < len(queue) && c.mainContext.Err() == nil; i++ {
		track := queue[i]
		stream, _ := track.Stream()
		c.eventloop.Play(beep.Seq(stream, beep.Callback(func() {
			play <- struct{}{}
		})))
		<-play
	}
}
