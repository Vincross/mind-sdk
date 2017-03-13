package cmd

import (
	"fmt"
	"mindsdk/cli/mindcli"
	"os"

	"github.com/spf13/cobra"
)

func NewScanCommand(cli *mindcli.MindCli) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "scan [OPTIONAL IP]",
		Short: "Scan your network or specific IP for robots",
		Run: func(cmd *cobra.Command, args []string) {
			waitDuration, _ := cmd.Flags().GetInt("waitDuration")
			robots, err := cli.Scan(waitDuration, args...)
			if err != nil {
				fmt.Println(err)
				os.Exit(-1)
			}
			if len(robots) < 1 {
				fmt.Printf("Could not find any robots.\n\n")
				fmt.Println("* Make sure robot is turned on.")
				fmt.Println("* Make sure you are on the same network as robot.")
				fmt.Printf("* Make sure port %d is available.\n", cli.RobotScanner.Config.Port)
				fmt.Println("* Make sure multicast is enabled on your router.")
			}
			for _, robot := range robots {
				fmt.Printf("%s %s\n", robot.IP, robot.Name)
			}
		},
	}
	cmd.Flags().Int("waitDuration", 3, "Time to wait for response from robot")
	return cmd
}
