// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

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
