package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

// 练习 7.11
// 增加额外的handler让客服端可以创建，读取，更新和删除数据库记录
// 例如，一个形如 /update?item=socks&price=6 的请求
// 会更新库存清单里一个货品的价格并且当这个货品不存在或价格无效时返回一个错误值
// （注意：这个修改会引入变量同时更新的问题）

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
	http.HandleFunc("/update", db.update)
	http.HandleFunc("/delete", db.delete)
	http.HandleFunc("/add", db.add)
	http.HandleFunc("/select", db.found)
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

func (db database) update(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price := req.URL.Query().Get("price")

	if _, ok := db[item]; !ok {
		fmt.Fprintf(w, "cant not found %q\n", item)
		w.WriteHeader(http.StatusNotFound)
		return
	} else if item == "" {
		fmt.Fprintf(w, "%q is empty\n", item)
		w.WriteHeader(http.StatusBadRequest)
		return
	} else if price == "" {
		fmt.Fprintf(w, "%q is empty\n", price)
		w.WriteHeader(http.StatusBadRequest)
		return
	} else {
		p, err := strconv.ParseFloat(price, 32)
		if err != nil {
			log.Fatal(err)
		}
		db[item] = dollars(p)
		fmt.Fprintf(w, "update %q = %.2f success", item, p)
	}
}

func (db database) delete(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")

	if _, ok := db[item]; !ok {
		fmt.Fprintf(w, "cant not found %q\n", item)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	delete(db, item)
	fmt.Fprintf(w, "delete %q success\n", item)
}

func (db database) add(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price := req.URL.Query().Get("price")
	p, err := strconv.ParseFloat(price, 32)
	if err != nil {
		fmt.Fprintf(w, "%v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	db[item] = dollars(p)
	fmt.Fprintf(w, "add %q = %.2f success\n", item, p)
}

func (db database) found(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if _, ok := db[item]; !ok {
		fmt.Fprintf(w, "cant not found %q\n", item)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "%s: %s\n", item, db[item])
}
