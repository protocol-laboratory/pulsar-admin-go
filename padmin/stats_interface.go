package padmin

type TopicStats interface {
	GetStats(string, string, string) (*TopicStatistics, error)
	GetPartitionedStats(string, string, string) (*TopicStatistics, error)
	GetStatsInternal(string, string, string) (*TopicInternalStats, error)
	GetPartitionedStatsInternal(string, string, string) (*PartitionedTopicInternalStats, error)
}

type LedgerInfo struct {
	LedgerId        int64  `json:"ledgerId"`
	Entries         int64  `json:"entries"`
	Size            int64  `json:"size"`
	Offloaded       bool   `json:"offloaded"`
	Metadata        string `json:"metadata"`
	UnderReplicated bool   `json:"underReplicated"`
}

type TopicInternalStats struct {
	EntriesAddedCounter                int64                  `json:"entriesAddedCounter"`
	NumberOfEntries                    int64                  `json:"numberOfEntries"`
	TotalSize                          int64                  `json:"totalSize"`
	CurrentLedgerEntries               int64                  `json:"currentLedgerEntries"`
	CurrentLedgerSize                  int64                  `json:"currentLedgerSize"`
	LastLedgerCreatedTimestamp         string                 `json:"lastLedgerCreatedTimestamp"`
	LastLedgerCreationFailureTimestamp string                 `json:"lastLedgerCreationFailureTimestamp"`
	WaitingCursorsCount                int32                  `json:"waitingCursorsCount"`
	PendingAddEntriesCount             int32                  `json:"pendingAddEntriesCount"`
	LastConfirmedEntry                 string                 `json:"lastConfirmedEntry"`
	State                              string                 `json:"state"`
	Ledgers                            []*LedgerInfo          `json:"ledgers"`
	Cursors                            map[string]CursorStats `json:"cursors"`
	SchemaLedgers                      []*LedgerInfo          `json:"schemaLedgers"`
	CompactedLedger                    *LedgerInfo            `json:"compactedLedger"`
}

type PartitionedTopicMetadata struct {
	Partitions int32             `json:"partitions"`
	Properties map[string]string `json:"properties"`
}

type PartitionedTopicInternalStats struct {
	Metadata   PartitionedMetadata            `json:"metadata"`
	Partitions map[string]*TopicInternalStats `json:"partitions"`
}

type TopicStatistics struct {
	MsgRateIn                                        float64                      `json:"msgRateIn,omitempty"`
	MsgThroughputIn                                  float64                      `json:"msgThroughputIn,omitempty"`
	MsgRateOut                                       float64                      `json:"msgRateOut,omitempty"`
	MsgThroughputOut                                 float64                      `json:"msgThroughputOut,omitempty"`
	BytesInCounter                                   int64                        `json:"bytesInCounter,omitempty"`
	MsgInCounter                                     int64                        `json:"msgInCounter,omitempty"`
	BytesOutCounter                                  int64                        `json:"bytesOutCounter,omitempty"`
	MsgOutCounter                                    int64                        `json:"msgOutCounter,omitempty"`
	AverageMsgSize                                   float64                      `json:"averageMsgSize,omitempty"`
	MsgChunkPublished                                bool                         `json:"msgChunkPublished,omitempty"`
	StorageSize                                      int64                        `json:"storageSize,omitempty"`
	BacklogSize                                      int64                        `json:"backlogSize,omitempty"`
	PublishRateLimitedTimes                          int64                        `json:"publishRateLimitedTimes,omitempty"`
	EarliestMsgPublishTimeInBacklogs                 int64                        `json:"earliestMsgPublishTimeInBacklogs,omitempty"`
	OffloadedStorageSize                             int64                        `json:"offloadedStorageSize,omitempty"`
	LastOffloadLedgerId                              int64                        `json:"lastOffloadLedgerId,omitempty"`
	LastOffloadSuccessTimeStamp                      int64                        `json:"lastOffloadSuccessTimeStamp,omitempty"`
	LastOffloadFailureTimeStamp                      int64                        `json:"lastOffloadFailureTimeStamp,omitempty"`
	OngoingTxnCount                                  int64                        `json:"ongoingTxnCount,omitempty"`
	AbortedTxnCount                                  int64                        `json:"abortedTxnCount,omitempty"`
	CommittedTxnCount                                int64                        `json:"committedTxnCount,omitempty"`
	Publishers                                       []Publisher                  `json:"publishers,omitempty"`
	WaitingPublishers                                int64                        `json:"waitingPublishers,omitempty"`
	Subscriptions                                    map[string]SubscriptionStats `json:"subscriptions,omitempty"`
	Replication                                      map[string]string            `json:"replication,omitempty"`
	DeduplicationStatus                              string                       `json:"deduplicationStatus,omitempty"`
	NonContiguousDeletedMessagesRanges               int64                        `json:"nonContiguousDeletedMessagesRanges,omitempty"`
	NonContiguousDeletedMessagesRangesSerializedSize int64                        `json:"nonContiguousDeletedMessagesRangesSerializedSize,omitempty"`
	DelayedMessageIndexSizeInBytes                   int64                        `json:"delayedMessageIndexSizeInBytes,omitempty"`
	Compaction                                       Compaction                   `json:"compaction"`
	OwnerBroker                                      string                       `json:"ownerBroker,omitempty"`
}

type CursorStats struct {
	MarkDeletePosition                       string           `json:"markDeletePosition"`
	ReadPosition                             string           `json:"readPosition"`
	WaitingReadOp                            bool             `json:"waitingReadOp"`
	PendingReadOps                           int              `json:"pendingReadOps"`
	MessagesConsumedCounter                  int64            `json:"messagesConsumedCounter"`
	CursorLedger                             int64            `json:"cursorLedger"`
	CursorLedgerLastEntry                    int64            `json:"cursorLedgerLastEntry"`
	IndividuallyDeletedMessages              string           `json:"individuallyDeletedMessages"`
	LastLedgerSwitchTimestamp                string           `json:"lastLedgerSwitchTimestamp"`
	State                                    string           `json:"state"`
	Active                                   bool             `json:"active"`
	NumberOfEntriesSinceFirstNotAckedMessage int64            `json:"numberOfEntriesSinceFirstNotAckedMessage"`
	TotalNonContiguousDeletedMessagesRange   int              `json:"totalNonContiguousDeletedMessagesRange"`
	SubscriptionHavePendingRead              bool             `json:"subscriptionHavePendingRead"`
	SubscriptionHavePendingReplayRead        bool             `json:"subscriptionHavePendingReplayRead"`
	Properties                               map[string]int64 `json:"properties"`
}

type Compaction struct {
	LastCompactionRemovedEventCount   int64 `json:"lastCompactionRemovedEventCount,omitempty"`
	LastCompactionSucceedTimestamp    int64 `json:"lastCompactionSucceedTimestamp,omitempty"`
	LastCompactionFailedTimestamp     int64 `json:"lastCompactionFailedTimestamp,omitempty"`
	LastCompactionDurationTimeInMills int64 `json:"lastCompactionDurationTimeInMills,omitempty"`
}

type SubscriptionStats struct {
	LastConsumedFlowTimestamp int64 `json:"lastConsumedFlowTimestamp,omitempty"`
	MsgBacklog                int   `json:"msgBacklog"`
}

type Publisher struct {
}
