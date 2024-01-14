package Sync

import (
	"notihub.panos.codes/Github"
	"notihub.panos.codes/Repository"
)

func Pulls(client *Github.CLI) {
	newPulls := Github.GetPulls(client)
	Repository.BulkInsertPullsIfNotExist(newPulls)
}
