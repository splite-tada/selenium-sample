package main

import (
	"flag"
	"fmt"
	"selenium-sample/process"
)

func main() {

	// 実行引数を取得
	flag.Parse()
	arg := flag.Args()[0]

	// 実行時引数によって、実行内容変更
	switch arg {
	case "login":
		process.LoginOnly()
		break
	default:
		fmt.Println("該当なし")
	}
}
