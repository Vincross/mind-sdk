package cmd

import (
	"fmt"
	"mindsdk/cli/mindcli"

	"github.com/spf13/cobra"
)

func NewGetDefaultRobotIPCommand(cli *mindcli.MindCli) *cobra.Command {
	return &cobra.Command{
		Use:   "get-default-robot-ip",
		Short: "Returns the IP of the default robot",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(cli.DefaultRobotIP())
		},
	}
}
