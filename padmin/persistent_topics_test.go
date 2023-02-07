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

func TestPersistentNonPartitionedTopics(t *testing.T) {
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
	err = admin.PersistentTopics.DeleteNonPartitioned(testTenant, testNs, testTopic)
	require.Nil(t, err)
	topicList, err = admin.PersistentTopics.ListNonPartitioned(testTenant, testNs)
	require.Nil(t, err)
	if len(topicList) != 0 {
		t.Fatal("topic list should be empty")
	}
	err = admin.Namespaces.Delete(testTenant, testNs)
	require.Nil(t, err)
	err = admin.Tenants.Delete(testTenant)
	require.Nil(t, err)
}

func TestPersistentPartitionedTopics(t *testing.T) {
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
	if topicList[0] != fmt.Sprintf("persistent://%s/%s/%s", testTenant, testNs, testTopic) {
		t.Fatal("topic name should be equal")
	}
	err = admin.PersistentTopics.DeletePartitioned(testTenant, testNs, testTopic)
	require.Nil(t, err)
	topicList, err = admin.PersistentTopics.ListPartitioned(testTenant, testNs)
	require.Nil(t, err)
	if len(topicList) != 0 {
		t.Fatal("topic list should be empty")
	}
	err = admin.Namespaces.Delete(testTenant, testNs)
	require.Nil(t, err)
	err = admin.Tenants.Delete(testTenant)
	require.Nil(t, err)
}

func Test_persistentTopics_ListNamespaceTopics(t *testing.T) {
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
	topicList, err := admin.PersistentTopics.ListNamespaceTopics(testTenant, testNs)
	require.Nil(t, err)
	t.Logf("topic list: %v", topicList)
}

func Test_persistentTopics_GetPartitionedMetadata(t *testing.T) {
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
	if topicList[0] != fmt.Sprintf("persistent://%s/%s/%s", testTenant, testNs, testTopic) {
		t.Fatal("topic name should be equal")
	}
	metadata, err := admin.PersistentTopics.GetPartitionedMetadata(testTenant, testNs, testTopic)
	require.Nil(t, err)
	require.EqualValues(t, metadata.Partitions, 2)
	require.Equal(t, metadata.Deleted, false)
}

func Test_persistentTopics_GetRetention(t *testing.T) {
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
	if topicList[0] != fmt.Sprintf("persistent://%s/%s/%s", testTenant, testNs, testTopic) {
		t.Fatal("topic name should be equal")
	}
	retention, err := admin.PersistentTopics.GetRetention(testTenant, testNs, testTopic)
	t.Logf("get err: %v", err)
	t.Logf("get retention: %v", retention)
}
