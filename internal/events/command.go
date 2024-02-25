package events

import (
	"os"
	"time"

	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

type Command interface {
	Exec(e *EventLoop) error
}

type EXIT struct{}

func (EXIT) Exec(e *EventLoop) error {
	os.Exit(0)
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
