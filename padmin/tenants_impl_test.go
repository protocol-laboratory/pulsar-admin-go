package padmin

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestTenants(t *testing.T) {
	broker := startTestBroker(t)
	defer broker.Close()
	admin := NewTestPulsarAdmin(t, broker.webPort)
	testTenant := RandStr(8)
	err := admin.Tenants.Create(testTenant, TenantInfo{
		AllowedClusters: []string{"standalone"},
	})
	require.Nil(t, err)
	tenants, err := admin.Tenants.List()
	require.Nil(t, err)
	assert.Contains(t, tenants, testTenant)
	err = admin.Tenants.Delete(testTenant)
	require.Nil(t, err)
}
