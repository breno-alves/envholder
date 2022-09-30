package main

import (
	"fmt"
	"os"

	"github.com/breno-alves/envholder/internal/cmd"
)

func main() {
	root := cmd.NewCommandHandler()
	if err := root.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
