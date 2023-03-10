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
	GetPartitionedStats(string, string, string) ([]*TopicStatistics, error)
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
