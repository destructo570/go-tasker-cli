package cmd

import (
	"fmt"
	"strconv"
	"strings"

	"tasker-cli/app/utils"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(updateCmd)
}

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a task",
	Long:  `Update a task`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 0 {

			id, err := strconv.Atoi(args[0])
			newTitle := args[1]

			if err != nil || len(strings.Trim(args[1], " ")) == 0 {
				fmt.Println("Invalid arguments passed")
				return
			}

			_, updateErr := utils.UpdateTask(id, newTitle)

			if updateErr != nil {
				fmt.Printf("Error updating task -[%d]: %v\n", id, updateErr)
				return
			}

			fmt.Printf("Task updated successfully - [%d]\n", id)
		}
	},
}
