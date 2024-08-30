package padmin

import "fmt"

type BookiesImpl struct {
	cli HttpClient
}

func (b *BookiesImpl) AllBookies() (*AllBookiesResp, error) {
	resp, err := b.cli.Get(UrlBookiesAll)
	if err != nil {
		return nil, err
	}
	var body = new(AllBookiesResp)
	if err := EasyReader(resp, body); err != nil {
		return nil, err
	}
	return body, nil
}

func (b *BookiesImpl) ListRacksInfo() (*ListRacksInfoResp, error) {
	resp, err := b.cli.Get(UrlBookiesRacksInfo)
	if err != nil {
		return nil, err
	}
	var body = new(ListRacksInfoResp)
	if err := EasyReader(resp, body); err != nil {
		return nil, err
	}
	return body, nil
}

// RemoveRacksInfo bookie is address:port
func (b *BookiesImpl) RemoveRacksInfo(bookie string) error {
	url := fmt.Sprintf(UrlBookiesRacksFormat, bookie)
	resp, err := b.cli.Delete(url)
	if err != nil {
		return err
	}
	return HttpCheck(resp)
}

func (b *BookiesImpl) GetRacksInfo(bookie string) (*BookieInfo, error) {
	url := fmt.Sprintf(UrlBookiesRacksFormat, bookie)
	resp, err := b.cli.Get(url)
	if err != nil {
		return nil, err
	}
	var body = new(BookieInfo)
	if err := EasyReader(resp, body); err != nil {
		return nil, err
	}
	return body, nil
}

func (b *BookiesImpl) UpdateRacksInfo(bookie string, info *BookieInfo) error {
	url := fmt.Sprintf(UrlBookiesRacksFormat, bookie)
	resp, err := b.cli.Post(url, info)
	if err != nil {
		return err
	}
	return HttpCheck(resp)
}

func newBookies(cli HttpClient) Bookies {
	return &BookiesImpl{cli: cli}
}
