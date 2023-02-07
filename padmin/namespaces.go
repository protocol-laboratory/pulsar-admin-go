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

import (
	"encoding/json"
	"fmt"
	"io"
)

type Namespaces interface {
	Create(tenant, namespace string) error
	Delete(tenant, namespace string) error
	List(tenant string) ([]string, error)
}

type NamespacesImpl struct {
	cli HttpClient
}

func newNamespaces(cli HttpClient) *NamespacesImpl {
	return &NamespacesImpl{cli: cli}
}

func (n *NamespacesImpl) Create(tenant, namespace string) error {
	resp, err := n.cli.Put(fmt.Sprintf(UrlNamespacesFormat, tenant, namespace), nil)
	if err != nil {
		return err
	}
	return HttpCheck(resp)
}

func (n *NamespacesImpl) Delete(tenant, namespace string) error {
	resp, err := n.cli.Delete(fmt.Sprintf(UrlNamespacesFormat, tenant, namespace))
	if err != nil {
		return err
	}
	return HttpCheck(resp)
}

func (n *NamespacesImpl) List(tenant string) ([]string, error) {
	resp, err := n.cli.Get(fmt.Sprintf(UrlNamespacesFormat, tenant, ""))
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
	if err != nil {
		return nil, err
	}
	return result, nil
}
