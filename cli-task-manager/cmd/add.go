package cmd

import (
	"fmt"
	"strings"

	"github.com/ekholme/learning-go-projects/cli-task-manager/db"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a task to the task list",
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")
		_, err := db.CreateTask(task)
		if err != nil {
			fmt.Println("Something went wrong:", err)
			return
		}
		fmt.Printf("Added \"%s\" to your task list.\n", task)
	},
}

//init() is a function that gets run before the main function gets run -- they set up things
//that need to be set up before an app runs
func init() {
	RootCmd.AddCommand(addCmd)
}
