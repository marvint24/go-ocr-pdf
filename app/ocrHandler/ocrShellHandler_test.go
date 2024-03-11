package ocrHandler

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetLanguages(t *testing.T) {
	filename := "DoOcr-deu+eng-balblup.pdf"
	expectedLanguage := "deu+eng"
	expectedFileName := "balblup.pdf"

	actualLanguage, actualFileName := extractLanguages(filename)
	assert.Equal(t, expectedLanguage, actualLanguage, fmt.Sprintf("Expected language to be %q", expectedLanguage))
	assert.Equal(t, expectedFileName, actualFileName, fmt.Sprintf("Expected fileName to be %q", expectedFileName))

	filename = "DoOcr-deu+eng-balblup-hi-was-geht.pdf"
	expectedLanguage = "deu+eng"
	expectedFileName = "balblup-hi-was-geht.pdf"

	actualLanguage, actualFileName = extractLanguages(filename)
	assert.Equal(t, expectedLanguage, actualLanguage, fmt.Sprintf("Expected language to be %q", expectedLanguage))
	assert.Equal(t, expectedFileName, actualFileName, fmt.Sprintf("Expected fileName to be %q", expectedFileName))

	filename = "DoOcr-deu+eng+jpn-a.pdf"
	expectedLanguage = "deu+eng+jpn"
	expectedFileName = "a.pdf"

	actualLanguage, actualFileName = extractLanguages(filename)
	assert.Equal(t, expectedLanguage, actualLanguage, fmt.Sprintf("Expected language to be %q", expectedLanguage))
	assert.Equal(t, expectedFileName, actualFileName, fmt.Sprintf("Expected fileName to be %q", expectedFileName))
}
