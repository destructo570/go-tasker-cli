package cmd

import (
	"fmt"

	"tasker-cli/app/utils"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all or a single task",
	Long:  `List all or a single task`,
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("-------------------------------------------")
		fmt.Println("| ID    |   Status      |    Description  |")
		fmt.Println("-------------------------------------------")

		status := utils.GetArgumentByIndex(args, 0)

		if status == nil {
			listAllTasks("")
		} else {
			listAllTasks(*status)
		}

	},
}

func getStatusTextLabel(status string) string {
	switch status {
	case "done":
		return "Done   "
	case "pending":
		return "Pending"
	default:
		return ""
	}

}

func listAllTasks(status string) {
	tasks := utils.GetAllTask()

	for _, task := range tasks {
		if status != "" && utils.GetSanitizedString(status) != task.Status {
			continue
		}

		fmt.Printf("%d           %s          %s \n", task.Id, getStatusTextLabel(task.Status), task.Title)
	}

}
