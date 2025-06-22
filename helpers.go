package ln

func Must[T any](x T, err error) T {
	if err != nil {
		Fatal("must failed", Err(err))
	}
	return x
}
