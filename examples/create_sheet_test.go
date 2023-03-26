package examples

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	os.MkdirAll("sample", os.ModePerm)
	code := m.Run()
	os.RemoveAll("sample")
	os.Exit(code)
}

// test
func TestCreateSheet(t *testing.T) {
	filePath := "sample/Example.xlsx"
	createSheet(filePath)

	// sample/Example.xlsxにファイルが作成されていることを確認
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		t.Error("ファイルが作成されていません")
	}
}
