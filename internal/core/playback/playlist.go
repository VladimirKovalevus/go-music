package playback

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const CONF_PATH = "."

type Playlist struct {
	list     []Track
	length   int
	position int
	Name     string
}

func (p *Playlist) Next() bool {
	return p.position < p.length-1
}
func (p *Playlist) Add(tracks ...Track) {
	p.list = append(p.list, tracks...)
}
func (p *Playlist) Remove(ind int) {
	if ind < len(p.list) && ind >= 0 {
		newlist := []Track{}
		for i := range p.list {
			if i == ind {
				continue
			}
			newlist = append(newlist, p.list[i])
		}
		p.list = newlist
	}
}
func (p *Playlist) PlaylistLen() int {
	if p == nil {
		return 0
	}
	return p.length
}
func (p *Playlist) GetPlaylistItem(i int) (Track, error) {
	if p == nil {
		return nil, fmt.Errorf("empty playlist")
	}
	if i > p.length || i < 0 {
		return nil, fmt.Errorf("out of bounds")
	}
	return p.list[i], nil
}

func (p *Playlist) DumpFile() error {
	f, err := os.Create(CONF_PATH + p.Name + ".spl")
	if err != nil {
		return err
	}
	for _, track := range p.list {
		f.WriteString(track.DumpString())
	}
	return f.Close()
}

func Parse() []*Playlist {
	var play []*Playlist
	dirs, _ := os.ReadDir(CONF_PATH)
	for _, dir := range dirs {
		if strings.HasSuffix(dir.Name(), ".spl") {
			playlist := parse(dir.Name())
			play = append(play, &playlist)
		}
	}
	fmt.Printf("%#v", play[0])
	return play
}

func parse(dir string) Playlist {
	var list []Track
	f, _ := os.Open(dir)
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		var source string
		var url string
		split := strings.Split(scanner.Text(), ":")
		source = split[0]
		url = split[1]
		list = append(list, prepareTrack(source, url))
	}
	return Playlist{Name: dir, list: list}
}

func prepareTrack(source, url string) Track {
	switch source {
	case "OS":
		return NewLocalTrack(url)
	}
	return nil
}
