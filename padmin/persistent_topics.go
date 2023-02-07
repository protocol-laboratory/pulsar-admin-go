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

type PersistentTopics interface {
	Topics
}

type PersistentTopicsImpl struct {
	cli HttpClient
}

func newPersistentTopics(cli HttpClient) *PersistentTopicsImpl {
	return &PersistentTopicsImpl{cli: cli}
}

func (p *PersistentTopicsImpl) CreateNonPartitioned(tenant, namespace, topic string) error {
	path := fmt.Sprintf(UrlPersistentTopicFormat, tenant, namespace, topic)
	resp, err := p.cli.Put(path, nil)
	if err != nil {
		return err
	}
	return HttpCheck(resp)
}

func (p *PersistentTopicsImpl) DeleteNonPartitioned(tenant, namespace, topic string) error {
	path := fmt.Sprintf(UrlPersistentTopicFormat, tenant, namespace, topic)
	resp, err := p.cli.Delete(path)
	if err != nil {
		return err
	}
	return HttpCheck(resp)
}

func (p *PersistentTopicsImpl) ListNonPartitioned(tenant, namespace string) ([]string, error) {
	path := fmt.Sprintf(UrlPersistentNamespaceFormat, tenant, namespace)
	resp, err := p.cli.Get(path)
	if err != nil {
		return nil, err
	}
	data, err := HttpCheckReadBytes(resp)
	if err != nil {
		return nil, err
	}
	var topics []string
	if err := json.Unmarshal(data, &topics); err != nil {
		return nil, err
	}
	return topics, nil
}

func (p *PersistentTopicsImpl) CreatePartitioned(tenant, namespace, topic string, numPartitions int) error {
	path := fmt.Sprintf(UrlPersistentPartitionedTopicFormat, tenant, namespace, topic)
	resp, err := p.cli.Put(path, numPartitions)
	if err != nil {
		return err
	}
	return HttpCheck(resp)
}

func (p *PersistentTopicsImpl) DeletePartitioned(tenant, namespace, topic string) error {
	path := fmt.Sprintf(UrlPersistentPartitionedTopicFormat, tenant, namespace, topic)
	resp, err := p.cli.Delete(path)
	if err != nil {
		return err
	}
	return HttpCheck(resp)
}

// ListPartitioned Get the list of partitioned topics under a namespace.
func (p *PersistentTopicsImpl) ListPartitioned(tenant, namespace string) ([]string, error) {
	path := fmt.Sprintf(UrlPersistentPartitionedNamespaceFormat, tenant, namespace)
	resp, err := p.cli.Get(path)
	if err != nil {
		return nil, err
	}
	data, err := HttpCheckReadBytes(resp)
	if err != nil {
		return nil, err
	}
	var topics []string
	if err := json.Unmarshal(data, &topics); err != nil {
		return nil, err
	}
	return topics, nil
}
