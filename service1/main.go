package main

import (
	"fmt"
	"github.com/ramilmsh/golangtest/libs/b"
)

func main() {
	fmt.Println((&b.B{B: "b"}).ToA())
}
