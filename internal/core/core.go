package core

import (
	"github.com/VladimirKovalevus/go-music/internal/core/events"
	"github.com/VladimirKovalevus/go-music/internal/core/killswitch"
	"github.com/VladimirKovalevus/go-music/internal/core/playback"
	"github.com/faiface/beep"
)

type Core struct {
	eventloop  *events.EventLoop
	killswitch *killswitch.Killswitch
}

func NewCore(e *events.EventLoop) *Core {
	return &Core{eventloop: e}
}

func (c *Core) Play(playlist playback.Playlist, index int) {
	c.killswitch.Cancel()
	c.killswitch = killswitch.NewKillswitch()
	go c.play(playlist, index)
}
func (c *Core) play(playlist playback.Playlist, index int) {
	queue := playlist.TracksFromIndex(index)
	play := make(chan struct{})
	for i := 0; i < len(queue) && c.killswitch.Err() == nil; {
		stream, _ := queue[i].Stream()
		c.eventloop.Play(beep.Seq(stream, beep.Callback(func() {
			play <- struct{}{}
		})))
		select {
		case <-play:
			{
				i++
			}
		case <-c.killswitch.Done():
			{
				return
			}
		}
	}
	c.killswitch.Cancel()
}
