package ln

import (
	"encoding/json"
	"os"
)

func Must[T any](x T, err error) T {
	if err != nil {
		Fatal("must failed", Err(err))
	}
	return x
}

func PrintJSON(x any) {
	enc := json.NewEncoder(os.Stderr)
	enc.SetIndent("", "  ")
	err := enc.Encode(x)
	if err != nil {
		Error("PrintJSON failed", Err(err))
	}
}
