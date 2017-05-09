package cmd

import (
	"mindsdk/cli/mindcli"

	"github.com/spf13/cobra"
)

func NewBetatestCommand(cli *mindcli.MindCli) *cobra.Command {
	return &cobra.Command{
		Use:   "beta-test",
		Short: "Install Skill in the robot's app",
		Long: "Run mind beta-test, then mind cli will return an address for installing the remote part of the Skill in the robot's app, a.k.a. the HEXA App. \n" +
			"Open the HEXA App, go to Skill Store -> Scan QR Code -> Enter a link, and type in the address mind cli returned, \n" +
			"then the app will download and install the remote part accordingly",
		Run: func(cmd *cobra.Command, args []string) {
			cli.RunBetatest(args...)
		},
	}
}
