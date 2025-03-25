package main

import (
	"os"
	"testing"
)

// 一時ファイルを作成し、指定の内容を書き込むヘルパー関数
func createTempFile(t *testing.T, content []byte) string {
	t.Helper()
	tmpfile, err := os.CreateTemp("", "testfile")
	if err != nil {
		t.Fatalf("一時ファイル作成エラー: %v", err)
	}
	if err := os.WriteFile(tmpfile.Name(), content, 0666); err != nil {
		t.Fatalf("一時ファイルへの書き込みエラー: %v", err)
	}
	return tmpfile.Name()
}

// テスト終了後、一時ファイルを削除するヘルパー関数
func removeTempFile(t *testing.T, filename string) {
	t.Helper()
	if err := os.Remove(filename); err != nil {
		t.Errorf("一時ファイル削除エラー: %v", err)
	}
}

// reverseString関数のテスト
func TestReverseString(t *testing.T) {
	original := "Hello, 世界"
	expected := "界世 ,olleH"
	got := reverseString(original)
	if got != expected {
		t.Errorf("reverseString() = %q, 期待値 %q", got, expected)
	}
}

// reverseFile関数のテスト
func TestReverseFile(t *testing.T) {
	original := "Hello"
	expected := "olleH"

	inputFile := createTempFile(t, []byte(original))
	defer removeTempFile(t, inputFile)

	// 出力先一時ファイル作成
	outputFile, err := os.CreateTemp("", "testfile_out")
	if err != nil {
		t.Fatalf("出力用一時ファイル作成エラー: %v", err)
	}
	outputFileName := outputFile.Name()
	outputFile.Close()
	defer removeTempFile(t, outputFileName)

	if err := reverseFile(inputFile, outputFileName); err != nil {
		t.Fatalf("reverseFile() エラー: %v", err)
	}
	data, err := os.ReadFile(outputFileName)
	if err != nil {
		t.Fatalf("出力ファイル読み込みエラー: %v", err)
	}
	if string(data) != expected {
		t.Errorf("reverseFile() 出力 = %q, 期待値 %q", string(data), expected)
	}
}

// copyFile関数のテスト
func TestCopyFile(t *testing.T) {
	original := "Copy this content"
	inputFile := createTempFile(t, []byte(original))
	defer removeTempFile(t, inputFile)

	outputFile, err := os.CreateTemp("", "testfile_copy")
	if err != nil {
		t.Fatalf("出力用一時ファイル作成エラー: %v", err)
	}
	outputFileName := outputFile.Name()
	outputFile.Close()
	defer removeTempFile(t, outputFileName)

	if err := copyFile(inputFile, outputFileName); err != nil {
		t.Fatalf("copyFile() エラー: %v", err)
	}
	data, err := os.ReadFile(outputFileName)
	if err != nil {
		t.Fatalf("出力ファイル読み込みエラー: %v", err)
	}
	if string(data) != original {
		t.Errorf("copyFile() 出力 = %q, 期待値 %q", string(data), original)
	}
}

// duplicateContents関数のテスト
func TestDuplicateContents(t *testing.T) {
	original := "dup"
	inputFile := createTempFile(t, []byte(original))
	defer removeTempFile(t, inputFile)

	n := 3
	expected := "dupdupdup"

	if err := duplicateContents(inputFile, n); err != nil {
		t.Fatalf("duplicateContents() エラー: %v", err)
	}

	data, err := os.ReadFile(inputFile)
	if err != nil {
		t.Fatalf("ファイル読み込みエラー: %v", err)
	}
	if string(data) != expected {
		t.Errorf("duplicateContents() 出力 = %q, 期待値 %q", string(data), expected)
	}
}

// replaceStringInFile関数のテスト
func TestReplaceStringInFile(t *testing.T) {
	original := "foo bar foo"
	needle := "foo"
	newstring := "baz"
	expected := "baz bar baz"

	inputFile := createTempFile(t, []byte(original))
	defer removeTempFile(t, inputFile)

	if err := replaceStringInFile(inputFile, needle, newstring); err != nil {
		t.Fatalf("replaceStringInFile() エラー: %v", err)
	}

	data, err := os.ReadFile(inputFile)
	if err != nil {
		t.Fatalf("ファイル読み込みエラー: %v", err)
	}
	if string(data) != expected {
		t.Errorf("replaceStringInFile() 出力 = %q, 期待値 %q", string(data), expected)
	}
}
