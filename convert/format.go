package convert

import (
	"fmt"
)

// units
var units = []string{"B", "KB", "MB", "GB", "TB", "PB"}

// ReadableSize convert byte to KB,MB,GB,TB,PB
func ReadableSize(raw float64) string {
	var t float64 = 1024
	var d float64 = 1

	for _, unit := range units {
		if raw < t {
			return fmt.Sprintf("%.1f%s", raw/d, unit)
		}
		d *= 1024
		t *= 1024
	}
	return "TooLarge"
}
