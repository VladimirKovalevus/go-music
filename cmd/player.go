package main

import (
	"log"
	"os"
	"time"

	"github.com/VladimirKovalevus/go-music/internal/events"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

func main() {
	f, err := os.Open("resources/syndafloden - мужская любовь.mp3")
	if err != nil {
		log.Fatalln(err)
	}
	streamer, format, err := mp3.Decode(f)
	if err != nil {
		log.Fatalln(err)
	}

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

	loop := events.NewEventLoop(streamer, format)
	loop.Play()
	loop.DispatchEvents()
}
