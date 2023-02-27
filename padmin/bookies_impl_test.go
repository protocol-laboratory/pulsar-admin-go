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
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestBookiesImpl_AllBookies(t *testing.T) {
	broker := startTestBroker(t)
	defer broker.Close()
	admin := NewTestPulsarAdmin(t, broker.webPort)
	testTenant := RandStr(8)
	err := admin.Tenants.Create(testTenant, TenantInfo{
		AllowedClusters: []string{"standalone"},
	})
	require.Nil(t, err)
	testNs := RandStr(8)
	err = admin.Namespaces.Create(testTenant, testNs)
	require.Nil(t, err)
	namespaces, err := admin.Namespaces.List(testTenant)
	require.Nil(t, err)
	assert.Contains(t, namespaces, fmt.Sprintf("%s/%s", testTenant, testNs))
	testTopic := RandStr(8)
	err = admin.NonPersistentTopics.CreatePartitioned(testTenant, testNs, testTopic, 2)
	require.Nil(t, err)
	topicList, err := admin.NonPersistentTopics.ListPartitioned(testTenant, testNs)
	require.Nil(t, err)
	if len(topicList) != 1 {
		t.Fatal("topic list should have one topic")
	}
	if topicList[0] != fmt.Sprintf("non-persistent://%s/%s/%s", testTenant, testNs, testTopic) {
		t.Fatal("topic name should be equal")
	}
	resp, err := admin.Bookies.AllBookies()
	require.Nil(t, err)
	t.Logf("%+v", resp)
}

func TestBookiesImpl_ListRacksInfo(t *testing.T) {
	broker := startTestBroker(t)
	defer broker.Close()
	admin := NewTestPulsarAdmin(t, broker.webPort)
	testTenant := RandStr(8)
	err := admin.Tenants.Create(testTenant, TenantInfo{
		AllowedClusters: []string{"standalone"},
	})
	require.Nil(t, err)
	testNs := RandStr(8)
	err = admin.Namespaces.Create(testTenant, testNs)
	require.Nil(t, err)
	namespaces, err := admin.Namespaces.List(testTenant)
	require.Nil(t, err)
	assert.Contains(t, namespaces, fmt.Sprintf("%s/%s", testTenant, testNs))
	testTopic := RandStr(8)
	err = admin.NonPersistentTopics.CreatePartitioned(testTenant, testNs, testTopic, 2)
	require.Nil(t, err)
	topicList, err := admin.NonPersistentTopics.ListPartitioned(testTenant, testNs)
	require.Nil(t, err)
	if len(topicList) != 1 {
		t.Fatal("topic list should have one topic")
	}
	if topicList[0] != fmt.Sprintf("non-persistent://%s/%s/%s", testTenant, testNs, testTopic) {
		t.Fatal("topic name should be equal")
	}
	_, err = admin.Bookies.AllBookies()
	require.Nil(t, err)
	resp, err := admin.Bookies.ListRacksInfo()
	require.Nil(t, err)
	t.Logf("get info: %+v", resp)
}

func TestBookiesImpl_GetRacksInfo(t *testing.T) {
	broker := startTestBroker(t)
	defer broker.Close()
	admin := NewTestPulsarAdmin(t, broker.webPort)
	testTenant := RandStr(8)
	err := admin.Tenants.Create(testTenant, TenantInfo{
		AllowedClusters: []string{"standalone"},
	})
	require.Nil(t, err)
	testNs := RandStr(8)
	err = admin.Namespaces.Create(testTenant, testNs)
	require.Nil(t, err)
	namespaces, err := admin.Namespaces.List(testTenant)
	require.Nil(t, err)
	assert.Contains(t, namespaces, fmt.Sprintf("%s/%s", testTenant, testNs))
	testTopic := RandStr(8)
	err = admin.NonPersistentTopics.CreatePartitioned(testTenant, testNs, testTopic, 2)
	require.Nil(t, err)
	topicList, err := admin.NonPersistentTopics.ListPartitioned(testTenant, testNs)
	require.Nil(t, err)
	if len(topicList) != 1 {
		t.Fatal("topic list should have one topic")
	}
	if topicList[0] != fmt.Sprintf("non-persistent://%s/%s/%s", testTenant, testNs, testTopic) {
		t.Fatal("topic name should be equal")
	}
	bookies, err := admin.Bookies.AllBookies()
	require.Nil(t, err)
	racksInfos, err := admin.Bookies.ListRacksInfo()
	require.Nil(t, err)
	if len(*racksInfos) == 0 {
		return
	}
	info, err := admin.Bookies.GetRacksInfo(bookies.Bookies[0].BookieId)
	require.Nil(t, err)
	t.Logf("rack info: %v", info)
}
