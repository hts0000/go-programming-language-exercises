package main

import (
	"fmt"
	"go-issues-cli/github"
	"log"
)

func main() {
	repoName := "hts0000/go-programming-language"
	token := "token " + "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"

	repo, err := github.NewRepoer(repoName, token)
	if err != nil {
		log.Fatal(err)
	}

	// 查
	for _, issue := range repo.Issues {
		fmt.Printf("#%-5d %v %9.9s %.55s\n", issue.Number, issue.CreatedAt, issue.User.UserName, issue.Title)
	}

	// 增
	// repo.CreateIssue("cli-test-2", "# cli-test-2")

	// 删
}
