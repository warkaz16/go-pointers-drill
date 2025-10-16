package main

import "fmt"

func main() {
	nums := []int{10, 20, 30}

	fmt.Println("До:", nums)

	p0 := &nums[0]
	*p0 += 5

	fmt.Println("До:", nums)
}
