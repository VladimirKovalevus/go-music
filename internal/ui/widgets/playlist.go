package widgets

import (
	"fyne.io/fyne"
	"fyne.io/fyne/widget"
)

type PlaylistWidget struct {
	widget.BaseWidget
	title *widget.Label
	icon  *widget.Icon
}

func NewWidget() *PlaylistWidget {
	pw := &PlaylistWidget{title: widget.NewLabel("")}
	return pw
}

func (p *PlaylistWidget) Update(title string, icon string) {
	resource, err := fyne.LoadResourceFromPath(icon)
	if err == nil {
		p.icon.SetResource(resource)
	}
	p.title.SetText(title)
}
