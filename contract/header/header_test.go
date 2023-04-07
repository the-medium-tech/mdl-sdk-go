package header

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHeaderType(t *testing.T) {
	assert.Equal(t, "fabric", HeaderTypeToString(FABRIC))
	assert.Equal(t, "ethereum", HeaderTypeToString(ETHEREUM))
	assert.Equal(t, "bitcoin", HeaderTypeToString(BITCOIN))
}
