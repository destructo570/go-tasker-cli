package cmd

import (
	"fmt"

	"tasker-cli/app/utils"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(showCmd)
}

var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show all or a single task",
	Long:  `Show all or a single task`,
	Run: func(cmd *cobra.Command, args []string) {
		tasks := utils.GetAllTask()
		fmt.Println("# ID   |   Status   |   Description")
		for _, task := range tasks {
			fmt.Printf("%d          %s          %s \n", task.Id, task.Status, task.Title)
		}
	},
}
