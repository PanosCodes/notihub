package main

import (
	"notihub.panos.codes/CLI"
	"notihub.panos.codes/Database"
	"notihub.panos.codes/Github"
	"notihub.panos.codes/Sync"
)

func main() {
	Database.Migrate()
	userInputs := CLI.Parse()

	cli := Github.CLI{Repo: userInputs["repo"]}
	Sync.Pulls(&cli)
}
