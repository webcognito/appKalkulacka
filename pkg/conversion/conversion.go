// Converting different types
package conversion

import (
	"fmt"
	"strconv"
)

func StringToFloat64(s string) (f float64) {
	if s == "" {
		return
	} else {
		f, err := strconv.ParseFloat(s, 64)
		if err != nil {
			fmt.Println(err)
		}
		return f
	}
}
