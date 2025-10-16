package main

import "fmt"

func main()  {
	temperature := 24

	tPtr := &temperature

	*tPtr += 5

	fmt.Println(*tPtr)
}