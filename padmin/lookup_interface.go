package padmin

type Lookup interface {
	// GetOwner get the owner broker of the given topic.
	// topicDomain: ref TopicDomain type, value persistent or non-persistent
	GetOwner(topicDomain TopicDomain, tenant, namespace, topic string) (*LookupData, error)
	// GetNamespaceBundle get the namespace bundle which the given topic belongs to.
	GetNamespaceBundle(topicDomain TopicDomain, tenant, namespace, topic string) (string, error)
}

type LookupData struct {
	BrokerUrl    string `json:"brokerUrl"`
	BrokerUrlTls string `json:"brokerUrlTls"`
	HttpUrl      string `json:"httpUrl"`    // web service http address
	HttpUrlTls   string `json:"httpUrlTls"` // web service https address
	NativeUrl    string `json:"nativeUrl"`
}

type TopicDomain string

const (
	TopicDomainPersistent    TopicDomain = "persistent"
	TopicDomainNonPersistent TopicDomain = "non-persistent"
)
