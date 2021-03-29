package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPwd(t *testing.T) {
	c := Cmd{
		command: "pwd",
	}
	actual := c.pwd()
	assert.Equal(t, "", actual)
}
