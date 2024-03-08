package main

import (
	"fmt"
	"strconv"
)

func main() {
	ParseBool, err := strconv.ParseBool("True")
	if err != nil {
		fmt.Println("Error", err.Error())
	} else {
		fmt.Println(ParseBool)
	}

	parseInt, err := strconv.Atoi("1000")
	if err != nil {
		fmt.Println("Error", err.Error())
	} else {
		fmt.Println(parseInt)
	}

	parseString := strconv.Itoa(1000)
	fmt.Println(parseString)

	formatInt := strconv.FormatInt(9999, 2) // 2 is base for binary
	fmt.Println(formatInt)
}
