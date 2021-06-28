package github

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

const URL = "https://api.github.com/repos/"

type Repo struct {
	RepoName string
	HTMLURL  string `json:"html_url"`
	Issues   []*Issue
	Token    string
}

type Issue struct {
	HTMLURL   string    `json:"html_url"`
	Title     string    `json:"title"`
	Number    uint64    `json:"number"`
	Context   string    `json:"body"` // markdown format
	CreatedAt time.Time `json:"created_at"`
	User      *User     `json:"user"`
	State     string    `json:"state"`
}

type User struct {
	UserName  string `json:"login"`
	UserRepos string `json:"html_url"`
}

func NewRepoer(repoName, token string) (*Repo, error) {
	repoURL := URL + repoName
	issuesURL := repoURL + "/issues?state=all"
	resp, err := http.Get(repoURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil,
			fmt.Errorf("require [%s] failed, http status code: [%v]",
				repoURL, resp.StatusCode)
	}
	var repo = Repo{
		RepoName: repoName,
		Token:    token,
	}
	err = json.NewDecoder(resp.Body).Decode(&repo)
	if err != nil {
		return nil, err
	}

	resp, err = http.Get(issuesURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil,
			fmt.Errorf("require [%s] failed, http status code: [%v]",
				issuesURL, resp.StatusCode)
	}

	err = json.NewDecoder(resp.Body).Decode(&repo.Issues)
	if err != nil {
		log.Fatal(err)
	}
	return &repo, nil
}

func (repo *Repo) CreateIssue(title, context string) error {
	client := &http.Client{}
	data := fmt.Sprintf(`{"title":"%s","body":"%s"}`, title, context)
	req, err := http.NewRequest("POST", URL+repo.RepoName+"/issues", strings.NewReader(data))
	if err != nil {
		return err
	}
	// req.Header.Add("Authorization", "token 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")
	req.Header.Add("Authorization", repo.Token)
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}

func (repo *Repo) DeleteIssue(id uint64) error {

	return nil
}
