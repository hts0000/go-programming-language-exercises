package main

// 如何增删改查issues:https://segmentfault.com/a/1190000015144126

import (
	"fmt"
	"github-issues-cli/github"
	"log"
)

const IssuesUrl = "https://api.github.com/search/issues"

func main() {
	userName, repoName := "hts0000", "go-programming-language"
	repo, err := github.SearchIssues(userName, repoName)
	if err != nil {
		log.Fatal(err)
	}
	for _, issue := range repo.Issues {
		fmt.Println(issue.CreatedAt, issue.Number, issue.Body)
	}
}
