package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Quantidade de processadores:", runtime.NumCPU())
}
