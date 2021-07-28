package main

import "fmt"

func main() {
	var b byte = 155
	var signed int8 = int8(b)

	fmt.Println(b)
	fmt.Println(signed)
	fmt.Println(int32(signed))
}
