package main

import "fmt"

func inc(n *int) {
	*n += 1
}

func inCopy(n int) {
	n += 1
}

func main() {
	x := 10
	fmt.Println("До:", x)
	inc(&x)
	// inCopy(x)
	fmt.Println("После:", x)
}
