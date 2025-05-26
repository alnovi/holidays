package utils

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMust(t *testing.T) {
	assert.Panics(t, func() {
		Must(errors.New("error"))
	})

	assert.NotPanics(t, func() {
		Must(nil)
	})
}

func TestMustMsg(t *testing.T) {
	assert.Panics(t, func() {
		MustMsg(errors.New("error"), "message")
	})

	assert.Panics(t, func() {
		MustMsg(errors.New("error"), "")
	})

	assert.NotPanics(t, func() {
		MustMsg(nil, "message")
	})
}
