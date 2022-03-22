package id

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestPkgID_NewID(t *testing.T) {
	id := NewID()
	require.IsType(t, uuid.UUID{}, id)
}
