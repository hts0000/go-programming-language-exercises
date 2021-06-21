package main

import (
	"fmt"
	"issues/github"
	"log"
	"os"
	"time"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	classify(result, 6)
	fmt.Printf("%d issues:\n", result.TotalCount)

}

func classify(result *github.IssuesSearchResult, month float64) {
	fmt.Printf("--------------Less than %.f months--------------\n", month)
	now := time.Now()
	for _, item := range result.Items {
		if m := now.Sub(item.CreatedAt).Hours() / 24 / 30; m <= month {
			fmt.Printf("#%-5d %v %9.9s %.55s\n",
				item.Number, item.CreatedAt, item.User.Login, item.Title)
		}
	}
}
