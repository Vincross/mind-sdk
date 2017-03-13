package cmd

import (
	"mindsdk/cli/mindcli"

	"github.com/spf13/cobra"
)

func NewPackCommand(cli *mindcli.MindCli) *cobra.Command {
	return &cobra.Command{
		Use:   "pack",
		Short: "Pack a Skill",
		Run: func(cmd *cobra.Command, args []string) {
			cli.X(append([]string{"mindcli-pack"}, args...)...)
		},
	}
}
