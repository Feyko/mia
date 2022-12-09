package sources

import "io"

type ThePirateBay struct {
}

func (t *ThePirateBay) Name() string {
	return "ThePirateBay"
}

func (t *ThePirateBay) Search(match string) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (t *ThePirateBay) Download(match string) (io.ReadCloser, error) {
	//TODO implement me
	panic("implement me")
}
