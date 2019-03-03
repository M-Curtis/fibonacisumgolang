package main

import "fmt"

func main() {
	x := 4000000
	fmt.Println("Sum of all even numbers in fibonaci sequence up to", x, "=", fibonacisum(x))
}

func fibonacisum(x int) int {
	a, b, c, s := 1, 1, 2, 0
	for ; c <= x; c = a + b {
		//println(c)
		a = b
		b = c
		if c%2 == 0 {
			if c == 2 {
				fmt.Print(c)
			} else {
				fmt.Print("+", c)
			}
			s += c
		}

	}
	fmt.Print("=", s, "\n")
	return (s)
}
