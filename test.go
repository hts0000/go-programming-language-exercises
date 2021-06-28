package main

import (
	"log"
	"net/http"
)

func main() {
	client := &http.Client{}
	// data := `{"number":1}`
	req, err := http.NewRequest("DELETE", "https://api.github.com/repos/hts0000/go-programming-language/issues/10", nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("Authorization", "token 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")
	req.Header.Add("Accept", "application/vnd.github.v3+json")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	resp.Body.Close()
}
