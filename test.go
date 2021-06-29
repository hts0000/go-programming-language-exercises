package main

// export EDITOR='program' 设置命令行默认编辑器
// export VISUAL='program' 设置GUI默认编辑器

import (
	"encoding/json"
	"fmt"
	"log"
)

type Issue struct {
	Title   string `json:"title"`
	Number  uint64 `json:"number"`
	State   string `json:"state"`
	Context string `json:"body"` // markdown format
}

func main() {
	var issue = Issue{
		Title:   "fuck world",
		Number:  1,
		State:   "closed",
		Context: "",
	}

	fmt.Println(issue)
	data := make([]byte, 0, 10)
	data, err := json.Marshal(issue)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))
}
