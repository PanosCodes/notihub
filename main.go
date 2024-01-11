package main

import (
	"fmt"

	"notihub.panos.codes/CLI"
	"notihub.panos.codes/Github"
)

func main() {
	userInputs := CLI.Parse()

	cli := Github.CLI{Repo: userInputs["repo"]}
	pulls := Github.GetPulls(&cli)

	fmt.Println(pulls)
}
