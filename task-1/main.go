package main

import "fmt"

func main() {
	age := 18
	p := &age
	fmt.Println(p, *p)
}

