package ui

import (
	"fyne.io/fyne"
	"fyne.io/fyne/widget"
	"github.com/VladimirKovalevus/go-music/internal/core/events"
	"github.com/VladimirKovalevus/go-music/internal/core/playback"
)

type UI_CONTROLLER interface {
	AddTrack(t playback.Track, p *playback.Playlist) bool
	RemoveTrack(t playback.Track, p *playback.Playlist) bool
	NewPlaylist(name string) bool
}

type UI struct {
	TrackList     fyne.CanvasObject
	Playlists     fyne.CanvasObject
	StopStart     fyne.CanvasObject
	PreviousTrack fyne.CanvasObject
	NextTrack     fyne.CanvasObject
	TrackInfo     fyne.CanvasObject
	EventLoop     *events.EventLoop
	Playlist      []playback.Playlist
}

func Init() {
	ui := UI{EventLoop: events.NewEventLoop(), Playlist: playback.Parse()}
	widget.NewList()
}
