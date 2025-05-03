package cmd

import (
	"fmt"
	"strconv"

	"tasker-cli/app/utils"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(removeCmd)
}

var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove a task",
	Long:  `Remove a task`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 0 {

			id, err := strconv.Atoi(args[0])

			if err != nil {
				fmt.Println("Invalid argument passed")
				return
			}

			tasks := utils.GetAllTask()

			for _, task := range tasks {

				if task.Id == id {
					_, err := utils.RemoveTask(id)

					if err != nil {
						fmt.Printf("Error removing task with Id: %d\n", id)
						return
					}

					fmt.Printf("Task removed successfully - [%d]\n", id)
					return
				}
			}
		}
	},
}
