package test

import (
	"github.com/stretchr/testify/assert"
	"learn-golang-restful-api/simple"
	"testing"
)

func TestConnection(t *testing.T) {
	connection, cleanup := simple.InitializedConnection("Database")
	assert.NotNil(t, connection)

	cleanup()
}
