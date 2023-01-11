package errors

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	ec := New()

	ec.Add(errors.New("err 1"), 1, "msg 1")
	ec.Add(errors.New("err 2"), 2, "msg 2")
	ec.Add(errors.New("err 3"), 3, "msg 3")
	fmt.Printf("ec=%v\n", ec)

	assert.Equal(t, 1, 2)
}
