package padmin

const (
	// UrlPath for the Admin API
	UrlPath             = "/admin/v2"
	UrlClusters         = UrlPath + "/clusters"
	UrlTenants          = UrlPath + "/tenants"
	UrlNamespacesFormat = UrlPath + "/namespaces/%s/%s"
)

// bookies
const (
	UrlBookiesAll         = UrlPath + "/bookies/all"
	UrlBookiesRacksInfo   = UrlPath + "/bookies/racks-info"
	UrlBookiesRacksFormat = UrlPath + "/bookies/racks-info/%s"
)

// lookup
const (
	UrlLookupBrokerFormat             = "/lookup/v2/topic/%s/%s/%s/%s"
	UrlLookupGetNamespaceBundleFormat = "/lookup/v2/topic/%s/%s/%s/%s/bundle"
)

// namespace
const (
	UrlNamespaceRetentionFormat                         = UrlPath + "/namespaces/%s/%s/retention"
	UrlNamespaceGetBacklogQuotaMapFormat                = UrlPath + "/namespaces/%s/%s/backlogQuotaMap"
	UrlNamespaceOperateBacklogQuotaFormat               = UrlPath + "/namespaces/%s/%s/backlogQuota"
	UrlNamespaceClearAllTopicsBacklogFormat             = UrlPath + "/namespaces/%s/%s/clearBacklog"
	UrlNamespaceClearSubscriptionBacklogFormat          = UrlPath + "/namespaces/%s/%s/clearBacklog/%s"
	UrlNamespaceClearAllTopicsBacklogForBundleFormat    = UrlPath + "/namespaces/%s/%s/%s/clearBacklog"
	UrlNamespaceClearSubscriptionBacklogForBundleFormat = UrlPath + "/namespaces/%s/%s/%s/clearBacklog/%s"
	UrlNamespaceCompactionThresholdFormat               = UrlPath + "/namespaces/%s/%s/compactionThreshold"
	UrlNamespaceMessageTTLFormat                        = UrlPath + "/namespaces/%s/%s/messageTTL"
)

// persistent
const (
	UrlPersistentNamespaceFormat                            = UrlPath + "/persistent/%s/%s"
	UrlPersistentTopicFormat                                = UrlPath + "/persistent/%s/%s/%s"
	UrlPersistentPartitionedNamespaceFormat                 = UrlPath + "/persistent/%s/%s/partitioned"
	UrlPersistentPartitionedTopicFormat                     = UrlPath + "/persistent/%s/%s/%s/partitions"
	UrlPersistentPartitionedRetentionFormat                 = UrlPath + "/persistent/%s/%s/%s/retention"
	UrlPersistentTopicGetBacklogQuotaMapFormat              = UrlPath + "/persistent/%s/%s/%s/backlogQuotaMap"
	UrlPersistentTopicOperateBacklogQuotaFormat             = UrlPath + "/persistent/%s/%s/%s/backlogQuota"
	UrlPersistentTopicEstimatedOfflineBacklogFormat         = UrlPath + "/persistent/%s%s%s/backlog"
	UrlPersistentTopicCalculateBacklogSizeByMessageIDFormat = UrlPath + "/persistent/%s/%s/%s/backlogSize"
	UrlPersistentTopicCompactionThresholdFormat             = UrlPath + "/persistent/%s/%s/%s/compactionThreshold"
	UrlPersistentTopicCompactionFormat                      = UrlPath + "/persistent/%s/%s/%s/compaction"
	UrlPersistentTopicMessageTTLFormat                      = UrlPath + "/persistent/%s/%s/%s/messageTTL"
	UrlPersistentTopicCreateMissedPartitionsFormat          = UrlPath + "/persistent/%s/%s/%s/createMissedPartitions"
	UrlPersistentGetLastMessageIdFormat                     = UrlPath + "/persistent/%s/%s/%s/lastMessageId"
	UrlPersistentGetInternalStatsForTopicFormat             = UrlPath + "/persistent/%s/%s/%s/internalStats"
	UrlPersistentGetInternalStatsForPartitionedTopicFormat  = UrlPath + "/persistent/%s/%s/%s/partitioned-internalStats"
	UrlPersistentGetStatsForTopicFormat                     = UrlPath + "/persistent/%s/%s/%s/stats"
	UrlPersistentGetStatsForPartitionedTopicFormat          = UrlPath + "/persistent/%s/%s/%s/partitioned-stats"
)

// non-persistent
const (
	UrlNonPersistentNamespaceFormat                            = UrlPath + "/non-persistent/%s/%s"
	UrlNonPersistentTopicFormat                                = UrlPath + "/non-persistent/%s/%s/%s"
	UrlNonPersistentPartitionedTopicFormat                     = UrlPath + "/non-persistent/%s/%s/%s/partitions"
	UrlNonPersistentPartitionedNamespaceFormat                 = UrlPath + "/non-persistent/%s/%s/partitioned"
	UrlNonPersistentPartitionedRetentionFormat                 = UrlPath + "/non-persistent/%s/%s/%s/retention"
	UrlNonPersistentTopicGetBacklogQuotaMapFormat              = UrlPath + "/non-persistent/%s/%s/%s/backlogQuotaMap"
	UrlNonPersistentTopicOperateBacklogQuotaFormat             = UrlPath + "/non-persistent/%s/%s/%s/backlogQuota"
	UrlNonPersistentTopicEstimatedOfflineBacklogFormat         = UrlPath + "/non-persistent/%s%s%s/backlog"
	UrlNonPersistentTopicCalculateBacklogSizeByMessageIDFormat = UrlPath + "/non-persistent/%s/%s/%s/backlogSize"
	UrlNonPersistentTopicCompactionThresholdFormat             = UrlPath + "/non-persistent/%s/%s/%s/compactionThreshold"
	UrlNonPersistentTopicCompactionFormat                      = UrlPath + "/non-persistent/%s/%s/%s/compaction"
	UrlNonPersistentTopicMessageTTLFormat                      = UrlPath + "/non-persistent/%s/%s/%s/messageTTL"
	UrlNonPersistentTopicsCreateMissedPartitionsFormat         = UrlPath + "/non-persistent/%s/%s/%s/createMissedPartitions"
	UrlNonPersistentGetLastMessageIdFormat                     = UrlPath + "/non-persistent/%s/%s/%s/lastMessageId"
	UrlNonPersistentGetInternalStatsForTopicFormat             = UrlPath + "/non-persistent/%s/%s/%s/internalStats"
	UrlNonPersistentGetInternalStatsForPartitionedTopicFormat  = UrlPath + "/non-persistent/%s/%s/%s/partitioned-internalStats"
	UrlNonPersistentGetStatsForTopicFormat                     = UrlPath + "/non-persistent/%s/%s/%s/stats"
	UrlNonPersistentGetStatsForPartitionedTopicFormat          = UrlPath + "/non-persistent/%s/%s/%s/partitioned-stats"
)
