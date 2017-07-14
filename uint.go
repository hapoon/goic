package goic

import (
	"fmt"
	"strconv"
)

// Uinter is implemented by any value that has a Uint method,
// which defines the “native” format for that value.
type Uinter interface {
	Uint() uint
}

// Uint returns uint value if convert succeeded.
// If convert failed, this returns error.
func Uint(v interface{}) (uint, error) {
	switch ty := v.(type) {
	case string:
		uv, err := strconv.ParseUint(ty, 10, 64)
		if err != nil {
			return 0, err
		}
		return uint(uv), nil
	case uint, uint8, uint16, uint32, uint64:
		return v.(uint), nil
	case int, int8, int16, int32, int64:
		return uint(v.(int)), nil
	default:
		return 0, fmt.Errorf("type error: %T", v)
	}
}
