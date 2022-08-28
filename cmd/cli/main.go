package main

import (
	"envholder/internal/cmd"
)

func main() {
	root := cmd.NewCommandHandler()
	root.Execute()
}
