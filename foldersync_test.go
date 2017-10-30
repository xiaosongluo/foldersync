package foldersync

import (
	"testing"
)

func TestPathExists(t *testing.T) {
	var path string
	var ok bool
	path = "foldersync.go"
	ok, _ = pathExists(path)
	if !ok {
		t.Errorf("Judgement for existence of the path(%s) makes mistake.", path)
	}
	path = "folder.go"
	ok, _ = pathExists(path)
	if ok {
		t.Errorf("Judgement for existence of the path(%s) makes mistake.", path)
	}
}
