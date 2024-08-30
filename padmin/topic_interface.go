package padmin

// Topics persistent topic and non-persistent topic interface
type Topics interface {
	CreateNonPartitioned(tenant, namespace, topic string) error
	CreatePartitioned(tenant, namespace, topic string, numPartitions int) error
	DeleteNonPartitioned(tenant, namespace, topic string) error
	DeletePartitioned(tenant, namespace, topic string) error
	ListNonPartitioned(tenant, namespace string) ([]string, error)
	ListPartitioned(tenant, namespace string) ([]string, error)
	ListNamespaceTopics(tenant, namespace string) ([]string, error)
	GetPartitionedMetadata(tenant, namespace, topic string) (*PartitionedMetadata, error)
	CreateMissedPartitions(tenant, namespace, topic string) error
	GetLastMessageId(tenant, namespace, topic string) (*MessageId, error)
	TopicRetention
}

// PartitionedMetadata partitioned topic metadata
type PartitionedMetadata struct {
	Deleted    bool  `json:"deleted"`
	Partitions int64 `json:"partitions"`
}
