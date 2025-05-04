package cmd

import (
	"fmt"
	"strconv"

	"tasker-cli/app/constants"
	"tasker-cli/app/utils"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(markDoneCmd)
}

var markDoneCmd = &cobra.Command{
	Use:   "mark-done",
	Short: "Marks a task done",
	Long:  `Marks a task done`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 0 {

			id, err := strconv.Atoi(args[0])

			if err != nil {
				fmt.Println("Invalid arguments passed")
				return
			}

			_, updateErr := utils.UpdateTaskStatus(id, constants.STATUS_DONE)

			if updateErr != nil {
				fmt.Printf("Error updating task -[%d]: %v\n", id, updateErr)
				return
			}

			fmt.Printf("Task updated successfully - [%d]\n", id)
		}
	},
}
