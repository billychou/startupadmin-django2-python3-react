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
	for i := 0; i < len(s); i++ {
		if HasPrefix(s[i:], substr) {
			return true
		}
	}
	return false
}

func basename(s string) string {
	// 将最后一个"/"和之前的部分全部都舍弃
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}

	// 保留最后一个'_'之前的所有内容
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}
	return s
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
	for i := 0; i < len(s); {
		r, size := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%d\t%c\n", i, r)
		i += size
	}
	// 4. range循环
	// range 关键字用于for 循环中迭代数组、切片、通道或者 map 的元素，在数组和切片中它返回元素的索引和对应的值，在集合中反馈 key-value 中的 key 值
	fmt.Println("range")
	for i, r := range "Hello, 世界" {
		fmt.Printf("%d\t%q\t%d\n", i, r, r)
	}
	// %d 整型
	// %q 该值对应的单引号括起来的 go 语法字符字面值，必要的时候会用安全的转义标示
	// %t 单词 true 或者 false， bool 值
	n := 0
	// for _, _ = range s {
	// 	n++
	// }
	// range 循环来处理 9个码点或者文字符号的编码
	for range s {
		n++
	}
	fmt.Println(n) //9
	fmt.Println(string(65))
	// 若将一个整数值转换成字符串，其中的值按照文字符号类型解读，并且产生代表该文字符号值的 utf8编码
	// 如果文字符号值非法，将被专门的替换字符取代，\uFFFD
	fmt.Println(string(12345566))
	fmt.Println(basename("a/b/c.go"))

}
