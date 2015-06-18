package hub

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Clone(t *testing.T) {
	o1 := GetTOptions()
	o2 := GetTOptions()

	o1.DiscoveryEnv = "o1"
	assert.NotEqual(t, o1.DiscoveryEnv, o2.DiscoveryEnv)
}
