package widgets

import (
	"fyne.io/fyne"
	"fyne.io/fyne/widget"
)

type PlaylistWidget struct {
	widget.BaseWidget
	title  *widget.Label
	button *widget.Button
}

func NewWidget() *PlaylistWidget {
	pw := &PlaylistWidget{title: widget.NewLabel("")}
	return pw
}

func (p *PlaylistWidget) Update(title string, icon string, callback func()) {
	resource, err := fyne.LoadResourceFromPath(icon)
	if err == nil {
		p.button.SetIcon(resource)
		p.button.OnTapped = callback
	}
	p.title.SetText(title)
}
