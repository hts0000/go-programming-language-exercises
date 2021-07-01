package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

// 练习 4.13
// 使用开放电影数据库的JSON服务接口
// 允许你检索和下载 https://omdbapi.com/ 上电影的名字和对应的海报图像
// 编写一个poster工具，通过命令行输入的电影名字，下载对应的海报

// 允许通过电影名称和IMDb编号来查找并下载海报
// IMDb编号：IMDb网站上搜索对应电影，打开该电影界面，在URL中即可找到电影对应的编号
// 如《Back to the Future》这部电影的编号为tt0088763

// 想要正常使用该接口，需要在[https://omdbapi.com/apikey.aspx]网站上申请一个apikey

// 根据名称搜索：http://www.omdbapi.com/?t=Back+to+the+Future&apikey=739c4db6
// 根据IMDb标签搜索：http://www.omdbapi.com/?i=tt0088763&apikey=739c4db6

const URL = "https://omdbapi.com"

type Movie struct {
	Actors     string `json:"Actors"`
	Awards     string `json:"Awards"`
	BoxOffice  string `json:"BoxOffice"`
	Country    string `json:"Country"`
	Dvd        string `json:"DVD"`
	Director   string `json:"Director"`
	Genre      string `json:"Genre"`
	Language   string `json:"Language"`
	Metascore  string `json:"Metascore"`
	Plot       string `json:"Plot"`
	Poster     string `json:"Poster"`
	Production string `json:"Production"`
	Rated      string `json:"Rated"`
	Ratings    []struct {
		Source string `json:"Source"`
		Value  string `json:"Value"`
	} `json:"Ratings"`
	Released   string `json:"Released"`
	Response   string `json:"Response"`
	Runtime    string `json:"Runtime"`
	Title      string `json:"Title"`
	Type       string `json:"Type"`
	Website    string `json:"Website"`
	Writer     string `json:"Writer"`
	Year       string `json:"Year"`
	ImdbID     string `json:"imdbID"`
	ImdbRating string `json:"imdbRating"`
	ImdbVotes  string `json:"imdbVotes"`
}

func main() {
	var (
		movie     Movie
		movieName string
		apiKey    string
		url       string
		id        string
	)
	flag.StringVar(&movieName, "name", "", "the name of the movie")
	flag.StringVar(&apiKey, "key", "", "the key of the omdbapi")
	flag.StringVar(&id, "id", "", "the movie id on the omdb")
	flag.Parse()

	if movieName == "" && id == "" {
		flag.Usage()
		os.Exit(1)
	} else if movieName != "" {
		url = fmt.Sprintf("%s/?t=%s&apikey=%s", URL, strings.Replace(movieName, " ", "+", -1), apiKey)
	} else if id != "" {
		url = fmt.Sprintf("%s/?i=%s&apikey=%s", URL, id, apiKey)
	}

	// fmt.Println(url)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		log.Fatal(resp.StatusCode)
	}

	err = json.NewDecoder(resp.Body).Decode(&movie)
	if err != nil {
		log.Fatal(err)
	}

	resp, err = http.Get(movie.Poster)
	if err != nil {
		log.Fatal(err)
	}

	fp, err := os.OpenFile(movie.Title+".jpg", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer fp.Close()

	n, err := io.Copy(fp, resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("《%s》 poster download success, %v bytes\n", movie.Title, n)
}
