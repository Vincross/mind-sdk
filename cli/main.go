package main

import (
	"fmt"
	"mindsdk/cli/cmd"
	"mindsdk/cli/mindcli"
	"os"
	"path/filepath"
	"runtime"

	"github.com/spf13/cobra"
)

func NewDefaultMindCliConfig() *mindcli.MindCliConfig {
	return &mindcli.MindCliConfig{
		Image:             "vincross/mindcli:latest",
		ContainerSkillDir: "/go/src/skill",
		ServeMPKPort:      8888,
		ServeRemotePort:   7597,
	}
}

func NewDefaultRobotScannerConfig() *mindcli.RobotScannerConfig {
	return &mindcli.RobotScannerConfig{
		Message: "VCSCAN",
		Port:    7590,
	}
}

func main() {
	home := homeDir()
	userConfig := mindcli.NewUserConfig(filepath.Join(home, ".mind.json"), filepath.Join(home, ".mind.auth"))
	robotScanner := mindcli.NewRobotScanner(NewDefaultRobotScannerConfig())
	cli := mindcli.NewMindCli(robotScanner, userConfig, NewDefaultMindCliConfig())
	mindCmd := newMindCommand(cli)
	if err := mindCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func newMindCommand(cli *mindcli.MindCli) *cobra.Command {
	var mindCmd = &cobra.Command{
		Use:   "mind",
		Short: "MIND Command-line Interface v0.6.1",
	}
	mindCmd.AddCommand(cmd.NewBuildCommand(cli))
	mindCmd.AddCommand(cmd.NewPackCommand(cli))
	mindCmd.AddCommand(cmd.NewGetDefaultRobotCommand(cli))
	mindCmd.AddCommand(cmd.NewGetDefaultRobotIPCommand(cli))
	mindCmd.AddCommand(cmd.NewInitCommand(cli))
	mindCmd.AddCommand(cmd.NewLoginCommand(cli))
	mindCmd.AddCommand(cmd.NewRunCommand(cli))
	mindCmd.AddCommand(cmd.NewScanCommand(cli))
	mindCmd.AddCommand(cmd.NewSetDefaultRobotCommand(cli))
	mindCmd.AddCommand(cmd.NewUpgradeCommand(cli))
	mindCmd.AddCommand(cmd.NewXCommand(cli))
	mindCmd.AddCommand(cmd.NewFlightTestCommand(cli))
	return mindCmd
}

func homeDir() string {
	if runtime.GOOS == "windows" {
		return os.Getenv("userprofile")
	}
	return os.Getenv("HOME")
}
