//go:build !debug

package ln

// PrintJSON prints value as a JSON on Stderr
// Use only for debugging purposes
func PrintJSON(x any) {}
