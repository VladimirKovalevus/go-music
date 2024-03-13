package core

import (
	"github.com/VladimirKovalevus/go-music/internal/core/events"
)

type Core struct {
	*events.EventLoop
}

func NewCore() *Core {
	return &Core{EventLoop: events.NewEventLoop()}
}
