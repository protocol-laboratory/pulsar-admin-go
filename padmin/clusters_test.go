package padmin

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestClusters(t *testing.T) {
	broker := startTestBroker(t)
	defer broker.Close()
	admin := NewTestPulsarAdmin(t, broker.webPort)
	clusters, err := admin.Clusters.List()
	require.Nil(t, err)
	assert.Contains(t, clusters, "standalone")
}
