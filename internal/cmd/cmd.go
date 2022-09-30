package cmd

import (
	"fmt"
	"log"

	"github.com/breno-alves/envholder/pkg/ssm"
	"github.com/spf13/cobra"
)

type CommandHandler struct {
	root *cobra.Command
}

func NewCommandHandler() *CommandHandler {
	cmdRoot := &CommandHandler{
		root: &cobra.Command{},
	}
	cmdRoot.root.AddCommand(Export(cmdRoot))

	return cmdRoot
}

func (ch *CommandHandler) Execute() error {
	return ch.root.Execute()
}

func Export(cmdRoot *CommandHandler) *cobra.Command {
	return &cobra.Command{
		Use:   "export [exporter] [path]",
		Short: "",
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			exporter := args[0]
			path := args[1]

			switch exporter {
			case "ssm":
				exp := ssm.NewSSM(path)
				variables, err := exp.ExportVariables(true)
				if err != nil {
					log.Fatal(err)
				}
				for idx := range variables {
					output := fmt.Sprintf("%s=%s\n", variables[idx].Name, variables[idx].Value)
					fmt.Println(output)
				}
			}
		},
	}
}
