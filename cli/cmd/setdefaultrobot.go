package cmd

import (
	"fmt"
	"mindsdk/cli/mindcli"
	"os"

	"github.com/spf13/cobra"
)

func NewSetDefaultRobotCommand(cli *mindcli.MindCli) *cobra.Command {
	return &cobra.Command{
		Use:   "set-default-robot [ROBOT NAME]",
		Short: "Set the default robot name",
		Long: "Set the default robot name to be used by `mind run`\n" +
			"Execute `mind scan` to search for available robots",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 1 {
				fmt.Println("Please provide robot name as argument")
				os.Exit(-1)
			}
			err := cli.SetDefaultRobotName(args[0])
			if err != nil {
				fmt.Println(err)
				os.Exit(-1)
			}
		},
	}
}
