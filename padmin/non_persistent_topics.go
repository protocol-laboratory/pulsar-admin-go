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
)

type INonPersistentTopics interface {
	ITopics
}

type nonPersistentTopics struct {
	cli HttpClient
}

func newNonPersistentTopics(cli HttpClient) *nonPersistentTopics {
	return &nonPersistentTopics{cli: cli}
}

func (n *nonPersistentTopics) CreateNonPartitioned(tenant, namespace, topic string) error {
	path := fmt.Sprintf(UrlNonPersistentTopicFormat, tenant, namespace, topic)
	resp, err := n.cli.Put(path, nil)
	if err != nil {
		return err
	}
	return HttpCheck(resp)
}

func (n *nonPersistentTopics) DeleteNonPartitioned(tenant, namespace, topic string) error {
	path := fmt.Sprintf(UrlNonPersistentTopicFormat, tenant, namespace, topic)
	resp, err := n.cli.Delete(path)
	if err != nil {
		return err
	}
	return HttpCheck(resp)
}

func (n *nonPersistentTopics) ListNonPartitioned(tenant, namespace string) ([]string, error) {
	path := fmt.Sprintf(UrlNonPersistentNamespaceFormat, tenant, namespace)
	resp, err := n.cli.Get(path)
	if err != nil {
		return nil, err
	}
	data, err := HttpCheckReadBytes(resp)
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

func (n *nonPersistentTopics) CreatePartitioned(tenant, namespace, topic string, numPartitions int) error {
	path := fmt.Sprintf(UrlNonPersistentPartitionedTopicFormat, tenant, namespace, topic)
	resp, err := n.cli.Put(path, numPartitions)
	if err != nil {
		return err
	}
	return HttpCheck(resp)
}

func (n *nonPersistentTopics) ListPartitioned(tenant, namespace string) ([]string, error) {
	path := fmt.Sprintf(UrlNonPersistentPartitionedNamespaceFormat, tenant, namespace)
	resp, err := n.cli.Get(path)
	if err != nil {
		return nil, err
	}
	var topics []string
	if err := EasyReader(resp, &topics); err != nil {
		return nil, err
	}
	return topics, nil
}

func (n *nonPersistentTopics) DeletePartitioned(tenant, namespace, topic string) error {
	path := fmt.Sprintf(UrlNonPersistentPartitionedTopicFormat, tenant, namespace, topic)
	resp, err := n.cli.Delete(path)
	if err != nil {
		return err
	}
	return HttpCheck(resp)
}

func (n *nonPersistentTopics) ListNamespaceTopics(tenant, namespace string) ([]string, error) {
	url := fmt.Sprintf(UrlNonPersistentNamespaceFormat, tenant, namespace)
	resp, err := n.cli.Get(url)
	if err != nil {
		return nil, err
	}
	var topics []string
	if err := EasyReader(resp, &topics); err != nil {
		return nil, err
	}
	return topics, nil
}
