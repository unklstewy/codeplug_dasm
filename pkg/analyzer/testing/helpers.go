package testing

import (
	"os"
	"path/filepath"
	"runtime"
	"testing"
)

// GetTestDataPath returns the absolute path to the testdata directory
func GetTestDataPath() string {
	_, file, _, _ := runtime.Caller(0)
	return filepath.Join(filepath.Dir(file), "..", "testdata")
}

// LoadTestCodeplug loads a codeplug file from the testdata directory
func LoadTestCodeplug(t *testing.T, relativePath string) []byte {
	t.Helper()

	path := filepath.Join(GetTestDataPath(), relativePath)
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("Failed to load test codeplug %s: %v", path, err)
	}
	return data
}

// SaveExpectedResult saves an expected result for future comparisons
func SaveExpectedResult(t *testing.T, data []byte, filename string) {
	t.Helper()

	path := filepath.Join(GetTestDataPath(), "expected", filename)
	if err := os.WriteFile(path, data, 0644); err != nil {
		t.Fatalf("Failed to save expected result %s: %v", path, err)
	}
}
