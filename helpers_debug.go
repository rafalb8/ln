//go:build debug

package ln

import (
	"encoding/json"
	"os"
)

func PrintJSON(x any) {
	enc := json.NewEncoder(os.Stderr)
	enc.SetIndent("", "  ")
	err := enc.Encode(x)
	if err != nil {
		Error("PrintJSON failed", Err(err))
	}
}
