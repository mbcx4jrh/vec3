package vec3

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var delta float64 = 0.0001

func TestDotProduct(t *testing.T) {
	v1 := Vector3{1, 2, 3}
	v2 := Vector3{2, 3, 4}
	res := Dot(v1, v2)

	assert.InDelta(t, 20.0, res, delta)
}

func TestCrossProduct(t *testing.T) {
	i := Vector3{1, 0, 0}
	j := Vector3{0, 1, 0}
	k := Vector3{0, 0, 1}

	assert := assert.New(t)

	assert.Equal(k, Cross(i, j))
	assert.Equal(i, Cross(j, k))
	assert.Equal(j, Cross(k, i))
}
