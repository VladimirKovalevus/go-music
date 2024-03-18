package ui

import (
	"fmt"

	"fyne.io/fyne"
	"fyne.io/fyne/widget"
	"github.com/VladimirKovalevus/go-music/internal/core"
	"github.com/VladimirKovalevus/go-music/internal/core/playback"
	"github.com/VladimirKovalevus/go-music/internal/ui/widgets"
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
func (u *UI) PlaylistCount() int {
	return len(u.Playlist)
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

func NewUi(c *core.Core) *UI {
	ui := &UI{}
	playlists := widget.NewList(func() int {
		return ui.PlaylistCount()
	}, func() fyne.CanvasObject {
		return widgets.NewWidget()
	}, func(lii widget.ListItemID, co fyne.CanvasObject) {
		play := ui.Playlist[lii]
		text, icon := play.Name, play.IconPath
		co.(*widgets.PlaylistWidget).Update(text, icon)
	})
	ui.Playlists = playlists
	return ui
}
