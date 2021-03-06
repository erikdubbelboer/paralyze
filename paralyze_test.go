package paralyze

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestParalyze(t *testing.T) {
	var someError = errors.New("some error")

	slowFn := func() (interface{}, error) {
		time.Sleep(time.Second)
		return "ok", nil
	}
	fastFn := func() (interface{}, error) {
		return 55, nil
	}
	errFn := func() (interface{}, error) {
		return nil, someError
	}
	results, errs := Paralyze(slowFn, fastFn, errFn)

	// Make sure both slices return are the correct length
	assert.Equal(t, 3, len(results))
	assert.Equal(t, 3, len(errs))

	// Assert that return values are in the correct positions
	assert.Equal(t, "ok", results[0])
	assert.Equal(t, 55, results[1])
	assert.Nil(t, results[2])

	// Assert that errors are
	assert.Nil(t, errs[0])
	assert.Nil(t, errs[1])
	assert.Equal(t, someError, errs[2])
}
