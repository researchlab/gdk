package convert

import (
	"fmt"
)

// ReadableSize convert byte to KB,MB,GB,TB,PB
func ReadableSize(raw float64) string {
	var t float64 = 1024
	var d float64 = 1

	if raw < t {
		return fmt.Sprintf("%.1fB", raw/d)
	}

	d *= 1024
	t *= 1024

	if raw < t {
		return fmt.Sprintf("%.1fKB", raw/d)
	}

	d *= 1024
	t *= 1024

	if raw < t {
		return fmt.Sprintf("%.1fMB", raw/d)
	}

	d *= 1024
	t *= 1024

	if raw < t {
		return fmt.Sprintf("%.1fGB", raw/d)
	}

	d *= 1024
	t *= 1024

	if raw < t {
		return fmt.Sprintf("%.1fTB", raw/d)
	}

	d *= 1024
	t *= 1024

	if raw < t {
		return fmt.Sprintf("%.1fPB", raw/d)
	}

	return "TooLarge"
}
