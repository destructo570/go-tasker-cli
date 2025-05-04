package cmd

import (
	"fmt"
	"strconv"

	"tasker-cli/app/constants"
	"tasker-cli/app/utils"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(markPendingCmd)
}

var markPendingCmd = &cobra.Command{
	Use:   "mark-pending",
	Short: "Marks a task pending",
	Long:  `Marks a task pending`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 0 {

			id, err := strconv.Atoi(args[0])

			if err != nil {
				fmt.Println("Invalid arguments passed")
				return
			}

			_, updateErr := utils.UpdateTaskStatus(id, constants.STATUS_PENDING)

			if updateErr != nil {
				fmt.Printf("Error updating task -[%d]: %v\n", id, updateErr)
				return
			}

			fmt.Printf("Task updated successfully - [%d]\n", id)
		}
	},
}
