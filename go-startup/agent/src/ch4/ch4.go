package main

import (
	"fmt"
)

// 显示传递一个数组的指针给函数，这样在函数内部对数组的任何修改都反应到原始数组上面，下面的程序
func zero(ptr *[32]byte) {
	for i := range ptr {
		ptr[i] = 0
	}
}

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

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
	// slice
	aArray := [...]int{0, 1, 2, 3, 4, 5}
	fmt.Println(aArray)
	reverse(aArray[:])
	fmt.Println(aArray)
	fmt.Println("NEXT")

	ages := map[string]int{
		"alice":   31,
		"charlie": 34,
	}

	fmt.Println(ages["alice"])
	// go使用双引号创建可解析的字符串字面量
	// 反引号用来创建原生的字符串

}
