package main

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"
)

func TestRun(t *testing.T) {
	var output bytes.Buffer

	err := run(&output)

	if err != nil {
		t.Fatalf("run returned an error: %v", err)
	}

	got := output.String()
	want := "FlowForge\n"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestRunWithFile(t *testing.T) {
	path := filepath.Join(t.TempDir(), "output.txt")

	file, err := os.Create(path)
	if err != nil {
		t.Fatal(err)
	}

	runErr := run(file)
	closeErr := file.Close()

	if runErr != nil {
		t.Fatal(runErr)
	}

	if closeErr != nil {
		t.Fatal(closeErr)
	}

	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}

	if string(data) != "FlowForge\n" {
		t.Errorf("unexpected file content: %q", data)
	}
}
