package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type CommandHandler struct {
	root *cobra.Command
}

type Command interface {
	Hello() *cobra.Command
	Login() *cobra.Command
}

func NewCommandHandler() *CommandHandler {
	viper.SetConfigFile(".env")
	viper.ReadInConfig()
	fmt.Println(viper.Get("user"))
	cmdRoot := &CommandHandler{
		root: &cobra.Command{},
	}

	cmdRoot.root.AddCommand(Login())
	return cmdRoot
}

func (ch *CommandHandler) Execute() error {
	return ch.root.Execute()
}

func Hello() *cobra.Command {
	return &cobra.Command{
		Use:   "hello [name]",
		Short: "retorna Olá + name passado",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Olá %s\n", args[0])
		},
	}
}

func Login() *cobra.Command {
	return &cobra.Command{
		Use:   "login [username] [password]",
		Short: "",
		//Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {

			fmt.Printf("Login\n")
		},
	}
}
