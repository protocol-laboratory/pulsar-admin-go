package padmin

import (
	"encoding/json"
	"io"
)

type Clusters interface {
	List() ([]string, error)
}

type ClustersImpl struct {
	cli HttpClient
}

func newClusters(cli HttpClient) *ClustersImpl {
	return &ClustersImpl{cli: cli}
}

func (c *ClustersImpl) List() ([]string, error) {
	resp, err := c.cli.Get(UrlClusters)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var data []byte
	data, err = io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	result := make([]string, 0)
	err = json.Unmarshal(data, &result)
	return result, err
}
