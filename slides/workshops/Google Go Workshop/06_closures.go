package main
import . "fmt"

func main() {
	s := []int{0, 2, 4, 6, 8}
	print_slice(func(i int) int { return s[i] })
}

func print_slice(s interface{}) {
	switch s := s.(type) {
	case func(int) int:
		for i := 0; i < 5; i++ {
			Printf("%v: %v\n", i, s(i))
		}
	}
}