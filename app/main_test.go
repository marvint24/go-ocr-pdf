package main

import (
	"log/slog"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetLogLevel(t *testing.T) {
	// Test case 1: LOG_LEVEL environment variable is set to a valid integer
	os.Setenv("LOG_LEVEL", "4")
	expectedLevel := slog.Level(4)
	actualLevel := getLogLevel()
	assert.Equal(t, expectedLevel, actualLevel, "Expected log level to be 4")

	// Test case 2: LOG_LEVEL environment variable is not set
	os.Unsetenv("LOG_LEVEL")
	expectedLevel = slog.Level(8)
	actualLevel = getLogLevel()
	assert.Equal(t, expectedLevel, actualLevel, "Expected log level to be 8")

	// Test case 3: LOG_LEVEL environment variable is set to an invalid integer
	os.Setenv("LOG_LEVEL", "invalid")
	expectedLevel = slog.Level(8)
	actualLevel = getLogLevel()
	assert.Equal(t, expectedLevel, actualLevel, "Expected log level to be 8")
}
