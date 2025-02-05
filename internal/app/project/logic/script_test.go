package logic

import (
	"testing"

	"lime/internal/app/project/model"
)

func TestNewVM2(t *testing.T) {
	// Create test CompileInfo
	info := model.CompileInfo{
		// Add necessary test data here
		// Capture stdout to verify the printed output
		// Since NewVM2 prints to stdout, we need to test the output
		CompileBase: model.CompileBase{
			Goos:   model.OS_WINDOWS,
			Goarch: model.ARCH_AMD64,
			Output: "test_output",
		},
	}

	data, err := NewVM2(info)
	if err != nil {
		t.Errorf("NewVM2 failed: %v", err)
	}

	t.Log("data -> ", data)
}
