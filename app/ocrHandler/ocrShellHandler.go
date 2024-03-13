package ocrHandler

import (
	"bytes"
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/exec"
	"strings"
	"time"

	"golang.org/x/sync/semaphore"
)

type OCRShellHandler struct {
	workingDir      string
	scannerInterval int
	semaphore       *semaphore.Weighted
}

/*
workingDir: is the directory where the OCRShellHandler will look for files to process
scanInterval: is the interval in miliseconds between scans of the workingDir
concurrentJobs: is the number of concurrent jobs that the OCRShellHandler will run
*/
func New(workingDir string, scanInterval int, concurrentJobs int) *OCRShellHandler {
	return &OCRShellHandler{
		workingDir:      workingDir,
		scannerInterval: scanInterval,
		semaphore:       semaphore.NewWeighted(int64(concurrentJobs)),
	}
}

func (handler *OCRShellHandler) processFile(filename string) {
	defer handler.semaphore.Release(1)
	handler.semaphore.Acquire(context.Background(), 1)

	languages, newName := extractLanguages(filename)
	if exists(newName) {
		slog.Error(fmt.Sprintf("File %q already exists. Skipping...", newName))
		return
	}
	if exists(fmt.Sprintf("./Original/%s", filename)) {
		slog.Error(fmt.Sprintf("Original File %q already exists. Skipping...", newName))
		return
	}

	slog.Info(fmt.Sprintf("Processing file %q with languages %q", filename, languages))

	runningName := strings.Replace(filename, "DoOcr", "RunningOcr", 1)
	err := os.Rename(filename, runningName)
	if err != nil {
		slog.Error(fmt.Sprintf("Error renaming file %q to %q: %s", filename, runningName, err))
		return
	}

	err = executeOcrmypdf(languages, runningName, newName)
	if err != nil {
		slog.Error(fmt.Sprintf("Error running OCR on file %q: %s", runningName, err))
		return
	}

	err = os.Rename(runningName, fmt.Sprintf("./Original/%s", filename))
	if err != nil {
		slog.Error(fmt.Sprintf("Error moving file %q to Original folder: %q", runningName, err))
		return
	}
	slog.Info(fmt.Sprintf("File %q processed successfully", filename))
}

func executeOcrmypdf(languages string, oldPath string, newPath string) error {
	slog.Debug(fmt.Sprintf("Running: ocrmypdf -l %q %q %q", languages, oldPath, newPath))
	cmd := exec.Command("ocrmypdf", "-l", languages, oldPath, newPath)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	slog.Debug(fmt.Sprintf("Output (stdout) from ocrmypdf: %q", stdout.String()))

	if err != nil {
		return fmt.Errorf("executeOcrmypdf error: \nOriginal err: %s,\nOutput (stderr) from ocrmypdf: %s", err, stderr.String())
	}
	return nil
}

func extractLanguages(filename string) (era string, resr string) {
	start := strings.Index(filename, "-")
	end := strings.Index(filename[start+1:], "-")
	return filename[start+1 : start+end+1], filename[start+end+2:]
}

func exists(path string) bool {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	if err != nil {
		slog.Error(fmt.Sprintf("Error checking if folder %q exists: %s", path, err))
		return false
	}
	return true
}

func (handler *OCRShellHandler) Start() {
	slog.Debug("Starting OCRShellHandler")

	handler.initWorkspace()

	for {
		files, err := getFiles(handler.workingDir)
		if err != nil {
			slog.Error(fmt.Sprintf("Error reading files from working directory: %q. Exiting...", err))
			return
		}

		for _, file := range files {
			if !strings.HasPrefix(file, "DoOcr") {
				continue
			}

			go handler.processFile(file)
		}
		time.Sleep(time.Duration(handler.scannerInterval) * time.Second)
	}
}
