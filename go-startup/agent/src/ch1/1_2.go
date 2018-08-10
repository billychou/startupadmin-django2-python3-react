// 1.包申明
package main

// 2. 引入包
import (
	"fmt"
	"os"
)

// 主函数
func main() {
	var b = true
	if b {
		fmt.Printf("welcome")
	}
	// 循环语句
	// 变量
	var s, sep string
	// 循环语句
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = "|"
	}
	fmt.Println(s)
}
