package examples

import (
	"testing"
)

func TestMergeCell(t *testing.T) {
	filePath := "sample/Example.xlsx"
	f, err := createSheet(filePath)
	if err != nil {
		t.Error(err)
	}

	err = f.MergeCell("Sheet1", "A1", "B1")
	if err != nil {
		t.Error(err)
	}
	f.Save()

	mergeCells, err1 := f.GetMergeCells("Sheet1")
	if err1 != nil {
		t.Error(err1)
	}
	if len(mergeCells) != 1 {
		t.Error("セルがマージされていません")
	}
}
