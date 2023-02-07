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

func TestNamespaces(t *testing.T) {
	broker := startTestBroker(t)
	defer broker.Close()
	admin := NewTestPulsarAdmin(t, broker.webPort)
	testNs := RandStr(8)
	err := admin.Namespaces.Create("public", testNs)
	require.Nil(t, err)
	namespaces, err := admin.Namespaces.List("public")
	require.Nil(t, err)
	assert.Contains(t, namespaces, fmt.Sprintf("public/%s", testNs))
	err = admin.Namespaces.Delete("public", testNs)
	require.Nil(t, err)
}

func TestNamespacesImpl_OperateNamespaceRetention(t *testing.T) {
	broker := startTestBroker(t)
	defer broker.Close()
	admin := NewTestPulsarAdmin(t, broker.webPort)
	testNs := RandStr(8)
	err := admin.Namespaces.Create("public", testNs)
	require.Nil(t, err)
	err = admin.Namespaces.SetNamespaceRetention("public", testNs, &RetentionConfiguration{
		RetentionSizeInMB:      100,
		RetentionTimeInMinutes: 10,
	})
	require.Nil(t, err)
	cfg, err := admin.Namespaces.GetNamespaceRetention("public", testNs)
	require.Nil(t, err)
	require.EqualValues(t, 100, cfg.RetentionSizeInMB)
	require.EqualValues(t, 10, cfg.RetentionTimeInMinutes)
	err = admin.Namespaces.RemoveNamespaceRetention("public", testNs)
	require.Nil(t, err)
	cfg, err = admin.Namespaces.GetNamespaceRetention("public", testNs)
	require.Nil(t, err)
	require.EqualValues(t, 0, cfg.RetentionSizeInMB)
	require.EqualValues(t, 0, cfg.RetentionTimeInMinutes)
}
