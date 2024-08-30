package padmin

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func NewTestPulsarAdmin(t *testing.T, port int) *PulsarAdmin {
	admin, err := NewPulsarAdmin(Config{
		Host: "localhost",
		Port: port,
	})
	require.Nil(t, err)
	return admin
}
