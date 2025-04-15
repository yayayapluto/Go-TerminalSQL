package utils

import (
	"fmt"
	"strconv"
)

func PrintField(len *int, val string) {
	lenStr := "20"
	if len != nil {
		lenStr = strconv.Itoa(*len)
	}

	fmt.Printf("| %-"+lenStr+"s |\n", val)
}
