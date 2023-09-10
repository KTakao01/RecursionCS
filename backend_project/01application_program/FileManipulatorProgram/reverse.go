package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	//"Usage: go run file_manipulator.go reverse <inputfile> <outputfile>"

	//コマンドラインから引数を取得する。
	if len(os.Args) < 4 {
		fmt.Println("Usage: go run file_manipulator.go reverse <inputfile> <outputfile>")
		return
	}

	command := os.Args[1]
	inputFile := os.Args[2]
	outputFile := os.Args[3]

	//反転させたいファイル（入力ファイル）を開く

	if command == "reverse" {
		file, err := os.Open(inputFile)
		if err != nil {
			fmt.Println("Error reading file:", err)
			return
		}
		defer file.Close()
		//入力ファイルから読み込みバイトスライスとして返す。

		content, err := io.ReadAll(file)
		if err != nil {
			fmt.Println("Error reading content:", err)
			return
		}
		//バイトスライスをstringに変換する
		reversedData := reverseString(string(content))
		//逆になった読み込みデータを出力ファイルを指定して書き込む。

		err = os.WriteFile(outputFile, []byte(reversedData), 0644)
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}
	}
}

//文字列を逆に変換する関数を作成する
//コードポイント単位(rune)で扱うとよい。文字列型、バイト型はUTF-8エンコードされた場においてUnicode文字はマルチバイトが多いので、不正な文字列が生成されやすい

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
