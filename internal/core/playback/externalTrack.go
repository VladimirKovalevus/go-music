package playback

import (
	"fmt"
	"io"

	"github.com/faiface/beep"
)

type ExternalTrack struct {
	Metadata
	source string
	url    string
}

type Data struct {
	Buffer          []byte
	BlockSize       int64
	FetchedIndicies []bool
}

func (e *ExternalTrack) Source() (string, string) {
	return e.source, e.url
}

func (e *ExternalTrack) Reader() (io.ReadSeekCloser, error) {
	file, err := e.fetchUrl()
	return file, err
}
func (e *ExternalTrack) Stream() (beep.StreamSeekCloser, error) {
	file, err := e.fetchStreamUrl()
	return file, err
}

func (e *ExternalTrack) DumpString() string {
	return fmt.Sprintf("%s:%s", e.source, e.url)
}

func (e *ExternalTrack) fetchUrl() (io.ReadSeekCloser, error) {

	return nil, nil
}
func (e *ExternalTrack) fetchStreamUrl() (beep.StreamSeekCloser, error) {
	return nil, nil
}
