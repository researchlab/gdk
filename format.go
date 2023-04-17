package gdk

import (
	"fmt"
)

// units
var units = []string{"B", "KB", "MB", "GB", "TB", "PB", "EB", "ZB", "YB"}

// BytesToReadable convert bytes to human readable string KB,MB,GB,TB,PB,EB,ZB,YB
func BytesToReadable(bytes float64, precision ...int) string {
	var (
		t      float64 = 1024
		d      float64 = 1
		format string  = "%.0f%s"
	)
	if len(precision) > 0 && precision[0] > 0 {
		format = fmt.Sprintf("%%.%df%%s", precision[0])
	}

	for _, unit := range units {
		if bytes < t {
			return fmt.Sprintf(format, bytes/d, unit)
		}
		d *= 1024
		t *= 1024
	}
	return "TooLarge"
}
