// 1.包申明
package main

// 2. 引入包
import (
	"fmt"
	"math"
	"os"
)

// 主函数
func main() {
	var b = true
	var u uint8 = 255
	fmt.Println(u, u+1, u*u)

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

	//  整数
	var x uint8 = 1<<1 | 1<<5
	var y uint8 = 1<<1 | 1<<2
	fmt.Printf("%08b\n", x)
	fmt.Printf("%08b\n", y)
	// 浮点数
	// float32 float64
	var f float32 = 16777216
	fmt.Println(f == f+1) // true
	// float 能精确标示的正整数范围有限
	const e = 2.71828
	for x := 0; x < 8; x++ {
		fmt.Printf("x = %d e^x = %8.3f\n", x, math.Exp(float64(x)))
	}
	// 字符串

}
