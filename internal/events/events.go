package events

import (
	"context"
	"fmt"
	"log"

	"github.com/faiface/beep"
	"github.com/faiface/beep/speaker"
)

type EventLoop struct {
	commands chan Command
	ctx      context.Context
	cancel   context.CancelFunc
	stream   beep.StreamSeekCloser
	gain     float64
	paused   bool
	form     beep.Format
}

func NewEventLoop(strm beep.StreamSeekCloser, frm beep.Format) *EventLoop {
	loop := &EventLoop{}
	loop.commands = make(chan Command, 1)
	loop.ctx, loop.cancel = context.WithCancel(context.Background())
	loop.stream = strm
	loop.form = frm
	go func() {
		for {
			if e := loop.ctx.Err(); e != nil {
				log.Println(e)
				return
			}
			com := <-loop.commands
			com.Exec(loop)
		}
	}()
	return loop
}
func (e *EventLoop) DispatchEvents() {
	for {
		if e := e.ctx.Err(); e != nil {
			log.Println(e)
			return
		}
		var input string
		fmt.Scan(&input)
		fmt.Println(input[0])
		switch input[0] {
		case 97:
			e.commands <- PLAYBACK{Amount: 5}
		case 100:
			e.commands <- PLAYBACK{Amount: -5}
		case 49:
			e.commands <- TRACK{Name: "resources/syndafloden - мужская любовь.mp3"}
		case 50:
			e.commands <- TRACK{Name: "resources/White Shore - Enjoy the Motion.mp3"}
		case 51:
			e.commands <- TRACK{Name: "resources/White Shore - Your Gold.mp3"}
		case 119:
			e.commands <- VOLUME{5}
		case 115:
			e.commands <- VOLUME{-5}
		case 112:
			e.commands <- EXIT{}
			return
		case 113:
			e.commands <- StartStop{}
		}
	}
}

func (e *EventLoop) PercentProgress() float64 {
	return float64(e.stream.Position()) / float64(e.stream.Len())
}
func (e *EventLoop) TimeProgress() float64 {
	return float64(e.stream.Len()) / float64(e.stream.Position())
}

func (e *EventLoop) VolumeEvent(Amount int) {
	e.commands <- VOLUME{Amount: int32(Amount)}
}

func (e *EventLoop) PlaybackEvent(Amount int) {
	e.commands <- PLAYBACK{Amount: int32(Amount)}
}
func (e *EventLoop) ExitEvent() {
	e.commands <- EXIT{}
}
func (e *EventLoop) Seek(pos float64) {
	newPos := int(float64(e.stream.Len()) * pos)
	e.stream.Seek(newPos)
}
func (e *EventLoop) StartStopEvent() {
	e.commands <- StartStop{}
}
func (e *EventLoop) ChangeTrackEvent(file string) {
	e.commands <- TRACK{Name: file}

}
func (e *EventLoop) Play() {
	speaker.Play(e.stream)
}
