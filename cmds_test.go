package main

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func TestPwd(t *testing.T) {
    pr, pw := io.Pipe()
    go Pwd(pw)
    buf := new(bytes.Buffer)
    buf.ReadFrom(pr)
    output := buf.String()

    dir, _ := os.Getwd()
    expectedOutput := dir + "\n"
    if output != expectedOutput {
        t.Errorf("Expected Pwd to output '%s', got '%s'", expectedOutput, output)
    }
}
