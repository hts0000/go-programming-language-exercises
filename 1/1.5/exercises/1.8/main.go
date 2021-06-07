package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		// fmt.Println("Before Comple", url)
		// err := compleUrl1(&url)
		// if err != nil {
		// 	fmt.Fprintf(os.Stderr, "exercises 1.8: comple err: %v\n", err)
		// 	os.Exit(1)
		// }
		// fmt.Println("After Comple", url)

		fmt.Println("Before Comple", url)
		compleUrl2(&url)
		fmt.Println("After Comple", url)
	}
}

func compleUrl1(url *string) error {
	if ok, err := regexp.MatchString("^http://*", *url); err != nil {
		fmt.Fprintf(os.Stderr, "exercises 1.8: regexp err: %v\n", err)
		return err
	} else if !ok {
		*url = "http://" + *url
	}
	return nil
}

func compleUrl2(url *string) {
	if ok := strings.HasPrefix(*url, "http://"); !ok {
		*url = "http://" + *url
	}
}
