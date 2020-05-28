package convert

import (
	"fmt"
)

// ReadableSize convert byte to KB,MB,GB,TB,PB
// 目前已知最大的数据存储计算单位是XB，但是具体这个X是什么的缩写满世界都没找到;
// 1B（Byte字节）；
// 1KB（Kilobyte） = 2^10 B = 1024 B；
// 1MB（Megabyte） = 2^10 KB = 1024 KB = 2^20 B；
// 1GB（Gigabyte） = 2^10 MB = 1024 MB = 2^30 B；
// 1TB（Terabyte） = 2^10 GB = 1024 GB = 2^40 B；
// 1PB（Petabyte） = 2^10 TB = 1024 TB = 2^50 B；
// 1EB（Exabyte） = 2^10 PB = 1024 PB = 2^60 B；
// 1ZB（Zettabyte） = 2^10 EB = 1024 EB = 2^70 B；
// 1YB（YottaByte） = 2^10 ZB = 1024 ZB = 2^80 B；
// 1BB（Brontobyte） = 2^10 YB = 1024 YB = 2^90 B；
// 1NB（NonaByte） = 2^10 BB = 1024 BB = 2^100 B；
// 1DB（DoggaByte） = 2^10 NB = 1024 NB = 2^110 B；
// 1CB(Corydonbyte) = 2^10 DB = 1024 DB = 2^120 B；
// 1XB(Xerobyte) = 2^10 CB = 1024 CB = 2^130 B；
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
