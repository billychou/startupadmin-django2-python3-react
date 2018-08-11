package main

import (
	"fmt"
)

func main() {
	fmt.Printf("本章节学习基本数据类型，包括数字、字符串、布尔值\n")
	var u uint8 = 255
	// 8字节无符号整型
	fmt.Println(u, u+1, u*u)
	//  255, 0, 1
	// u 的类型决定了 u+1 和 u*u 的类型
	fmt.Printf("")
	fmt.Printf("%b\n", 255)
	fmt.Printf("%b\n", 255+1)
	fmt.Printf("%b\n", 255*255) // 1111111000000001

	// fmt
	x := 1<<1 | 1<<5
	var y uint8 = 1<<1 | 1<<2
	fmt.Printf("%08b\n", x)
	fmt.Printf("%08b\n", y)
}
