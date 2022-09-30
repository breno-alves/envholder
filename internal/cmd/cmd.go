package cmd

import (
	"fmt"
	"log"

	"github.com/breno-alves/envholder/pkg/exporters"
	"github.com/breno-alves/envholder/pkg/transformers"
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
	cmdRoot.root.PersistentFlags().String("format", "dotenv", "Output format expected")

	return cmdRoot
}

func (ch *CommandHandler) Execute() error {
	return ch.root.Execute()
}

func Export(cmdRoot *CommandHandler) *cobra.Command {
	return &cobra.Command{
		Use:   "export [exporter] [path]",
		Short: "",
		Long:  "",
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			export := args[0]
			path := args[1]
			format, _ := cmd.Flags().GetString("format")

			transformer := transformers.NewTransformer(format)
			exporter := exporters.NewExporter(export, path)

			variables, err := exporter.Export()
			if err != nil {
				log.Fatal(err)
			}

			for idx := range variables {
				variable := variables[idx]
				output := transformer.Transform(variable)
				fmt.Printf(output)
			}
		},
	}
}
