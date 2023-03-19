package main

import (
	"clock"
	"fmt"
)

func main() {

	c := clock.New(-1, 15)
	fmt.Println(int(c))
	fmt.Println(c)
}
