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
	stream   beep.StreamSeekCloser
	form     beep.Format
}

func NewEventLoop(strm beep.StreamSeekCloser, frm beep.Format) *EventLoop {
	loop := &EventLoop{}
	loop.commands = make(chan Command, 1)
	loop.ctx = context.Background()
	loop.stream = strm
	loop.form = frm
	go func() {
		for {
			com := <-loop.commands
			switch v := com.(type) {
			case PLAYBACK:
				{
					log.Println(v.Exec(loop))
				}
			case EXIT:
				{
					log.Println(v.Exec(loop))
				}
			case TRACK:
				{
					log.Println(v.Exec(loop))
				}
			}
		}
	}()
	return loop
}
func (e *EventLoop) DispatchEvents() {
	for {
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
		}
	}
}
func (e *EventLoop) Play() {
	speaker.Play(e.stream)
}
