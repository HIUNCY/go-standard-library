package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Muhamad Zainul Kamal", "Zai"))
	fmt.Println(strings.Split("Muhamad Zainul Kamal", " "))
	fmt.Println(strings.ToLower("Muhamad Zainul Kamal"))
	fmt.Println(strings.ToUpper("Muhamad Zainul Kamal"))
	fmt.Println(strings.Trim("     Muhamad Zainul Kamal    ", " "))
	fmt.Println(strings.ReplaceAll("Muhamad Zainul Kamal", "Zainul", "Ajay"))
}
