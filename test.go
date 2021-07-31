package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"golang.org/x/net/html"
)

// export EDITOR='program' 设置命令行默认编辑器
// export VISUAL='program' 设置GUI默认编辑器

func main() {
	// fileName := "test*.txt"

	// fp, err := os.CreateTemp("", fileName)
	// fmt.Println(fp.Name())
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer fp.Close()
	// defer func() {
	// 	if os.Remove(fp.Name()).Error() != "" {
	// 		log.Printf("cant not remove tmp file %s, err: %v", fp.Name(), os.Remove(fp.Name()).Error())
	// 	}
	// }()
	// editorPath := `D:\Program Files\Sublime Text 3\sublime_text.exe`
	// editor := `sublime_text.exe`
	// cmd := &exec.Cmd{
	// 	Path:   editorPath,
	// 	Args:   []string{editor, fp.Name()},
	// 	Stdin:  os.Stdin,
	// 	Stdout: os.Stdout,
	// 	Stderr: os.Stderr,
	// }
	// // cmd := exec.Command(editorPath, fp.Name())
	// fmt.Println(cmd.Args)
	// err = cmd.Run()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println()

	// cmd := &exec.Cmd{
	// 	Path:   "/usr/bin/vim",
	// 	Args:   []string{filepath.Base("/usr/bin/vim"), "/tmp/test.txt"},
	// 	Stdin:  os.Stdin,
	// 	Stdout: os.Stdout,
	// 	Stderr: os.Stderr,
	// }
	// fmt.Println(cmd.Args)
	// cmd.Run()
	// fmt.Println(cmd.Run().Error())

	// fmt.Println(4398046511618 & (1 << 0))
	// fmt.Printf("%#b\n", 4398046511618)
	// fmt.Printf("%#b\n", 4398046511616)
	x := 4398046511618
	x = x - ((x >> 1) & 0x5555555555555555)
	x = (x & 0x3333333333333333) + ((x >> 2) & 0x3333333333333333)
	x = (x + (x >> 4)) & 0x0f0f0f0f0f0f0f0f
	x = x + (x >> 8)
	x = x + (x >> 16)
	x = x + (x >> 32)
	fmt.Println(int(x & 0x7f))

	n := -100.111
	fmt.Println(+100.111)
	fmt.Println(1 + n)
	t2, _ := time.ParseDuration("-1h")
	fmt.Println(time.Now().Add(t2))

	tempdir, err := ioutil.TempDir("", "test")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(tempdir)
	defer func() {
		if err := os.RemoveAll(tempdir); err != nil {
			log.Print(err)
		}
	}()

	resp, err := http.Get("http://shouce.jb51.net/gopl-zh/ch8/ch8-06.html")
	if err != nil {
		log.Fatal(err)
	}
	doc, err := html.Parse(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	b := &bytes.Buffer{}
	err = html.Render(b, doc)
	if err != nil {
		log.Fatal(err)
	}
	file, err := os.Create("./index.html")
	if err != nil {
		log.Fatal(err)
	}
	_, err = io.Copy(file, b)
	if err != nil {
		log.Fatal(err)
	}
}
