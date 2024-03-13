package ocrHandler

import (
	"fmt"
	"log/slog"
	"os"
)

const languagesFolder = "/languages"
const targetFolder = "/usr/share/tesseract-ocr/4.00/tessdata"

func (handler *OCRShellHandler) initWorkspace() {
	if !exists(handler.workingDir) {
		slog.Error(fmt.Sprintf("Folder %q does not exist. Exiting...", handler.workingDir))
		return
	}

	err := os.Chdir(handler.workingDir)
	if err != nil {
		slog.Error(fmt.Sprintf("Error changing working directory to %q: %s. Exiting...", handler.workingDir, err))
		return
	}

	if !exists("./Original") {
		err = os.Mkdir("./Original", 0755)
		if err != nil {
			slog.Error(fmt.Sprintf("Error creating Original folder: %s. Exiting...", err))
			return
		}
		slog.Info("Created ./Original folder")
	}

	// Copying language files
	languageFiles, err := getFiles(languagesFolder)
	if err != nil {
		slog.Error(fmt.Sprintf("Error reading language files from /languages: %s.", err))
	}
	for _, file := range languageFiles {
		source := fmt.Sprintf("%s/%s", languagesFolder, file)
		destination := fmt.Sprintf("%s/%s", targetFolder, file)
		slog.Info(fmt.Sprintf("Moving language file %q to %q", source, destination))
		err = copyFile(source, destination)
		if err != nil {
			slog.Error(fmt.Sprintf("Error moving language file %q to %q: %s.", source, destination, err))
		}
	}

}
