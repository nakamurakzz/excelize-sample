package examples

import (
	"os"
	"testing"
)

// test
func TestCreateSheet(t *testing.T) {
	filePath := "sample/Example.xlsx"
	_, err := createSheet(filePath)

	// sample/Example.xlsxにファイルが作成されていることを確認
	if _, err = os.Stat(filePath); os.IsNotExist(err) {
		t.Error("ファイルが作成されていません")
	}
}
