package ui

import (
	"fmt"
	"time"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/container"
	"fyne.io/fyne/widget"
	"github.com/VladimirKovalevus/go-music/internal/core/events"
)

func Init() {

	loop := events.NewEventLoop()
	loop.Play()
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
		loop.Seek(f)
	}
	timelable := widget.NewLabel("123")
	progg.Step = 0.1

	go func() {
		for {
			time.Sleep(time.Second / 10)
			progg.SetValue(loop.PercentProgress() * 100)
			current, overall := loop.TimeProgress()
			timelable.SetText(fmt.Sprintf("%02.f:%02d:%02d/%02.f:%02d:%02d",
				current.Hours(), int(current.Minutes())%60, int(current.Seconds())%60,
				overall.Hours(), int(overall.Minutes())%60, int(overall.Seconds())%60,
			))
		}
	}()

	w.SetContent(
		container.NewHBox(
			container.NewVBox(
				container.NewHBox(btnl, label, btnr),
				btnStop,
				container.NewHBox(btn1, btn2, btn3),
				progg, timelable),
		),
	)
	w.ShowAndRun()
}
