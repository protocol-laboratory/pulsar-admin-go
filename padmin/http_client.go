package padmin

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"
)

type HttpClient interface {
	Get(path string) (*http.Response, error)
	Put(path string, body any) (*http.Response, error)
	Post(path string, body any) (*http.Response, error)
	Delete(path string) (*http.Response, error)
	Do(*http.Request) (*http.Response, error)
}

type HttpClientImpl struct {
	urlPrefix string
	cli       *http.Client
}

func (h *HttpClientImpl) Get(path string) (*http.Response, error) {
	url := h.urlPrefix + path
	return h.cli.Get(url)
}

func (h *HttpClientImpl) Put(path string, body any) (*http.Response, error) {
	url := h.urlPrefix + path
	data, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	return h.cli.Do(req)
}

func (h *HttpClientImpl) Delete(path string) (*http.Response, error) {
	url := h.urlPrefix + path
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}
	return h.cli.Do(req)
}

func (h *HttpClientImpl) Post(path string, body any) (*http.Response, error) {
	url := h.urlPrefix + path
	data, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	return h.cli.Do(req)
}

func (h *HttpClientImpl) Do(req *http.Request) (*http.Response, error) {
	return h.cli.Do(req)
}

func newHttpClient(config Config) (HttpClient, error) {
	var cli *http.Client
	if config.TlsEnable {
		cli = &http.Client{
			Timeout: time.Millisecond * time.Duration(config.ConnectionTimeout),
			Transport: &http.Transport{
				TLSClientConfig: config.TlsConfig,
			},
		}
	} else {
		cli = &http.Client{
			Timeout: time.Millisecond * time.Duration(config.ConnectionTimeout),
		}
	}
	return &HttpClientImpl{
		urlPrefix: config.urlPrefix,
		cli:       cli,
	}, nil
}
