package padmin

import (
	"crypto/tls"
	"strconv"
)

type Config struct {
	urlPrefix string
	// Host pulsar service address, default localhost
	Host string
	// Port pulsar service port, default 8080
	Port int
	// TlsEnable enable tls, default false
	TlsEnable bool
	// TlsConfig tls config
	TlsConfig *tls.Config
	// ConnectionTimeout connect timeout, default 0, zero means no timeout
	ConnectionTimeout int64
}

type PulsarAdmin struct {
	Clusters            Clusters
	Tenants             Tenants
	Namespaces          Namespaces
	PersistentTopics    PersistentTopics
	NonPersistentTopics NonPersistentTopics
	Bookies             Bookies
	Lookup              Lookup
}

func NewDefaultPulsarAdmin() (*PulsarAdmin, error) {
	return NewPulsarAdmin(Config{})
}

func NewPulsarAdmin(config Config) (*PulsarAdmin, error) {
	if config.Host == "" {
		config.Host = "localhost"
	}
	if config.Port == 0 {
		config.Port = 8080
	}
	config.urlPrefix = "http://" + config.Host + ":" + strconv.Itoa(config.Port)
	client, err := newHttpClient(config)
	if err != nil {
		return nil, err
	}
	return &PulsarAdmin{
		Clusters:            newClusters(client),
		Tenants:             newTenants(client),
		Namespaces:          newNamespaces(client),
		PersistentTopics:    newPersistentTopics(client),
		NonPersistentTopics: newNonPersistentTopics(client),
		Bookies:             newBookies(client),
		Lookup:              newLookup(client),
	}, nil
}
