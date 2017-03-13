package cmd

import (
	"mindsdk/cli/mindcli"

	"github.com/spf13/cobra"
)

func NewUpgradeCommand(cli *mindcli.MindCli) *cobra.Command {
	return &cobra.Command{
		Use:   "upgrade",
		Short: "Upgrades mindcli container to latest version",
		Run: func(cmd *cobra.Command, args []string) {
			cli.UpgradeImage()
		},
	}
}
