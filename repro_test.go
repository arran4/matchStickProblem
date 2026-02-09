package main

import (
	"image/gif"
	"os"
	"os/exec"
	"testing"
)

func TestLimitFlag(t *testing.T) {
	// Clean up any previous test output
	const testOut = "test_out.gif"
	defer os.Remove(testOut)

	// Run main.go with a limit of 5 frames
	cmd := exec.Command("go", "run", "main.go", "-limit", "5", "-out", testOut)
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Logf("Output: %s", output)
		t.Fatalf("Failed to run main.go: %v", err)
	}

	// Read the generated GIF
	f, err := os.Open(testOut)
	if err != nil {
		t.Fatalf("Failed to open test output: %v", err)
	}
	defer f.Close()

	g, err := gif.DecodeAll(f)
	if err != nil {
		t.Fatalf("Failed to decode GIF: %v", err)
	}

	// Verify the number of frames
	// Expected: 1 initial frame + 5 generated frames = 6 frames
	expected := 6
	if len(g.Image) != expected {
		t.Errorf("Expected %d frames, got %d", expected, len(g.Image))
	}
}
