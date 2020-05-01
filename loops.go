package main

import "fmt"

func main() {
	var str = "some string"
	for pos, char := range str {
		fmt.Println(char, pos)
	}
}
