package cmd

import (
	"mindsdk/cli/mindcli"

	"github.com/spf13/cobra"
)

func NewInitCommand(cli *mindcli.MindCli) *cobra.Command {
	return &cobra.Command{
		Use:   "init [SKILLNAME] [OPTINAL SKILLID]",
		Short: "Initialize and scaffold a new Skill",
		Run: func(cmd *cobra.Command, args []string) {
			cli.X(append([]string{"mindcli-init"}, args...)...)
		},
	}
}
