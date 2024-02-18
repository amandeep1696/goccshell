package main

import (
	"testing"
)


func TestNewHistory(t *testing.T) {
    h := NewHistory()
    if h == nil {
        t.Error("Expected NewHistory to return a non-nil value")
    }
    if h.file == nil {
        t.Error("Expected NewHistory to open a file for history commands")
    }
}

