package playback

import (
	"io"
	"os"

	"github.com/dhowden/tag"
	"github.com/faiface/beep"
)

type Track interface {
	GetName() string
	GetAlbum() string
	GetArtist() string
	GetIcon() *tag.Picture
	Load() (beep.StreamSeekCloser, chan struct{})
}

type Metadata struct {
	Title  string
	Album  string
	Artist string
	Icon   *tag.Picture
}

func (m Metadata) GetName() string {
	return m.Title
}
func (m Metadata) GetAlbum() string {
	return m.Album
}
func (m Metadata) GetArtist() string {
	return m.Artist
}
func (m Metadata) GetIcon() *tag.Picture {
	return m.Icon
}

type LocalTrack struct {
	Metadata
	filePath string
}

func NewLocalTrack(filePath string) *LocalTrack {
	myMeta := Metadata{}
	f, _ := os.Open(filePath)
	meta, err := tag.ReadFrom(f)
	f.Close()
	if err != nil {
		myMeta.Icon = meta.Picture()
		myMeta.Album = meta.Album()
		myMeta.Artist = meta.Artist()
		myMeta.Title = meta.Title()
	}
	return &LocalTrack{Metadata: myMeta, filePath: filePath}
}

func (l *LocalTrack) GetReader() (io.ReadCloser, error) {
	file, err := os.Open(l.filePath)
	return file, err
}
