package cmd

import (
	"mindsdk/cli/mindcli"

	"github.com/spf13/cobra"
)

func NewRunCommand(cli *mindcli.MindCli) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "run [OPTIONAL ROBOT NAME]",
		Short: "Run Skill on robot",
		Long: "Run Skill on robot.\n" +
			"If robot name is not provided, `mind` will use the default robot",
		Run: func(cmd *cobra.Command, args []string) {
			noInstall, _ := cmd.Flags().GetBool("noInstall")
			cli.RunSkill(noInstall, args...)
		},
	}
	cmd.Flags().Bool("noInstall", false, "Run skill without installing.")
	return cmd
}
