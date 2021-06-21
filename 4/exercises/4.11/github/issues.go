package github

type Repo struct {
	RepoUrl string
	Token   string
	Issues
}

type Issues struct {
	Title   string
	Context string
	Labels  string
}

func (is *Issues) Add() {

}

func (is *Issues) Remove() {

}

func (is *Issues) Update() {

}

func (is *Issues) Search() {

}
