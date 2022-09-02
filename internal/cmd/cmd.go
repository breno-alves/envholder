package cmd

import (
	"envholder/internal/api"
	"envholder/pkg/config"
	"fmt"
	"github.com/spf13/cobra"
)

type CommandHandler struct {
	root *cobra.Command
	api  *api.Api
}

type Command interface {
	Hello() *cobra.Command
	Login() *cobra.Command
}

func NewCommandHandler() *CommandHandler {
	cmdRoot := &CommandHandler{
		root: &cobra.Command{},
		api:  api.NewApi(),
	}
	cmdRoot.root.AddCommand(Login(cmdRoot))
	cmdRoot.root.AddCommand(ListProjects(cmdRoot))

	return cmdRoot
}

func (ch *CommandHandler) Execute() error {
	return ch.root.Execute()
}
func ListProjects(cmdRoot *CommandHandler) *cobra.Command {
	return &cobra.Command{
		Use:   "project [option]",
		Short: "",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			configData, err := config.ReadConfig()
			if err != nil {
				fmt.Println("missing config file")
				panic(err)
			}
			
			_, err = cmdRoot.api.ListProjects(configData.AccessToken)
			if err != nil {
				panic(err)
			}
		},
	}
}

func Login(cmdRoot *CommandHandler) *cobra.Command {
	return &cobra.Command{
		Use:   "login [username] [password]",
		Short: "",
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			response, err := cmdRoot.api.Login(args[0], args[1])
			if err != nil {
				panic(err)
			}
			fmt.Println("Login successful")
			err = config.WriteConfig(response.AccessToken)
			if err != nil {
				panic(err)
			}
		},
	}
}
