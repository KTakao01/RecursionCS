//インターフェースによる抽象化、DIになれるための練習
package main

import(

)
//main,reverseへの切り分け

//具体型の定義
//os.,os.WriteFileのラッパー関数ReadAll,WriteFile

//ioパッケージのfunc ReadAll(r Reader) ([]byte, error)

//ReadAllとWriteFileをメソッドとしてもつ型の定義defaultFileOperations
type defaultFileOperations struct{}

func (d *defaultFileOperations) ReadAll()([]byte,error){
	return io.ReadAll()
}

//ReadAll,WriteFileの抽象化

//インターフェースの導入//メソッドを持つ型の集まりを表す型
type FileReader interface
//便宜的な具体型の導入

//コマンド名による動作の切り分け(共通)

//個別コマンド（ここではreverse)の実装
//入力、出力ファイルの指定が必要

//ファイルをインスタンスとして扱う
//インスタンスを読み込む
//読み込んだデータを逆順にする
//逆順にしたデータを出力用ファイルにかき出し

//データを逆順にする関数


//設計プロセス

//没案1 reverse.go
//mainの中に全ての関数をベタガキする。

//ボツ案2 reverse2.go
//DIせずテストコードを書かない

//ボツ案3 reverse3.go
//パッケージ違いでややこしい。遅い。メソッドが増える（openしないといけない）と面倒。
//osパッケージのfunc Open(name string) (*File, error)
//osパッケージのfunc WriteFile(filename string, data []byte, perm FileMode) error
//func (f *File) Read(b []byte) (n int, err error):File型はReaderインターフェースを実装している
