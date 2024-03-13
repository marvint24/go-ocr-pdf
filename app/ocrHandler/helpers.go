package ocrHandler

import (
	"fmt"
	"os"
)

func getFiles(direcrory string) ([]string, error) {
	fileNames := []string{}
	dirEntries, err := os.ReadDir(direcrory)
	if err != nil {
		return nil, err
	}

	for _, entry := range dirEntries {
		fileNames = append(fileNames, entry.Name())
	}

	return fileNames, nil
}

func copyFile(source string, destination string) error {
	srcData, err := os.ReadFile(source)
	if err != nil {
		return fmt.Errorf("error reading file %q: %s", source, err)
	}

	destFile, err := os.Create(destination)
	if err != nil {
		return fmt.Errorf("error creating new file %q: %s", destination, err)
	}
	defer destFile.Close()

	_, err = destFile.Write(srcData)
	if err != nil {
		return fmt.Errorf("error writing to file %q: %s", destination, err)
	}

	return nil
}
