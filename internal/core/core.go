package core

import (
	"github.com/VladimirKovalevus/go-music/internal/core/events"
	"github.com/VladimirKovalevus/go-music/internal/core/playback"
)

type Core struct {
	e         *events.EventLoop
	playlists []playback.Playlist
}

func NewCore() *Core {
	return &Core{e: events.NewEventLoop()}
}
