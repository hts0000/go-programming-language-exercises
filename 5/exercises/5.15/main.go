package main

import (
	"fmt"
	"log"
)

// 练习5.15
// 编写类似sum的可变参数函数max和min
// 考虑不传参时，max和min该如何处理，
// 再编写至少接收1个参数的版本

func main() {
	m, err := max1(1, 2, 3, 4, 100, 99)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(m)

	m, err = min1(1, 2, 3, 4, 100, 99)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(m)

	// m, err = min1()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(m)

	m, err = max2(1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(m)

	m, err = min2(1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(m)
}

// 接收任意多参数
func max1(nums ...int) (int, error) {
	if len(nums) == 0 {
		return 0, fmt.Errorf("max: Parameter is null")
	}

	m := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > m {
			m = nums[i]
		}
	}
	return m, nil
}

// 接收任意多参数
func min1(nums ...int) (int, error) {
	if len(nums) == 0 {
		return 0, fmt.Errorf("min: Parameter is null")
	}

	m := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < m {
			m = nums[i]
		}
	}
	return m, nil
}

// 接收至少一个参数
func max2(num int, nums ...int) (int, error) {
	m := num
	for _, n := range nums {
		if n > m {
			m = n
		}
	}
	return m, nil
}

// 接收至少一个参数
func min2(num int, nums ...int) (int, error) {
	m := num
	for _, n := range nums {
		if n < m {
			m = n
		}
	}
	return m, nil
}
