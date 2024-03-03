package ui

import (
	"fmt"
	"log"
	"time"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/container"
	"fyne.io/fyne/widget"
	"github.com/VladimirKovalevus/go-music/internal/core"
)

func Init() {

	appCore := core.NewCore()
	a := app.New()

	w := a.NewWindow("Syndaudio")
	w.SetFixedSize(true)
	w.Resize(fyne.NewSize(400, 200))

	btnl := widget.NewButton("left", func() {
		appCore.PlaybackEvent(-5)
	})

	btnr := widget.NewButton("right", func() {
		appCore.PlaybackEvent(5)
	})

	btnStop := widget.NewButton("start/stop", func() {
		appCore.StartStopEvent()
	})
	btn1 := widget.NewButton("1", func() {
		appCore.ChangeTrackEvent("resources/White Shore - Enjoy the Motion.mp3")
	})
	btn2 := widget.NewButton("2", func() {
		appCore.ChangeTrackEvent("resources/White Shore - Your Gold.mp3")
	})
	btn3 := widget.NewButton("3", func() {
		appCore.ChangeTrackEvent("resources/syndafloden - мужская любовь.mp3")
	})
	progg := widget.NewSlider(float64(0), float64(100))
	progg.OnChanged = func(f float64) {
		appCore.Seek(f)
	}
	timelable := widget.NewLabel("123")
	progg.Step = 0.1

	list := widget.NewList(func() int {
		return appCore.PlaylistLen()
	},
		func() fyne.CanvasObject {
			return widget.NewLabel("____")
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			trck, err := appCore.GetPlaylistItem(i)
			if err != nil {
				log.Println(err)
			}
			o.(*widget.Label).SetText(trck.GetName())
		},
	)

	go func() {
		for {
			time.Sleep(time.Second / 10)
			progg.SetValue(appCore.PercentProgress() * 100)
			current, overall := appCore.TimeProgress()
			timelable.SetText(fmt.Sprintf("%02.f:%02d:%02d/%02.f:%02d:%02d",
				current.Hours(), int(current.Minutes())%60, int(current.Seconds())%60,
				overall.Hours(), int(overall.Minutes())%60, int(overall.Seconds())%60,
			))
		}
	}()

	w.SetContent(
		container.NewHBox(
			container.NewVBox(
				container.NewHBox(btnl, btnStop, btnr),
				container.NewHBox(btn1, btn2, btn3),
				progg, timelable),
			list,
		),
	)
	w.ShowAndRun()
}
