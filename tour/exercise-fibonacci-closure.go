package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	x := 0
	y := 1
	cnt := 0
	return func() int {
		z := 0
		if cnt == 0 {
			z = x
		} else if cnt == 1 {
			z = y
		} else {
			z = x + y
			x = y
			y = z
		}
		cnt++
		return z
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
