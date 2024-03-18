package ui

import (
	"fmt"

	"fyne.io/fyne"
	"github.com/VladimirKovalevus/go-music/internal/core"
	"github.com/VladimirKovalevus/go-music/internal/core/playback"
)

type UI struct {
	Core            *core.Core
	TrackList       fyne.CanvasObject
	Playlists       fyne.CanvasObject
	stopStart       fyne.CanvasObject
	previousTrack   fyne.CanvasObject
	nextTrack       fyne.CanvasObject
	TrackInfo       fyne.CanvasObject
	Playlist        []playback.Playlist
	CurrentPlaylist uint32
}

func (u *UI) SelectPlaylist() {
}
func (u *UI) ReloadTracks() {
}
func (u *UI) NewPlayList(name string, icon []byte) {
}
func (u *UI) AddTrackToPlaylist(p *playback.Playlist, t *playback.Track) {
}
func (u *UI) NextTrack() {
	fmt.Println(u.nextTrack)
}
func (u *UI) PreviousTrack() {
	fmt.Println(u.previousTrack)
}
func (u *UI) StopStart() {
	fmt.Println(u.stopStart)
}

func NewUi() UI {
	// widget.NewList()
	return UI{}
}
