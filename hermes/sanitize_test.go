package hermes

import (
	"testing"
)

func TestFromRoot(t *testing.T) {
	if FromRoot("/usr/tmp/A/file1", "/usr/tmp/B") {
		t.Error("Unexpected issue, this statement should not run?")
	}

	if FromRoot("/", "/usr/tmp/A") {
		t.Error("Accessed above root")
	}

	if FromRoot("/usr/tmp/B", "/usr/tmp/A") {
		t.Error("Invalid access")
	}

	if !FromRoot("/usr/tmp/A/file1", "/usr/tmp/A") {
		t.Error("Unexpected issue, comes from root?")
	}
}

func TestSplitPath(t *testing.T) {
	path := "/usr/tmp/A/file1"
	sep := "/"

	curr, parent := SplitPath(path, sep)

	if curr != "file1" {
		t.Error("Incorrect split with current")
	}

	if parent != "/usr/tmp/A" {
		t.Error("Incorrect split with parent")
	}
}
