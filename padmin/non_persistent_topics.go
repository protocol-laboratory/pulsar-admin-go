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
	"errors"
	"fmt"
	"io"
)

type NonPersistentTopics struct {
	cli HttpClient
}

func newNonPersistentTopics(cli HttpClient) *NonPersistentTopics {
	return &NonPersistentTopics{cli: cli}
}

func (n *NonPersistentTopics) CreateNonPartitioned(tenant, namespace, topic string) error {
	path := fmt.Sprintf(UrlNonPersistentTopicFormat, tenant, namespace, topic)
	resp, err := n.cli.Put(path, nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if !StatusOk(resp.StatusCode) {
		str, err := ReadAll(resp.Body)
		if err != nil {
			return err
		}
		return errors.New(str)
	}
	return nil
}

func (n *NonPersistentTopics) DeleteNonPartitioned(tenant, namespace, topic string) error {
	path := fmt.Sprintf(UrlNonPersistentTopicFormat, tenant, namespace, topic)
	resp, err := n.cli.Delete(path)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if !StatusOk(resp.StatusCode) {
		str, err := ReadAll(resp.Body)
		if err != nil {
			return err
		}
		return errors.New(str)
	}
	return nil
}

func (n *NonPersistentTopics) ListNonPartitioned(tenant, namespace string) ([]string, error) {
	path := fmt.Sprintf(UrlNonPersistentNamespaceFormat, tenant, namespace)
	resp, err := n.cli.Get(path)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if !StatusOk(resp.StatusCode) {
		str, err := ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		return nil, errors.New(str)
	}
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	topics := make([]string, 0)
	err = json.Unmarshal(data, &topics)
	if err != nil {
		return nil, err
	}
	return topics, nil
}
