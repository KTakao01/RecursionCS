package main

import (
	"fmt"
	"os"
)

func main() {
	message := "What is your favorite food?\n"
	fmt.Print(message)             // 標準出力へ
	fmt.Fprint(os.Stderr, message) // 標準エラー出力へも同時に書き出し
}
