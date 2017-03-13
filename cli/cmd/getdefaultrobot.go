package cmd

import (
	"fmt"
	"mindsdk/cli/mindcli"

	"github.com/spf13/cobra"
)

func NewGetDefaultRobotCommand(cli *mindcli.MindCli) *cobra.Command {
	return &cobra.Command{
		Use:   "get-default-robot",
		Short: "Returns the name of the default robot",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(cli.DefaultRobotName())
		},
	}
}
