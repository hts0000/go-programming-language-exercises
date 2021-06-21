package main

// 练习 4.10
// 修改issues程序，根据问题的时间进行分类
// 比如不到一个月的、不到一年的、超过一年

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

	fmt.Printf("%d issues:\n", result.TotalCount)
	// 返回6个月以内的issues
	classify(result, 6)

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
