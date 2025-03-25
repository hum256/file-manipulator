package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
)

// ファイルの読み込みと書き込みを共通化
func readFileContent(path string) ([]byte, error) {
	return os.ReadFile(path)
}

func writeFileContent(path string, data []byte) error {
	return os.WriteFile(path, data, 0666)
}

// 文字列を反転する
func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// ファイルの内容を反転し、別ファイルに書き出す
func reverseFile(inputpath, outputpath string) error {
	data, err := readFileContent(inputpath)
	if err != nil {
		return err
	}
	reversed := []byte(reverseString(string(data)))
	return writeFileContent(outputpath, reversed)
}

// ファイルの内容をコピーし、別ファイルに書き出す
func copyFile(inputpath, outputpath string) error {
	data, err := readFileContent(inputpath)
	if err != nil {
		return err
	}
	return writeFileContent(outputpath, data)
}

// ファイルの内容を指定回数繰り返し、上書きする
func duplicateContents(inputpath string, n int) error {
	data, err := readFileContent(inputpath)
	if err != nil {
		return err
	}
	repeated := bytes.Repeat(data, n)
	return writeFileContent(inputpath, repeated)
}

// ファイルの内容から指定文字列を検索し、置換する
func replaceStringInFile(inputpath, needle, newstring string) error {
	data, err := readFileContent(inputpath)
	if err != nil {
		return err
	}
	replaced := bytes.ReplaceAll(data, []byte(needle), []byte(newstring))
	return writeFileContent(inputpath, replaced)
	
}

func main() {	
	args := os.Args
	if len(args) < 4 || len(args) > 5 {
		fmt.Println("コマンドか引数が間違っています。")
		return
	}

	command := args[1]
	var err error

	switch command {
	case "reverse":
		// 引数：reverse、入力ファイルパス、出力ファイルパス
		err = reverseFile(args[2], args[3])
	case "copy":
		// 引数：copy、入力ファイルパス、出力ファイルパス
		err = copyFile(args[2], args[3])
	case "duplicate-contents":
		// 引数：duplicate-contents、入出力ファイルパス、倍数
		var n int
		n, err = strconv.Atoi(args[3])
		if err != nil {
			fmt.Println("倍数の指定が不正です：", err)
			return
		}
		err = duplicateContents(args[2], n)
	case "replace-string":
		// 引数：replace-string、入出力ファイルパス、検索文字列、置換文字列
		if len(args) != 5 {
			fmt.Println("replace-string コマンドはファイルパス、検索文字列、置換文字列を指定してください。")
			return
		}
		err = replaceStringInFile(args[2], args[3], args[4])
	default:
		fmt.Println("不明なコマンドです：", command)
		return
	}

	if err != nil {
		fmt.Println("エラー：", err)
	}
}