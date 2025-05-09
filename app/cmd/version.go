package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Tasker",
	Long:  `All software has versions. This is Tasker's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Tasker CLI v0.1.1")
	},
}
