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
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestLookupImpl_GetOwner(t *testing.T) {
	broker := startTestBroker(t)
	defer broker.Close()
	admin := NewTestPulsarAdmin(t, broker.webPort)
	testTenant := RandStr(8)
	testNs := RandStr(8)
	testTopic := RandStr(8)
	err := admin.Tenants.Create(testTenant, TenantInfo{
		AllowedClusters: []string{"standalone"},
	})
	require.Nil(t, err)
	err = admin.Namespaces.Create(testTenant, testNs)
	require.Nil(t, err)
	err = admin.PersistentTopics.CreateNonPartitioned(testTenant, testNs, testTopic)
	require.Nil(t, err)
	topicList, err := admin.PersistentTopics.ListNonPartitioned(testTenant, testNs)
	require.Nil(t, err)
	if len(topicList) != 1 {
		t.Fatal("topic list should have one topic")
	}
	if topicList[0] != fmt.Sprintf("persistent://%s/%s/%s", testTenant, testNs, testTopic) {
		t.Fatal("topic name should be equal")
	}
	owner, err := admin.Lookup.GetOwner(TopicDomainPersistent, testTenant, testNs, testTopic)
	require.Nil(t, err)
	require.EqualValues(t, "pulsar://localhost:6650", owner.BrokerUrl)
	require.EqualValues(t, "http://localhost:8080", owner.HttpUrl)
	require.EqualValues(t, "pulsar://localhost:6650", owner.NativeUrl)

	testTenant = RandStr(8)
	testNs = RandStr(8)
	testTopic = RandStr(8)
	err = admin.Tenants.Create(testTenant, TenantInfo{
		AllowedClusters: []string{"standalone"},
	})
	require.Nil(t, err)
	err = admin.Namespaces.Create(testTenant, testNs)
	require.Nil(t, err)
	err = admin.NonPersistentTopics.CreateNonPartitioned(testTenant, testNs, testTopic)
	require.Nil(t, err)
	topicList, err = admin.NonPersistentTopics.ListNonPartitioned(testTenant, testNs)
	require.Nil(t, err)
	if len(topicList) != 1 {
		t.Fatal("topic list should have one topic")
	}
	if topicList[0] != fmt.Sprintf("non-persistent://%s/%s/%s", testTenant, testNs, testTopic) {
		t.Fatal("topic name should be equal")
	}
	owner, err = admin.Lookup.GetOwner(TopicDomainNonPersistent, testTenant, testNs, testTopic)
	require.Nil(t, err)
	require.EqualValues(t, "pulsar://localhost:6650", owner.BrokerUrl)
	require.EqualValues(t, "http://localhost:8080", owner.HttpUrl)
	require.EqualValues(t, "pulsar://localhost:6650", owner.NativeUrl)
}

func TestLookupImpl_GetNamespaceBundle(t *testing.T) {
	broker := startTestBroker(t)
	defer broker.Close()
	admin := NewTestPulsarAdmin(t, broker.webPort)
	testTenant := RandStr(8)
	testNs := RandStr(8)
	testTopic := RandStr(8)
	err := admin.Tenants.Create(testTenant, TenantInfo{
		AllowedClusters: []string{"standalone"},
	})
	require.Nil(t, err)
	err = admin.Namespaces.Create(testTenant, testNs)
	require.Nil(t, err)
	err = admin.PersistentTopics.CreatePartitioned(testTenant, testNs, testTopic, 2)
	require.Nil(t, err)
	topicList, err := admin.PersistentTopics.ListPartitioned(testTenant, testNs)
	require.Nil(t, err)
	if len(topicList) != 1 {
		t.Fatal("topic list should have one topic")
	}
	bundle, err := admin.Lookup.GetNamespaceBundle(TopicDomainPersistent, testTenant, testNs, testTopic+"-partition-0")
	require.Nil(t, err)
	t.Logf("bundle: %s", bundle) // 0x00000000_0x40000000
	bundle, err = admin.Lookup.GetNamespaceBundle(TopicDomainPersistent, testTenant, testNs, testTopic+"-partition-1")
	require.Nil(t, err)
	t.Logf("bundle: %s", bundle) // 0x40000000_0x80000000
}
