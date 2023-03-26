package examples

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	os.MkdirAll("sample", os.ModePerm)
	code := m.Run()
	os.RemoveAll("sample") // テストの実行結果をエクセルで確認するためには、ここにブレークポイントを設定する
	os.Exit(code)
}
