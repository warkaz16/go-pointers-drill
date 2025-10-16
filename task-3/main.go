package main

import "fmt"

func main()  {
	price := 199.0
	var pricePtr *float64
	pricePtr = &price

	fmt.Println(*pricePtr)
}
