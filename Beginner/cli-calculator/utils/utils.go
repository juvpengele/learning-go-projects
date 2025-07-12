package utils

import (
	"fmt"
	"strings"
)


func ShowResult(key string, result *float64) {
	ShowResultWithBorder(key, result, "=")
}

func ShowResultWithBorder(key string, result *float64, border string) {
	line := strings.Repeat(border, 40)
	fmt.Println(line)
	fmt.Printf("%s %.2f\n", key, *result)
	fmt.Println(line)
}