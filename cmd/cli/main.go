package main

import (
	"envholder/internal/cmd"
	"fmt"
	"os"
)

func main() {
	root := cmd.NewCommandHandler()
	if err := root.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
