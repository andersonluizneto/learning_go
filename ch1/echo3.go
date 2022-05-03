// Echo3 show your command lines arguments
package main

import (
	"fmt"
	"os"
	"strings"
)

func main()  {
	fmt.Println(strings.Join(os.Args[1:], " "))
}