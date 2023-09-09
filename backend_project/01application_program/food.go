package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	data := []byte("What is your favorite food?\n")
	//stdoutに直接バイトを書き込む
	os.Stdout.Write(data)
	os.Stdout.Sync() //stdoutをフラッシュ

	reader := bufio.NewReader(os.Stdin)
	food, _ := reader.ReadString('\n')

	food = food[:len(food)-1]
	fmt.Printf("Thanks for letting me know your favorite food is %s.\n", food)
}
