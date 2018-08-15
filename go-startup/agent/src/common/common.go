package main

import (
	"flag"
	"fmt"
)

func main() {
	logTime := flag.String("time", "", "log time")
	appName := flag.String("app", "", "appname")
	flag.Parse()
	fmt.Println(*logTime)
	fmt.Println(*appName)
}
