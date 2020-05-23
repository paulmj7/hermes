package main

import (
	"testing"
)

func ItemInfoTest(t *testing.T) {
	isFile, size, dateMod, err := ItemInfo("test.json")
	if err != nil {
		t.Error("Error in ItemInfo")
	}

	if !isFile {
		t.Error("Item is a file")
	}

	if size == -1 {
		t.Error("Size should be > 0")
	}

	if dateMod == "" {
		t.Error("Date modified should not be blank")
	}
}
