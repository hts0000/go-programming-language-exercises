package main

//练习 7.12
// 修改/list的handler让它把输出打印成一个HTML的表格而不是文本
// html/template包(§4.6)可能会对你有帮助

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type dollars float32
type database map[string]dollars

func main() {
	db := database{
		"Apple":  0.73,
		"Banana": 0.33,
		"Tesla":  23456.78,
	}
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "hello world!\n")
	})
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	log.Fatal(http.ListenAndServe("localhost:8001", nil))
}

func (d dollars) String() string { return fmt.Sprintf("%.2f", d) }

var tplt = template.Must(template.New("test").Parse(`
<!DOCTYPE html>
	<html>
	  <head>
		<title>ex7.9</title>
		  <style>
			table {
			  border-collapse: collapse;
			}
			td, th {
			  border: solid 1px;
			  padding: 0.5em;
			  text-align: right;
			}
		  </style>
	  </head>
	  <body>
		<table>
		  <tr>
			<!-- 在这里设置请求字段 -->
			<th><a>Item</a></th>
			<th><a>Price</a></th>
		  </tr>
		  {{range $key, $value := .}}
		  <tr>
			<td>{{ $key }}</td>
			<td>{{ $value }}</td>
		  </tr>
		  {{end}}
		</table>
	  </body>
	</html>
	`))

func (db database) list(w http.ResponseWriter, req *http.Request) {
	err := tplt.Execute(w, db)
	if err != nil {
		fmt.Fprintf(w, "\n%v", err)
	}
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	fmt.Fprintf(w, "%s\n", price)
}
