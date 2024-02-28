package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
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
	// //////////////////////////////
	a := app.New()

	w := a.NewWindow("Syndaudio")
	w.SetFixedSize(true)
	w.Resize(fyne.NewSize(400, 200))
	label := widget.NewLabel("some text")

	btnl := widget.NewButton("left", func() {
		loop.PlaybackEvent(-5)
	})

	btnr := widget.NewButton("right", func() {
		loop.PlaybackEvent(5)
	})

	btnStop := widget.NewButton("start/stop", func() {
		loop.StartStopEvent()
	})

	btn1 := widget.NewButton("1", func() {
		loop.ChangeTrackEvent("resources/White Shore - Enjoy the Motion.mp3")
	})
	btn2 := widget.NewButton("2", func() {
		loop.ChangeTrackEvent("resources/White Shore - Your Gold.mp3")
	})
	btn3 := widget.NewButton("3", func() {
		loop.ChangeTrackEvent("resources/syndafloden - мужская любовь.mp3")
	})
	progg := widget.NewSlider(float64(0), float64(100))
	progg.OnChanged = func(f float64) {
		loop.Seek(f / 100)
		fmt.Println(f / 100)
	}

	go func() {
		for {
			fmt.Println(loop.PercentProgress())
			time.Sleep(time.Second / 10)
			progg.SetValue(loop.PercentProgress() * 100)
		}
	}()

	w.SetContent(
		container.NewVBox(
			container.NewHBox(btnl, label, btnr),
			btnStop,
			container.NewHBox(btn1, btn2, btn3),
			progg),
	)
	w.ShowAndRun()
}
