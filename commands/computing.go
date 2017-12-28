package commands

import (
	"github.com/spf13/cobra"
)

func computing() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "computing",
		Short: "Manage computing resources",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	cmd.AddCommand(
		describeInstances(),
		runInstances(),
	)

	return cmd
}
