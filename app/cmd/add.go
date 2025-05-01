package cmd

import (
	"fmt"
	"strings"

	"tasker-cli/app/utils"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(addCmd)
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new task",
	Long:  `Add a new task and assign it an ID`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 0 {

			count := 0

			//Iterate over each arg/task provided.
			// Checks if it's not empty spaces before processing by trimming the arg.
			for i := range args {
				task := strings.Trim(args[i], " ")

				if len(task) > 0 {
					count++
					utils.AddTask(task)
					fmt.Println(task)
				}
			}

			fmt.Printf("%d tasks added.\n", count)
		} else {
			fmt.Println("No task provided")
		}
	},
}
