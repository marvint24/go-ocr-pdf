package main

import (
	"fmt"
	"log/slog"
	ocrhandler "ocrTool/ocrHandler"
	"os"
	"path"
	"strconv"
)

const (
	defaultLogLevel       = 8
	defaultScanInterval   = 10
	defaultConcurrentJobs = 4
	basePath              = "/data"
)

func getLogLevel() slog.Level {
	logLevelInt, err := strconv.Atoi(os.Getenv("LOG_LEVEL"))
	if err != nil {
		slog.Error(fmt.Sprintf("Error parsing environment var LOG_LEVEL: %s", err))
		logLevelInt = defaultLogLevel
	}

	return slog.Level(logLevelInt)
}

func readEnvInt(envVar string, defaultValue int) int {
	envValue := os.Getenv(envVar)
	if envValue == "" {
		return defaultValue
	}

	value, err := strconv.Atoi(envValue)
	if err != nil {
		slog.Error(fmt.Sprintf("Error parsing environment var %q: %s", envVar, err))
		return defaultValue
	}

	return value
}

func main() {
	logLevel := getLogLevel()
	slog.SetLogLoggerLevel(logLevel)

	scanFolder := os.Getenv("SCAN_FOLDER")
	scanInterval := readEnvInt("SCAN_INTERVAL", defaultScanInterval)
	concurrentJobs := readEnvInt("CONCURRENT_JOBS", defaultConcurrentJobs)

	slog.Info(fmt.Sprintf("Starting OCRTool with log level %d, scan folder %q, scan interval %d, and concurrent jobs %d", logLevel, scanFolder, scanInterval, concurrentJobs))

	ocrhandler := ocrhandler.New(path.Join(basePath, scanFolder), scanInterval, concurrentJobs)

	ocrhandler.Start()

}
