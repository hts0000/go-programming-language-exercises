package main

import (
	"fmt"
	"log"
	"xkcd-cli/xkcd"
)

// 练习 4.12
// 流行的web漫画服务xkcd也提供了JSON接口
// 例如，一个 https://xkcd.com/571/info.0.json 请求将返回一个很多人喜爱的571编号的详细描述
// 下载每个链接（只下载一次）然后创建一个离线索引
// 编写一个xkcd工具，使用这些离线索引，打印和命令行输入的检索词相匹配的漫画的URL

// 获取网站信息：https://xkcd.com/info.0.json
// 其中num为最新图片的编号

func main() {
	comics, err := xkcd.NewComics()
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(comics)

	fmt.Println(comics.Search("403"))
}
