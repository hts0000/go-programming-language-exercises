package github

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const URL = "https://api.github.com/repos"

type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Issues     []*Issue
}

type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string    // in Markdown format
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

func SearchIssues(userName, repoName string) (*IssuesSearchResult, error) {
	issueURL := URL + "/" + userName + "/" + repoName + "/issues"
	fmt.Println(issueURL)
	resp, err := http.Get(issueURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// We must close resp.Body on all execution paths.
	// (Chapter 5 presents 'defer', which makes this simpler.)
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	var result IssuesSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, nil
}
