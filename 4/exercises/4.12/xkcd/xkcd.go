package xkcd

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

const URL = "https://xkcd.com/info.0.json"

type Comic struct {
	Month     string `json:"month"`
	Num       uint64 `json:"num"`
	Link      string `json:"link"`
	Year      string `json:"year"`
	News      string `json:"new"`
	SafeTitle string `json:"safe_title"`
	// Transcript string `json:"transcript"`
	Alt   string `json:"alt"`
	Img   string `json:"img"`
	Title string `json:"title"`
	Day   string `json:"day"`
}

type Comics struct {
	Total uint64
	Comic []*Comic
}

func NewComics() (*Comics, error) {
	comics := Comics{}
	comic := Comic{}
	resp, err := require(URL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&comic)
	if err != nil {
		return nil, err
	}

	comics.Total = comic.Num
	var url string
	for i := uint64(1); i < 404; i++ {
		// if i == 404 { // https://xkcd.com/404/info.0.json page not found
		// 	continue
		// }
		url = fmt.Sprintf("https://xkcd.com/%d/info.0.json", i)
		resp, err = require(url)
		if err != nil {
			return nil, err
		}
		fmt.Printf("require %s success\n", url)
		defer resp.Body.Close()

		err = json.NewDecoder(resp.Body).Decode(&comic)
		if err != nil {
			return nil, err
		}
		comics.Comic = append(comics.Comic, &comic)
	}
	return &comics, nil
}

func require(url string) (*http.Response, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("require [%s] failed, error: %v", url, err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("require [%s] failed, http status code: %v", url, resp.StatusCode)
	}
	return resp, nil
}

func (comics *Comics) Search(num string) string {
	n, _ := strconv.Atoi(num)
	return comics.Comic[n-1].Img
}
