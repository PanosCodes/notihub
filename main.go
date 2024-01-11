package main

import (
	"fmt"

	"notihub.panos.codes/CLI"
)

func main() {
	userInputs := CLI.Parse()
	fmt.Println(userInputs)
}
