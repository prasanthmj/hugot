//go:build !darwin

package hugot

func getDefaultLibraryPath() string {
	return ""
}