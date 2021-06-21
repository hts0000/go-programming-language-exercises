package main

import (
	"github-issues-cli/github"
)

func main() {
	var is github.Repo
	is.Issues.Add()
	is.Issues.Remove()
	is.Issues.Update()
	is.Issues.Search()
}
