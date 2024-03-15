package ui

import (
	"fyne.io/fyne"
	"fyne.io/fyne/container"
	"fyne.io/fyne/widget"
	"github.com/VladimirKovalevus/go-music/internal/core/events"
	"github.com/VladimirKovalevus/go-music/internal/core/playback"
)

func NewTrackLayout(t playback.Track, e *events.EventLoop) fyne.CanvasObject {
	btn := widget.NewButton(t.Title(), func() {
	})
	return container.NewHBox(btn)
}
