package playback

import (
	"fmt"
	"io"
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

func (e *ExternalTrack) DumpString() string {
	return fmt.Sprintf("%s:%s", e.source, e.url)
}

func (e *ExternalTrack) fetchUrl() (io.ReadSeekCloser, error) {
	//TODO

	return nil, nil
}
