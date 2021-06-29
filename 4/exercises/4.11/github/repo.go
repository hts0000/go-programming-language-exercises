package github

import (
	"encoding/json"
	"fmt"
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
	Content   string    `json:"body"` // markdown format
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
	resp, err := getUrl(repoURL)
	if err != nil {
		fmt.Println("###############################")
		return nil, err
	}

	var repo = Repo{
		RepoName: repoName,
		Token:    token,
	}
	err = json.NewDecoder(resp.Body).Decode(&repo)
	if err != nil {
		return nil, err
	}

	issuesURL := repoURL + "/issues?state=all"
	resp, err = getUrl(issuesURL)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}

	err = json.NewDecoder(resp.Body).Decode(&repo.Issues)
	if err != nil {
		return nil, err
	}
	return &repo, nil
}

func (repo *Repo) CreateIssue(data Issue) error {
	return httpRequire("POST", URL+repo.RepoName+"/issues", repo.Token, data)
}

func (repo *Repo) UpdateIssue(data Issue) error {
	return httpRequire("PATCH", URL+repo.RepoName+"/issues/"+fmt.Sprint(data.Number), repo.Token, data)
}

func getUrl(url string) (resp *http.Response, err error) {
	resp, err = http.Get(url)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil,
			fmt.Errorf("require [%s] failed, http status code: [%v]",
				url, resp.StatusCode)
	}
	return
}

func httpRequire(method, url, token string, data Issue) error {
	jsonStr, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("create: conv struct failed, error: %v", err)
	}

	fmt.Println(string(jsonStr))

	req, err := http.NewRequest(method, url, strings.NewReader(string(jsonStr)))
	if err != nil {
		return fmt.Errorf("create request failed: %v, method = [%s], url = [%s]", err, method, url)
	}
	req.Header.Add("Authorization", token)
	req.Header.Add("Accept", "application/vnd.github.v3+json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("request [%s] failed: %v, method = [%s]", url, err, method)
	}
	defer resp.Body.Close()

	if method == "POST" {
		if resp.StatusCode != http.StatusCreated {
			return fmt.Errorf("create issues failed, http status code: [%v]", resp.StatusCode)
		}
	} else {
		if resp.StatusCode != http.StatusOK {
			return fmt.Errorf("request [%s] failed, http state code: %v, method = [%s]", url, resp.StatusCode, method)
		}
	}

	return nil
}
