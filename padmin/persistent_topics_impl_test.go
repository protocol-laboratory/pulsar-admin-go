package padmin

import (
	"fmt"
	"github.com/stretchr/testify/assert"
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

func Test_persistentTopics_OperateTopicRetention(t *testing.T) {
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
	err = admin.Namespaces.SetNamespaceRetention(testTenant, testNs, &RetentionConfiguration{
		RetentionSizeInMB:      100,
		RetentionTimeInMinutes: 10,
	})
	require.Nil(t, err)
	cfg, err := admin.Namespaces.GetNamespaceRetention(testTenant, testNs)
	require.Nil(t, err)
	require.EqualValues(t, 100, cfg.RetentionSizeInMB)
	require.EqualValues(t, 10, cfg.RetentionTimeInMinutes)
	err = admin.PersistentTopics.SetTopicRetention(testTenant, testNs, testTopic, &RetentionConfiguration{
		RetentionSizeInMB:      50,
		RetentionTimeInMinutes: 5,
	})
	assert.Error(t, err)
}

func TestPersistentTopicsImpl_OperateTopicMessageTTL(t *testing.T) {
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
	err = admin.Namespaces.SetNamespaceMessageTTL(testTenant, testNs, 30)
	require.Nil(t, err)
	_, err = admin.PersistentTopics.GetTopicMessageTTL(testTenant, testNs, testTopic)
	require.Error(t, err)
	err = admin.PersistentTopics.SetTopicMessageTTL(testTenant, testNs, testTopic, 30)
	require.Error(t, err)
	_, err = admin.PersistentTopics.GetTopicMessageTTL(testTenant, testNs, testTopic)
	require.Error(t, err)
}

func TestPersistentTopicsImpl_OperateTopicCompactionThreshold(t *testing.T) {
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
	threshold, err := admin.PersistentTopics.GetTopicCompactionThreshold(testTenant, testNs, testTopic)
	require.Error(t, err)
	require.EqualValues(t, 0, threshold)
}

func TestPersistentTopicsImpl_CreateMissedPartitions(t *testing.T) {
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
	err = admin.PersistentTopics.CreateMissedPartitions(testTenant, testNs, testTopic)
	require.Nil(t, err)
}

func TestPersistentTopicsImpl_GetLastMessageId(t *testing.T) {
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
	msg, err := admin.PersistentTopics.GetLastMessageId(testTenant, testNs, testTopic+"-partition-0")
	require.Nil(t, err)
	require.Greater(t, msg.LedgerId, int64(0))
}

func TestPersistentTopicsImpl_GetStatsInternalForPartitionTopic(t *testing.T) {
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
	stats, err := admin.PersistentTopics.GetPartitionedStatsInternal(testTenant, testNs, testTopic)
	require.Nil(t, err)
	require.EqualValues(t, 2, stats.Metadata.Partitions)
	t.Logf("get stats: %+v", stats)
}

func TestPersistentTopicsImpl_GetStatsInternalForTopic(t *testing.T) {
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
	stats, err := admin.PersistentTopics.GetStatsInternal(testTenant, testNs, testTopic)
	require.Nil(t, err)
	t.Logf("get stats: %+v", stats)
}

func TestPersistentTopicsImpl_GetStats(t *testing.T) {
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
	stats, err := admin.PersistentTopics.GetStats(testTenant, testNs, testTopic)
	require.Nil(t, err)
	t.Logf("get stats: %+v", stats)
}

func TestPersistentTopicsImpl_GetPartitionedStats(t *testing.T) {
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
	statsList, err := admin.PersistentTopics.GetPartitionedStats(testTenant, testNs, testTopic)
	require.Nil(t, err)
	t.Logf("get partitioned stats:%+v", statsList)
}
