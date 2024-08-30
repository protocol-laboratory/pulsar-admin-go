package padmin

import "time"

type NamespaceBacklog interface {
	// GetNamespaceBacklogQuota Get backlog quota map on a namespace.
	GetNamespaceBacklogQuota(tenant, namespace string) (*BacklogQuotaResp, error)
	// SetNamespaceBacklogQuota  Set a backlog quota for all the topics on a namespace.
	SetNamespaceBacklogQuota(tenant, namespace string, cfg *BacklogQuota) error
	// RemoveNamespaceBacklogQuota Remove a backlog quota policy from a namespace.
	RemoveNamespaceBacklogQuota(tenant, namespace string, opts ...Option) error
	//ClearNamespaceAllTopicsBacklog Clear backlog for all topics on a namespace.
	ClearNamespaceAllTopicsBacklog(tenant, namespace string) error
	// ClearNamespaceSubscriptionBacklog Clear backlog for a given subscription on all topics on a namespace.
	ClearNamespaceSubscriptionBacklog(tenant, namespace, subscription string) error
	// ClearNamespaceAllTopicsBacklogForBundle Clear backlog for all topics on a namespace bundle.
	ClearNamespaceAllTopicsBacklogForBundle(tenant, namespace, bundle string) error
	// ClearNamespaceSubscriptionBacklogForBundle Clear backlog for a given subscription on all topics on a namespace bundle.
	ClearNamespaceSubscriptionBacklogForBundle(tenant, namespace, subscription, bundle string) error
}

type TopicBacklog interface {
	// GetTopicBacklogQuota Get backlog quota map on a topic.
	GetTopicBacklogQuota(tenant, namespace, topic string) (*BacklogQuotaResp, error)
	// SetTopicBacklogQuota Set a backlog quota for a topic.
	SetTopicBacklogQuota(tenant, namespace, topic string, cfg *BacklogQuota) error
	// RemoveTopicBacklogQuota Remove a backlog quota policy from a topic.
	RemoveTopicBacklogQuota(tenant, namespace, topic string, opts ...Option) error
	// EstimatedOfflineTopicBacklog Get estimated backlog for offline topic.
	EstimatedOfflineTopicBacklog(tenant, namespace, topic string) (*OfflineTopicStats, error)
	// CalculateBacklogSizeByMessageID Calculate backlog size by a message ID (in bytes).
	CalculateBacklogSizeByMessageID(tenant, namespace, topic string) error
}

type BacklogQuotaType string

const (
	DestinationStorage BacklogQuotaType = "destination_storage"
	MessageAge         BacklogQuotaType = "message_age"
)

type BacklogQuotaResp struct {
	DestinationStorage BacklogQuota `json:"destination_storage,omitempty"`
	MessageAge         BacklogQuota `json:"message_age,omitempty"`
}

type RetentionPolicy string

const (
	ProducerRequestHold     RetentionPolicy = "producer_request_hold"
	ProducerException       RetentionPolicy = "producer_exception"
	ConsumerBacklogEviction RetentionPolicy = "consumer_backlog_eviction"
)

type BacklogQuota struct {
	Limit     int64
	LimitSize int64
	LimitTime int64
	Policy    RetentionPolicy
}

type OfflineTopicStats struct {
	StorageSize       int64
	TotalMessages     int64
	MessageBacklog    int64
	BrokerName        string
	TopicName         string
	DataLedgerDetails []LedgerDetail
	CursorDetails     map[string]CursorDetail
	StatGeneratedAt   time.Time
}

type LedgerDetail struct {
	Entries   int64
	Timestamp int64
	Size      int64
	LedgerId  int64
}

type CursorDetail struct {
	CursorBacklog  int64
	CursorLedgerId int64
}
