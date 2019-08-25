package main

import (
	"fmt"
	"go-callpass/callpass"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		argInput := os.Args[1]
		fmt.Println(callpass.GetCallPass(argInput))
	}
}
