package padmin

type MessageId struct {
	LedgerId       int64 `json:"ledgerId"`
	EntryId        int64 `json:"entryId"`
	PartitionIndex int32 `json:"partitionIndex"`
	BatchIndex     int32 `json:"batchIndex"`
}
