//go:build debug || !release

package ln

import (
	"encoding/json"
	"os"
)

// PrintJSON prints value as a JSON on Stderr
// Use only for debugging purposes
func PrintJSON(x any) {
	err := json.NewEncoder(os.Stderr).Encode(x)
	if err != nil {
		Error("PrintJSON failed", Err(err))
	}
}
