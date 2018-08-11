package main

import (
	"fmt"
	"unicode/utf8"
)

// 这些函数取自于原生字节序列，后续需要查阅go标准库，看go标准库源码
func HasPrefix(s, prefix string) bool {
	// 判断prefix是否为s的前缀
	// 两个条件
	// 1. s的长度大于等于prefix
	// 2. s从0到prefix长度 等于prefix
	return len(s) >= len(prefix) && s[:len(prefix)] == prefix
}

func HasSuffix(s, suffix string) bool {
	// 判断suffix是否为s的后缀
	s_length := len(s) - len(suffix) 
	return len(s) >= len(suffix) && s[s_length:] == suffix
}

func Contains(s, substr string) bool {
	// 判断substr是否为s的子字符串
	// 这个自己还没完全掌握
	for i := 0; i< len(s);i++ {
		if HasPrefix(s[i:], substr) {
			return true
		}
	}
	return false
}

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
	// 
	// string
	// 1. 判断HasPrefix函数
	fmt.Println("welcome", "sksk")
	// 执行的时候，傻逼了。哈哈
	fmt.Println(HasPrefix(`welcome`, `sksks`))
	fmt.Println(HasPrefix(`welcome`, `wel`))
    // 2. 判断HasSuffix
	fmt.Println(HasSuffix(`welcome`, `wel`))
	fmt.Println(HasSuffix(`welcome`, `come`))
 	// 3. 原生utf8编码
	s := "Hello, 世界"
	fmt.Println(len(s))
	fmt.Println(utf8.RuneCountInString(s))
	for i := 0; i<len(s); {
		r, size := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%d\t%c\n", i, r)
		i += size
	}
}
