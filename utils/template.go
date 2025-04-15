package utils

import (
	"fmt"
	"strings"
)

func Separator(symbol string, len int) {
	fmt.Println(strings.Repeat(symbol, len+4))
}
