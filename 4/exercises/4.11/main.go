package main

// 练习 4.11：
// 编写一个工具
// 允许用户在命令行创建、读取、更新和关闭GitHub上的issue
// 当必要的时候自动打开用户默认的编辑器用于输入文本信息

// https://github.com/Julineo/golang1training/blob/master/4/4.11/borrowed/main.go

import (
	"flag"
	"fmt"
	"go-issues-cli/github"
	"log"
	"os"
	"runtime"
)

func main() {
	var (
		repoName string
		token    string
	)
	flag.StringVar(&repoName, "repo", "", "set repo name, example: hts0000/go-programming-language")
	flag.StringVar(&token, "token", "", "set token value, example: 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")

	flag.Parse()

	if repoName == "" || token == "" {
		flag.Usage()
		os.Exit(1)
	}

	token = "token " + token

	repo, err := github.NewRepoer(repoName, token)
	if err != nil {
		log.Fatal(err)
	}

	// 查
	for _, issue := range repo.Issues {
		fmt.Printf("#%-5d %v %9.9s %.55s\n", issue.Number, issue.CreatedAt, issue.User.UserName, issue.Title)
	}

	var issue = github.Issue{}
	// 增
	// 将新增的内容存入一个结构体
	issue = github.Issue{
		State:   "open",
		Title:   "hello github-issues-cli - test1",
		Context: "",
	}
	err = repo.CreateIssue(issue)
	if err != nil {
		log.Fatal(err)
	}

	// 改
	// 将修改的内容存入一个结构体
	issue = github.Issue{
		Number:  11,
		State:   "open",
		Title:   "hello github-issues-cli - test2",
		Context: "",
	}
	err = repo.UpdateIssue(issue)
	if err != nil {
		log.Fatal(err)
	}
}

func editor(filePath string) error {
	fp, err := os.CreateTemp("", filePath)
	if err != nil {
		return err
	}
	var editor = ""
	switch runtime.GOOS {
	case "linux":
		editor = "vim"
	}
}
