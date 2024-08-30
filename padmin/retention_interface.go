package padmin

// NamespaceRetention operate interface about retention configuration for namespace.
type NamespaceRetention interface {
	// GetNamespaceRetention get retention configuration for namespace.
	GetNamespaceRetention(tenant, namespace string) (*RetentionConfiguration, error)
	// SetNamespaceRetention set retention configuration for namespace.
	SetNamespaceRetention(tenant, namespace string, cfg *RetentionConfiguration) error
	// RemoveNamespaceRetention remove retention configuration for namespace.
	RemoveNamespaceRetention(tenant, namespace string) error
}

// TopicRetention operate interface about retention configuration for specified topic.
type TopicRetention interface {
	// GetTopicRetention get retention configuration for namespace.
	GetTopicRetention(tenant, namespace, topic string) (*RetentionConfiguration, error)
	// SetTopicRetention set retention configuration for namespace.
	SetTopicRetention(tenant, namespace, topic string, cfg *RetentionConfiguration) error
	// RemoveTopicRetention remove retention configuration for namespace.
	RemoveTopicRetention(tenant, namespace, topic string) error
}

// RetentionConfiguration retention configuration
type RetentionConfiguration struct {
	RetentionSizeInMB      int64 `json:"retentionSizeInMB"`
	RetentionTimeInMinutes int64 `json:"retentionTimeInMinutes"`
}
