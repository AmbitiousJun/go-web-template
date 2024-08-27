package array_test

import (
	"go_web_template/internal/model/entity"
	"go_web_template/internal/util/array"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSafetyGet(t *testing.T) {
	arr := []interface{}{1, entity.NewTime(time.Now())}
	_, ok := array.SafetyGet(arr, -1, 0)
	assert.False(t, ok)
	_, ok = array.SafetyGet(arr, 2, 0)
	assert.False(t, ok)
	res, ok := array.SafetyGet(arr, 0, 0)
	assert.True(t, ok)
	assert.Equal(t, arr[0], res)
	_, ok = array.SafetyGet(arr, 0, (*time.Time)(nil))
	assert.False(t, ok)
	res1, ok := array.SafetyGet(arr, 1, (*entity.Time)(nil))
	assert.True(t, ok)
	assert.Equal(t, arr[1], res1)
}
