package ui

import (
	"fyne.io/fyne"
	"github.com/VladimirKovalevus/go-music/internal/core"
	"github.com/VladimirKovalevus/go-music/internal/core/playback"
)

type UI struct {
	TrackList       fyne.CanvasObject
	Playlists       fyne.CanvasObject
	StopStart       fyne.CanvasObject
	PreviousTrack   fyne.CanvasObject
	NextTrack       fyne.CanvasObject
	TrackInfo       fyne.CanvasObject
	Core            core.Core
	Playlist        []playback.Playlist
	CurrentPlaylist uint32
}

func (u *UI) SelectPlaylist() {
}

func (u *UI) ReloadTracks() {

}

func Init() UI {
	return UI{}
}
