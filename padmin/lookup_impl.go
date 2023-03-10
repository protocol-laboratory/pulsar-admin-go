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

type LookupImpl struct {
	cli HttpClient
}

func (l *LookupImpl) GetOwner(topicDomain TopicDomain, tenant, namespace, topic string) (*LookupData, error) {
	url := fmt.Sprintf(UrlLookupBrokerFormat, topicDomain, tenant, namespace, topic)
	resp, err := l.cli.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var body = new(LookupData)
	if err := EasyReader(resp, body); err != nil {
		return nil, err
	}
	return body, nil
}

func (l *LookupImpl) GetNamespaceBundle(topicDomain TopicDomain, tenant, namespace, topic string) (string, error) {
	url := fmt.Sprintf(UrlLookupGetNamespaceBundleFormat, topicDomain, tenant, namespace, topic)
	resp, err := l.cli.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	res, err := ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return res, nil
}

func newLookup(cli HttpClient) Lookup {
	return &LookupImpl{cli: cli}
}
