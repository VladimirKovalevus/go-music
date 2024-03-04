package core

import (
	"fmt"

	"github.com/VladimirKovalevus/go-music/internal/core/events"
	"github.com/VladimirKovalevus/go-music/internal/core/playback"
)

type Core struct {
	*events.EventLoop
	*playback.Playlist
	Playlists []*playback.Playlist
}

func NewCore() *Core {
	playlists := playback.Parse()
	return &Core{EventLoop: events.NewEventLoop(), Playlists: playlists, Playlist: playlists[0]}
}
func (c *Core) ShowPlaylist() {
	fmt.Println(c.Playlist)
}

func (c *Core) SetCurrentPlaylist(i int) {
	c.Playlist = c.Playlists[i]
}

// func (c *Core)
