package main

import (
	"fmt"
)

func main() {
	var i int = 12
	j := i

	fmt.Println(i, j)
	fmt.Printf("%v,%T\n", i, i)
	fmt.Printf("%v,%T\n", j, j)
}
