package cmd

import (
	"fmt"
	"mindsdk/cli/mindcli"
	"os"

	"github.com/spf13/cobra"
)

func NewXCommand(cli *mindcli.MindCli) *cobra.Command {
	return &cobra.Command{
		Use:   "x [COMMAND] [ARGS..]",
		Short: "Run a command inside of a cross-compiling capable docker container",
		Long: "Run a command inside of a cross-compiling capable docker container.\n\n" +
			"Typically used for cross compiling C/C++ and Golang applications.\n" +
			"Provide name of a build script, run make inside of the current folder or execute\n" +
			"any other linux command inside of a container set up for compiling to the ARM architecture.\n\n" +
			"Examples:\n" +
			"$ mind x ./build.sh`    # Executes build.sh inside container.\n" +
			"$ mind x make`          # Run Makefile inside container.\n" +
			"$ mind x bash`          # Return a shell inside container.\n\n" +
			"For more information about what a docker container is\n" +
			"please see: https://www.docker.com/",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 1 {
				fmt.Println("Need to provide command to execute")
				os.Exit(-1)
			}
			cli.X(args...)
		},
	}
}
