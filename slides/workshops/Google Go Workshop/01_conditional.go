package main

import . "fmt"

func main() {
	s := []int{0, 2, 4, 6, 8}
	for i := 0; i < len(s); i++ {
		Printf("%v: %v\n", i, s[i])
	}
}