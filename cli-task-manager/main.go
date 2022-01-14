package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/ekholme/learning-go-projects/cli-task-manager/cmd"
	"github.com/ekholme/learning-go-projects/cli-task-manager/db"
	"github.com/mitchellh/go-homedir"
)

func main() {
	home, _ := homedir.Dir()
	dbPath := filepath.Join(home, "tasks.db")
	must(db.Init(dbPath))
	must(cmd.RootCmd.Execute())
}

//this isn't going to be the best approach for a web application, but for this case it's fine?
func must(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
