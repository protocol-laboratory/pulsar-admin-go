package padmin

// NamespaceMessageTTL Discard data after some time (by automatically acknowledging)
type NamespaceMessageTTL interface {
	// GetNamespaceMessageTTL Get the message TTL for the namespace
	GetNamespaceMessageTTL(tenant, namespace string) (int64, error)
	// SetNamespaceMessageTTL Set the message TTL for the namespace
	SetNamespaceMessageTTL(tenant, namespace string, seconds int64) error
	// RemoveNamespaceMessageTTL Remove the message TTL for the namespace
	RemoveNamespaceMessageTTL(tenant, namespace string) error
}

// TopicMessageTTL Discard data after some time (by automatically acknowledging)
type TopicMessageTTL interface {
	// GetTopicMessageTTL Get the message TTL for the topic
	GetTopicMessageTTL(tenant, namespace, topic string) (int64, error)
	// SetTopicMessageTTL Set the message TTL for the topic
	SetTopicMessageTTL(tenant, namespace, topic string, seconds int64) error
	// RemoveTopicMessageTTL Remove the message TTL for the topic
	RemoveTopicMessageTTL(tenant, namespace, topic string) error
}
