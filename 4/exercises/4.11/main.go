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
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

func main() {
	var (
		repoName   string
		token      string
		action     string
		editorFlag bool
	)
	flag.StringVar(&repoName, "repo", "", "set repo name, example: hts0000/go-programming-language")
	flag.StringVar(&token, "token", "", "set token value, example: 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")
	flag.StringVar(&action, "a", "search", "set action, support: search, searchall, update, create")
	flag.BoolVar(&editorFlag, "e", false, "whether to use a text editor")
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

	var (
		number  uint64
		state   string
		title   string
		content string
		issue   github.Issue
	)

	switch action {
	case "search":
		fmt.Print("Please input issue number: ")
		fmt.Scanln(&number)
		for _, issue := range repo.Issues {
			if issue.Number == number {
				fmt.Printf("#%-5d %v %9.9s %.55s\n", issue.Number, issue.CreatedAt, issue.User.UserName, issue.Title)
				break
			}
		}
	case "searchall":
		for _, issue := range repo.Issues {
			fmt.Printf("#%-5d %v %9.9s %.55s\n", issue.Number, issue.CreatedAt, issue.User.UserName, issue.Title)
		}
	case "update":
		if editorFlag {
			fmt.Print("Please input issue number, State, Title: ")
			fmt.Scanln(&number, &state, &title)
			fmt.Print("\nRetrieve the content will use editor\n")
			content, err = editor()
			if err != nil {
				log.Fatal(err)
			}
		} else {
			fmt.Print("Please input issue number, State, Title, Context: ")
			fmt.Scanln(&number, &state, &title, &content)
		}
		issue = github.Issue{
			Number:  number,
			State:   state,
			Title:   title,
			Content: content,
		}
		err = repo.UpdateIssue(issue)
		if err != nil {
			log.Fatal(err)
		}
	case "create":
		if editorFlag {
			fmt.Print("Please input issue State, Title: ")
			fmt.Scanln(&state, &title)
			fmt.Print("\nRetrieve the content will use editor\n")
			content, err = editor()
			if err != nil {
				log.Fatal(err)
			}
		} else {
			fmt.Print("Please input issue State, Title, Context: ")
			fmt.Scanln(&state, &title, &content)
		}
		issue = github.Issue{
			State:   state,
			Title:   title,
			Content: content,
		}
		err = repo.CreateIssue(issue)
		if err != nil {
			log.Fatal(err)
		}
	default:
		flag.Usage()
		os.Exit(1)
	}
}

// 无法正常工作
// 打开编辑器后，无法正常写入，提示：另一个程序使用该文件，进程无法访问
func editor() (content string, err error) {
	// fmt.Print("Please input editor path: ")
	// var editorPath string
	// fmt.Scanln(&editorPath)
	editorPath := `D:\Program Files\Sublime Text 3\sublime_text.exe`
	fp, err := os.CreateTemp("", "tmp*.txt")
	if err != nil {
		return "", err
	}
	defer fp.Close()
	defer func() {
		if os.Remove(fp.Name()).Error() != "" {
			content = ""
			err = fmt.Errorf("cant not remove tmp file %s, err: %v", fp.Name(), os.Remove(fp.Name()).Error())
		}
	}()

	cmd := &exec.Cmd{
		Path:   editorPath,
		Args:   []string{editorPath, fp.Name()},
		Stdin:  os.Stdin,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}
	fmt.Println(cmd.Args)

	fmt.Printf("Start editor %s...\n", editorPath)
	err = cmd.Run()
	if err != nil {
		return "", err
	}
	s, err := ioutil.ReadFile(fp.Name())
	content = string(s)
	return
}
