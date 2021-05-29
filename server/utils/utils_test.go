package utils_test

import (
	"testing"

	"github.com/hokkung/go-gateway/server/utils"
	"github.com/stretchr/testify/assert"
)

func TestIsAllowedUserAgent(t *testing.T) {
	assert := assert.New(t)

	result, err := utils.IsAllowedUserAgent("ios-1")

	assert.Equal(true, result)
	assert.Equal(nil, err)
}
