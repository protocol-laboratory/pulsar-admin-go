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

type NonPersistentTopics interface {
	Topics
}

type NonPersistentTopicsImpl struct {
	cli HttpClient
}

func newNonPersistentTopics(cli HttpClient) *NonPersistentTopicsImpl {
	return &NonPersistentTopicsImpl{cli: cli}
}

func (n *NonPersistentTopicsImpl) CreateNonPartitioned(tenant, namespace, topic string) error {
	path := fmt.Sprintf(UrlNonPersistentTopicFormat, tenant, namespace, topic)
	resp, err := n.cli.Put(path, nil)
	if err != nil {
		return err
	}
	return HttpCheck(resp)
}

func (n *NonPersistentTopicsImpl) DeleteNonPartitioned(tenant, namespace, topic string) error {
	path := fmt.Sprintf(UrlNonPersistentTopicFormat, tenant, namespace, topic)
	resp, err := n.cli.Delete(path)
	if err != nil {
		return err
	}
	return HttpCheck(resp)
}

func (n *NonPersistentTopicsImpl) ListNonPartitioned(tenant, namespace string) ([]string, error) {
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

func (n *NonPersistentTopicsImpl) CreatePartitioned(tenant, namespace, topic string, numPartitions int) error {
	path := fmt.Sprintf(UrlNonPersistentPartitionedTopicFormat, tenant, namespace, topic)
	resp, err := n.cli.Put(path, numPartitions)
	if err != nil {
		return err
	}
	return HttpCheck(resp)
}

func (n *NonPersistentTopicsImpl) ListPartitioned(tenant, namespace string) ([]string, error) {
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

func (n *NonPersistentTopicsImpl) DeletePartitioned(tenant, namespace, topic string) error {
	path := fmt.Sprintf(UrlNonPersistentPartitionedTopicFormat, tenant, namespace, topic)
	resp, err := n.cli.Delete(path)
	if err != nil {
		return err
	}
	return HttpCheck(resp)
}

func (n *NonPersistentTopicsImpl) ListNamespaceTopics(tenant, namespace string) ([]string, error) {
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

func (n *NonPersistentTopicsImpl) GetPartitionedMetadata(tenant, namespace, topic string) (*PartitionedMetadata, error) {
	url := fmt.Sprintf(UrlNonPersistentPartitionedTopicFormat, tenant, namespace, topic)
	resp, err := n.cli.Get(url)
	if err != nil {
		return nil, err
	}
	var metadata = new(PartitionedMetadata)
	if err := EasyReader(resp, metadata); err != nil {
		return nil, err
	}
	return metadata, nil
}

func (n *NonPersistentTopicsImpl) GetTopicRetention(tenant, namespace, topic string) (*RetentionConfiguration, error) {
	url := fmt.Sprintf(UrlNonPersistentPartitionedRetentionFormat, tenant, namespace, topic)
	resp, err := n.cli.Get(url)
	if err != nil {
		return nil, err
	}
	var retention = new(RetentionConfiguration)
	if err := EasyReader(resp, retention); err != nil {
		return nil, err
	}
	return retention, nil
}

func (n *NonPersistentTopicsImpl) SetTopicRetention(tenant, namespace, topic string, cfg *RetentionConfiguration) error {
	if cfg == nil {
		return fmt.Errorf("config empty")
	}
	url := fmt.Sprintf(UrlNonPersistentPartitionedRetentionFormat, tenant, namespace, topic)
	resp, err := n.cli.Post(url, cfg)
	if err != nil {
		return err
	}
	return HttpCheck(resp)
}

func (n *NonPersistentTopicsImpl) RemoveTopicRetention(tenant, namespace, topic string) error {
	url := fmt.Sprintf(UrlNonPersistentPartitionedRetentionFormat, tenant, namespace, topic)
	resp, err := n.cli.Delete(url)
	if err != nil {
		return err
	}
	return HttpCheck(resp)
}
