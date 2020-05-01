package main

import "fmt"

func main() {
	arr := []int{101, 202, 303}
	fmt.Println(len(arr), cap(arr))

	arr = append(arr, 404)
	fmt.Println(len(arr), cap(arr))

	for idx := range arr {
		fmt.Println(idx, arr[idx])
	}
}
