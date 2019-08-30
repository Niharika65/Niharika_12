package main

import "fmt"

func swap(x, y int) (int, int) {
	return y, x
}

func main() {
	a, b := swap(2, 3)
	fmt.Println(a, b)
	c,d :=5,6
	c,d=d,c
	fmt.Println(c,d)
}
