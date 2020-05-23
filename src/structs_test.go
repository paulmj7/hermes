package main

import (
	"testing"
)

func TestConfigStruct(t *testing.T) {
	config := ReadConfig("test.json")

	if config.Port != ":5000" {
		t.Error("The port read is not correct")
	}

	if config.Roots[0].Path != "/usr/tmp/A" {
		t.Error("The first root read is not correct")
	}

	if config.Roots[1].Path != "/usr/tmp/B" {
		t.Error("The second root read is not correct")
	}
}

func TestToString(t *testing.T) {
	config := ReadConfig("test.json")

	rootStrings := ToString(config.Roots)

	if rootStrings[0] != "/usr/tmp/A" {
		t.Error("ToString failed for first root")
	}

	if rootStrings[1] != "/usr/tmp/B" {
		t.Error("ToString failed for second root")
	}
}
