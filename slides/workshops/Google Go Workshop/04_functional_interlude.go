package main
import . "fmt"

func main() {
	print_slice(0, 2, 4, 6, 8)
}

func print_slice(s ...int) {
	for i, v := range s {
		Printf("%v: %v\n", i, v)
	}
}