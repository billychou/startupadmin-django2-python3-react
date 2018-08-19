package main

import (
	"fmt"
)

func main() {
	var a [3]int
	fmt.Println(a[0])
	fmt.Println(a[len(a)-1])
	for i, v := range a {
		fmt.Printf("%d %d\n", i, v)
	}

	for _, v := range a {
		fmt.Printf("%d\n", v)
	}

	var q = [3]int{1, 2, 3}
	for i, v := range q {
		fmt.Printf("%d\t%d\n", i, v)
	}
	// 数组长度由初始化数组的元素个数
	qArray := [...]int{1, 2, 3}
	for i, v := range qArray {
		fmt.Printf("%d\t%d\n", i, v)
	}
	// 数组复制
	r := [...]int{3: -1}
	for i, v := range r {
		fmt.Printf("%d\t%d\n", i, v)
	}
}
