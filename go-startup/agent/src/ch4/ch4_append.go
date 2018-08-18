package ch4

import (
	"fmt"
	"os"
)

func appendDemo() {
	var runes []rune
	for _,r := range "Hello, 世界" {
		//runes.append(r)
		runes = append(runes, r)
	}
	fmt.Printf("%q\n", runes)
}



