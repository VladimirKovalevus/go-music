package playback

import (
	"io"
	"log"
	"os"
	"strings"

	"github.com/dhowden/tag"
	"github.com/faiface/beep"
)

type Track interface {
	Title() string
	Album() string
	Artist() string
	Icon() *tag.Picture
	Reader() (io.ReadCloser, error)
	Stream() (beep.StreamSeekCloser, error)
	DumpString() string
	Source() (string, string)
}

type Metadata struct {
	title  string
	album  string
	artist string
	icon   *tag.Picture
}

func (m Metadata) Title() string {
	return m.title
}
func (m Metadata) Album() string {
	return m.album
}
func (m Metadata) Artist() string {
	return m.artist
}
func (m Metadata) Icon() *tag.Picture {
	return m.icon
}

type LocalTrack struct {
	Metadata
	filePath string
}

func NewLocalTrack(filePath string) *LocalTrack {
	myMeta := Metadata{}
	f, err := os.Open(filePath)
	if err != nil {
		log.Println(err)
		return nil
	}
	meta, err := tag.ReadFrom(f)

	f.Close()

	if err == nil && meta != nil {

		myMeta.icon = meta.Picture()
		myMeta.album = meta.Album()
		myMeta.artist = meta.Artist()
		myMeta.title = meta.Title()
	} else {
		ind := strings.LastIndex(filePath, "/")
		myMeta.title = filePath[ind+1:]
		myMeta.artist = "Download"
		myMeta.album = "Unknown"
	}
	return &LocalTrack{Metadata: myMeta, filePath: filePath}
}

func (l *LocalTrack) Reader() (io.ReadCloser, error) {
	file, err := os.Open(l.filePath)
	return file, err
}
func (l *LocalTrack) Stream() (beep.StreamSeekCloser, error) {
	return nil, nil
}
func (l *LocalTrack) DumpString() string {
	return "OS:" + l.filePath
}
func (l *LocalTrack) Source() (string, string) {
	return "OS", l.filePath
}
