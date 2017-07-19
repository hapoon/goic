package goic

import (
	"fmt"
)

// String returns string value.
// Same as fmt.Sprintf("%v",v).
func String(v interface{}) string {
	return fmt.Sprintf("%v", v)
}
