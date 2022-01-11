package main

import (
	"fmt"
	"path/filepath"

	"github.com/ekholme/learning-go-projects/cli-task-manager/cmd"
	"github.com/ekholme/learning-go-projects/cli-task-manager/db"
	"github.com/mitchellh/go-homedir"
)

func main() {
	home, _ := homedir.Dir()
	dbPath := filepath.Join(home, "tasks.db")
	err := db.Init(dbPath)
	if err != nil {
		panic(err)
	}
	fmt.Println("db stuff worked")
	cmd.RootCmd.Execute()
}
