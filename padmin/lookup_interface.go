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

type Lookup interface {
	// GetOwner get the owner broker of the given topic.
	// topicDomain: ref TopicDomain type, value persistent or non-persistent
	GetOwner(topicDomain TopicDomain, tenant, namespace, topic string) (*LookupData, error)
	// GetNamespaceBundle get the namespace bundle which the given topic belongs to.
	GetNamespaceBundle(topicDomain TopicDomain, tenant, namespace, topic string) (string, error)
}

type LookupData struct {
	BrokerUrl    string `json:"brokerUrl"`
	BrokerUrlTls string `json:"brokerUrlTls"`
	HttpUrl      string `json:"httpUrl"`    // web service http address
	HttpUrlTls   string `json:"httpUrlTls"` // web service https address
	NativeUrl    string `json:"nativeUrl"`
}

type TopicDomain string

const (
	TopicDomainPersistent    TopicDomain = "persistent"
	TopicDomainNonPersistent TopicDomain = "non-persistent"
)
