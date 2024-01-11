package CLI

import (
	"flag"
)

func Parse() map[string]string {
	repoPath := flag.String("repo", "", "Repository to get pull requests from")
	flag.Parse()

	if *repoPath == "" {
		panic("Please specify a repository to get pulls from")
	}

	return map[string]string {
		"repo": *repoPath,
	}
}
