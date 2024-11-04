package env

import (
	"cmp"
	"os"
	"reflect"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

type types interface {
	~bool | ~[]string | cmp.Ordered
}

func init() {
	_ = godotenv.Load()
}

// Get retrieves a variable from the environment. If not found, it returns the default value.
func Get[T types](envName string, defaultValue ...T) T {
	var (
		def T
		ret any
		err error
	)

	value := os.Getenv(envName)

	if len(defaultValue) > 0 {
		def = defaultValue[0]
	}

	switch any(def).(type) {
	case string:
		ret = value

	case bool:
		ret, err = strconv.ParseBool(value)

	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		ret, err = strconv.Atoi(value)
		ret = reflect.ValueOf(ret).Convert(reflect.TypeOf(def)).Interface()

	case float32, float64:
		ret, err = strconv.ParseFloat(value, 64)
		ret = reflect.ValueOf(ret).Convert(reflect.TypeOf(def)).Interface()

	case []string:
		if strings.Contains(value, ";") {
			ret = strings.Split(value, ";")
		} else {
			ret = strings.Split(value, ",")
		}
	}

	switch {
	case value == "":
		ret = def
	case err != nil:
		panic(err)
	}

	return ret.(T)
}
