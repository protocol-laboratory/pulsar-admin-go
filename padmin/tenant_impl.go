package padmin

import (
	"encoding/json"
	"io"
)

type TenantsImpl struct {
	cli HttpClient
}

func newTenants(cli HttpClient) *TenantsImpl {
	return &TenantsImpl{cli: cli}
}

func (t *TenantsImpl) Create(tenantName string, info TenantInfo) error {
	path := UrlTenants + "/" + tenantName
	resp, err := t.cli.Put(path, info)
	if err != nil {
		return err
	}
	return HttpCheck(resp)
}

func (t *TenantsImpl) Delete(tenantName string) error {
	url := UrlTenants + "/" + tenantName
	resp, err := t.cli.Delete(url)
	if err != nil {
		return err
	}
	return HttpCheck(resp)
}

func (t *TenantsImpl) List() ([]string, error) {
	resp, err := t.cli.Get(UrlTenants)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	result := make([]string, 0)
	err = json.Unmarshal(bytes, &result)
	return result, err
}
