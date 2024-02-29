package playback

type Playlist struct {
	list     []Track
	length   int
	position int
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
