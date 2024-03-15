package ui

import (
	"fyne.io/fyne"
	"github.com/VladimirKovalevus/go-music/internal/core"
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
	Core          core.Core
	Playlist      []playback.Playlist
}

func Init() {
	// ui := UI{EventLoop: events.NewEventLoop(), Playlist: playback.Parse()}
	// widget.NewList()
}
