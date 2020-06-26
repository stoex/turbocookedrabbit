package main_test

import (
	"testing"

	"github.com/houseofcat/turbocookedrabbit/pkg/tcr"
	"github.com/stretchr/testify/assert"
)

func TestReadConfig(t *testing.T) {
	fileNamePath := "testseasoning.json"

	assert.FileExists(t, fileNamePath)

	config, err := tcr.ConvertJSONFileToConfig(fileNamePath)

	assert.Nil(t, err)
	assert.NotEqual(t, "", config.PoolConfig.ConnectionPoolConfig.URI, "RabbitMQ URI should not be blank.")
}
