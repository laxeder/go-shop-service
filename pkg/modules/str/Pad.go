package str

import (
	"fmt"
	"strings"
)

// pad end: 4 positions - 0000
func PadVersion(version string) string {
	doc := OnlyNumbers(version)
	padSpaces := fmt.Sprintf("%-4s", doc)
	pad := strings.Replace(padSpaces, " ", "0", 4)
	return pad[:4]
}

// pad start: 4 positions - 0000
func PadYear(year int) string {
	padSpaces := fmt.Sprintf("%4v", year)
	pad := strings.Replace(padSpaces, " ", "0", 4)
	return pad[:4]
}

// pad start: 2 positions - 00
func PadMonth(month int) string {
	padSpaces := fmt.Sprintf("%2v", month)
	pad := strings.Replace(padSpaces, " ", "0", 2)
	return pad[:2]
}

// pad start: 2 positions - 00
func PadDay(day int) string {
	padSpaces := fmt.Sprintf("%2v", day)
	pad := strings.Replace(padSpaces, " ", "0", 2)
	return pad[:2]
}

// pad start: 19 positions - 0000000000000000000
func PadTimestampNano(timestamp int64) string {
	padSpaces := fmt.Sprintf("%19v", timestamp)
	pad := strings.Replace(padSpaces, " ", "0", 19)
	return pad[:19]
}

// pad start: 10 positions : 0000000000
func PadTimestamp(timestamp int64) string {
	padSpaces := fmt.Sprintf("%10v", timestamp)
	pad := strings.Replace(padSpaces, " ", "0", 10)
	return pad[:10]
}
