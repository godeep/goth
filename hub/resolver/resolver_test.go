package resolver

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_RawResolve(t *testing.T) {
	resolved, err := Resolver().GetAddress("prod.app.serve.tam.tam")
	assert.NoError(t, err)
	t.Logf("%+v", resolved)
}
