package events

import (
	"context"
	"fmt"
	"log"
	"math"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/speaker"
)

type EventLoop struct {
	commands chan Command
	ctx      context.Context
	cancel   context.CancelFunc
	gain     float64
	paused   bool
	form     beep.Format
	stream   beep.StreamSeekCloser
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

func (e *EventLoop) PercentProgress() float64 {
	return float64(e.stream.Position()) / float64(e.stream.Len())
}
func (e *EventLoop) TimeProgress() (time.Duration, time.Duration) {
	return e.form.SampleRate.D(e.stream.Position()), e.form.SampleRate.D(e.stream.Len())
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
	speaker.Lock()
	defer speaker.Unlock()
	newPos := int(float64(e.stream.Len()) * pos / 100)
	fmt.Println(newPos-e.stream.Position(), e.form.SampleRate.N(time.Second))
	if int(math.Abs(float64(newPos-e.stream.Position()))) < e.form.SampleRate.N(time.Second) {
		return
	}
	e.stream.Seek(e.form.SampleRate.N(e.form.SampleRate.D(newPos)))
}
func (e *EventLoop) StartStopEvent() {
	e.commands <- StartStop{}
}
func (e *EventLoop) ChangeTrackEvent(file string) {
	e.commands <- CHANGE_TRACK{Name: file}
}
func (e *EventLoop) Play() {
	speaker.Play(e.stream)
}
