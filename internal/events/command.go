package events

import (
	"os"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/effects"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

type Command interface {
	Exec(e *EventLoop) error
}

type EXIT struct{}

func (EXIT) Exec(e *EventLoop) error {
	e.cancel()
	return nil
}

type PLAYBACK struct {
	Amount int32
}

func (p PLAYBACK) Exec(e *EventLoop) error {

	streamer := e.stream
	format := e.form
	streamer.Seek(streamer.Position() + format.SampleRate.N(time.Duration(p.Amount)*time.Second))
	return nil
}

type TRACK struct {
	Name string
}

func (t TRACK) Exec(e *EventLoop) error {

	f, err := os.Open(t.Name)
	if err != nil {
		return err
	}
	speaker.Lock()

	e.stream.Close()
	streamer, format, err := mp3.Decode(f)
	if err != nil {
		return err
	}
	e.stream = streamer
	e.form = format
	speaker.Unlock()
	speaker.Clear()
	speaker.Play(e.stream)
	return nil
}

type VOLUME struct {
	Amount int32
}

func (v VOLUME) Exec(e *EventLoop) error {

	// effects.Volume{}
	e.gain += float64(v.Amount)
	speaker.Lock()
	vol := &effects.Volume{Streamer: e.stream, Volume: e.gain / 100, Base: 10}
	ctrl := &beep.Ctrl{Streamer: vol, Paused: e.paused}
	speaker.Unlock()
	speaker.Clear()
	speaker.Play(ctrl)

	return nil
}

type StartStop struct {
}

func (v StartStop) Exec(e *EventLoop) error {

	e.paused = !e.paused
	ctrl := &beep.Ctrl{Streamer: e.stream, Paused: e.paused}
	speaker.Clear()
	speaker.Play(ctrl)
	return nil
}
