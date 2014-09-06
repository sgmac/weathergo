package helpers

import (
	"fmt"
	"time"
)

func DoLine(line string) string {

	var header string
	header = "%s\n"

	for i := 0; i < len(line); i++ {
		header += "="
	}
	header += "\n"
	return header
}

func EpochToTime(epoch int64) string {
	hour, min, _ := time.Unix(epoch, 0).Clock()
	return fmt.Sprintf("%d:%d", hour, min)
}
