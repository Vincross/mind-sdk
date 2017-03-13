package cmd

import (
	"mindsdk/cli/mindcli"

	"github.com/spf13/cobra"
)

func NewBuildCommand(cli *mindcli.MindCli) *cobra.Command {
	return &cobra.Command{
		Use:   "build",
		Short: "Build a Skill",
		Run: func(cmd *cobra.Command, args []string) {
			cli.X(append([]string{"mindcli-build"}, args...)...)
		},
	}
}
