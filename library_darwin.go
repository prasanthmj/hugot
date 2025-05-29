//go:build darwin

package hugot

import (
	"os"
	"path/filepath"
)

func getDefaultLibraryPath() string {
	possiblePaths := []string{
		"/usr/local/lib/libonnxruntime.dylib",
		"/opt/homebrew/lib/libonnxruntime.dylib",
		filepath.Join(os.Getenv("HOME"), ".local/lib/libonnxruntime.dylib"),
		"libonnxruntime.dylib",
	}
	
	for _, path := range possiblePaths {
		if _, err := os.Stat(path); err == nil {
			return path
		}
	}
	
	return "libonnxruntime.dylib"
}