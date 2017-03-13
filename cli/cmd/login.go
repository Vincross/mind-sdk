package cmd

import (
	"fmt"
	"mindsdk/cli/mindcli"
	"os"

	"github.com/spf13/cobra"
)

func NewLoginCommand(cli *mindcli.MindCli) *cobra.Command {
	return &cobra.Command{
		Use:   "login [EMAIL] [PASSWORD]",
		Short: "Authenticate yourself",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 2 {
				fmt.Println("Please provide email and password")
				os.Exit(-1)
			}
			err := cli.Login(args[0], args[1])
			if err != nil {
				fmt.Println(err)
				os.Exit(-1)
			}
			fmt.Println("Login succeeded")
		},
	}
}
