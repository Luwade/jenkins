package main

import (
	"fmt"
	"github.com/AsynkronIT/goconsole"
)

func main()  {
	sum_integer := Sum()
	fmt.Println("The sum of two integers are: ", sum_integer)
	console.ReadLine()
}

func Sum() int {
	sum := 1 + 1
	return sum
}