package time

import (
	"testing"

	"github.com/project-flogo/core/data/expression/function"
	"github.com/stretchr/testify/assert"
)

var in = &fnNow{}

func init() {
	function.ResolveAliases()
}

func TestSample(t *testing.T) {
	final1, _ := in.Eval()
	assert.NotNil(t, final1)
}
