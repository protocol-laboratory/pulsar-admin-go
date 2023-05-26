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
	BytesInCounter                                   uint64                       `json:"bytesInCounter,omitempty"`
	MsgInCounter                                     uint64                       `json:"msgInCounter,omitempty"`
	BytesOutCounter                                  uint64                       `json:"bytesOutCounter,omitempty"`
	MsgOutCounter                                    uint64                       `json:"msgOutCounter,omitempty"`
	AverageMsgSize                                   float64                      `json:"averageMsgSize,omitempty"`
	MsgChunkPublished                                bool                         `json:"msgChunkPublished,omitempty"`
	StorageSize                                      uint64                       `json:"storageSize,omitempty"`
	BacklogSize                                      uint64                       `json:"backlogSize,omitempty"`
	PublishRateLimitedTimes                          uint64                       `json:"publishRateLimitedTimes,omitempty"`
	EarliestMsgPublishTimeInBacklogs                 uint64                       `json:"earliestMsgPublishTimeInBacklogs,omitempty"`
	OffloadedStorageSize                             uint64                       `json:"offloadedStorageSize,omitempty"`
	LastOffloadLedgerId                              uint64                       `json:"lastOffloadLedgerId,omitempty"`
	LastOffloadSuccessTimeStamp                      uint64                       `json:"lastOffloadSuccessTimeStamp,omitempty"`
	LastOffloadFailureTimeStamp                      uint64                       `json:"lastOffloadFailureTimeStamp,omitempty"`
	OngoingTxnCount                                  uint64                       `json:"ongoingTxnCount,omitempty"`
	AbortedTxnCount                                  uint64                       `json:"abortedTxnCount,omitempty"`
	CommittedTxnCount                                uint64                       `json:"committedTxnCount,omitempty"`
	Publishers                                       []string                     `json:"publishers,omitempty"`
	WaitingPublishers                                uint64                       `json:"waitingPublishers,omitempty"`
	Subscriptions                                    map[string]SubscriptionStats `json:"subscriptions,omitempty"`
	Replication                                      map[string]string            `json:"replication,omitempty"`
	DeduplicationStatus                              string                       `json:"deduplicationStatus,omitempty"`
	NonContiguousDeletedMessagesRanges               uint64                       `json:"nonContiguousDeletedMessagesRanges,omitempty"`
	NonContiguousDeletedMessagesRangesSerializedSize uint64                       `json:"nonContiguousDeletedMessagesRangesSerializedSize,omitempty"`
	DelayedMessageIndexSizeInBytes                   uint64                       `json:"delayedMessageIndexSizeInBytes,omitempty"`
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
	LastCompactionRemovedEventCount   uint64 `json:"lastCompactionRemovedEventCount,omitempty"`
	LastCompactionSucceedTimestamp    uint64 `json:"lastCompactionSucceedTimestamp,omitempty"`
	LastCompactionFailedTimestamp     uint64 `json:"lastCompactionFailedTimestamp,omitempty"`
	LastCompactionDurationTimeInMills uint64 `json:"lastCompactionDurationTimeInMills,omitempty"`
}

type SubscriptionStats struct {
	LastConsumedFlowTimestamp uint64 `json:"lastConsumedFlowTimestamp,omitempty"`
	MsgBacklog                int    `json:"msgBacklog"`
}
