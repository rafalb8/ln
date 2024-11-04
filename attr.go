package log

import (
	"time"

	"golang.org/x/exp/constraints"
)

// An Attr is a key-value pair.
type Attr struct {
	Key   string
	Value any
}

// Group returns group Attr
func Group(key string, v ...Attr) Attr {
	return Attr{Key: key, Value: v}
}

// Bool returns bool Attr
func Bool(key string, v bool) Attr {
	return Attr{Key: key, Value: v}
}

// Int returns int64 Attr
func Int[T constraints.Signed](key string, v T) Attr {
	return Attr{Key: key, Value: int64(v)}
}

// Uint returns uint64 Attr
func Uint[T constraints.Unsigned](key string, v T) Attr {
	return Attr{Key: key, Value: uint64(v)}
}

// Float returns float64 Attr
func Float[T constraints.Float](key string, v T) Attr {
	return Attr{Key: key, Value: float64(v)}
}

// String returns string Attr
func String(key, value string) Attr {
	return Attr{Key: key, Value: value}
}

// Time returns time Attr
func Time(key string, v time.Time) Attr {
	return Attr{Key: key, Value: v}
}

// Duration returns duration Attr
func Duration(key string, v time.Duration) Attr {
	return Attr{Key: key, Value: v}
}

// Err returns error Attr
func Err(v error) Attr {
	return Attr{Key: "error", Value: v}
}

// Any returns any Attr
//
// Use only if value is of a custom type or if specific Attr does not exist
func Any(key string, value any) Attr {
	return Attr{Key: key, Value: value}
}
