package cmd

import "github.com/spf13/cobra"

var RootCmd = &cobra.Command{
	Use:   "taskmgr",
	Short: "Taskmgr is a CLI task manager",
	Run: func(cmd *cobra.Command, args []string) {

	},
}
