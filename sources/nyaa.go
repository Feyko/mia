package sources

import "io"

type Nyaa struct {
}

func (n *Nyaa) Name() string {
	return "Nyaa"
}

func (n *Nyaa) Search(match string) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (n *Nyaa) Download(match string) (io.ReadCloser, error) {
	//TODO implement me
	panic("implement me")
}
