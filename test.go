package main

// export EDITOR='program' 设置命令行默认编辑器
// export VISUAL='program' 设置GUI默认编辑器

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
	fileName := "test*.txt"

	fp, err := os.CreateTemp("", fileName)
	fmt.Println(fp.Name())
	if err != nil {
		log.Fatal(err)
	}
	defer fp.Close()
	defer func() {
		if os.Remove(fp.Name()).Error() != "" {
			log.Printf("cant not remove tmp file %s, err: %v", fp.Name(), os.Remove(fp.Name()).Error())
		}
	}()
	editorPath := `D:\Program Files\Sublime Text 3\sublime_text.exe`
	editor := `sublime_text.exe`
	cmd := &exec.Cmd{
		Path:   editorPath,
		Args:   []string{editor, fp.Name()},
		Stdin:  os.Stdin,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}
	// cmd := exec.Command(editorPath, fp.Name())
	fmt.Println(cmd.Args)
	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println()
}
