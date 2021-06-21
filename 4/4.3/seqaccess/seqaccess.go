package main

import (
	"fmt"
	"sort"
)

// 顺序访问map的方式
// 获取map的key，根据排序的key对map进行访问

func main() {
	m := map[string]int{
		"cc": 2,
		"dd": 1,
		"aa": 3,
		"ee": 3,
		"zz": 1,
		"AA": 1,
	}
	keys := func(m map[string]int) []string {
		keys := make([]string, 0, len(m))
		for key := range m {
			keys = append(keys, key)
		}
		return keys
	}(m)

	sort.Strings(keys)

	for _, key := range keys {
		fmt.Println(key, m[key])
	}
}
