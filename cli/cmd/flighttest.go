package cmd

import (
	"mindsdk/cli/mindcli"

	"github.com/spf13/cobra"
)

func NewFlightTestCommand(cli *mindcli.MindCli) *cobra.Command {
	return &cobra.Command{
		Use:   "flight-test",
		Short: "Flight test a Skill on mobile device",
		Long: "Open the HEXA app and navigate to : Skill Store -> Scan QR code -> Enter a link,\n" +
			"and type in the address that this command returns.\n" +
			"The HEXA app will then download and install the remote part of the Skill from your/this PC.",
		Run: func(cmd *cobra.Command, args []string) {
			cli.RunFlightTest(args...)
		},
	}
}
