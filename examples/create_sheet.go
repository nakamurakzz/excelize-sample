package examples

import (
	"fmt"

	"github.com/xuri/excelize/v2"
)

func createSheet(filePath string) {
	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	// ワークシートを作成する
	index, err := f.NewSheet("Sheet1")
	if err != nil {
		fmt.Println(err)
		return
	}

	// セルの値を設定
	f.SetCellValue("Sheet1", "A1", "Hello world.")
	f.SetCellValue("Sheet2", "B1", 100)

	// ワークブックのデフォルトワークシートを設定
	f.SetActiveSheet(index)

	// 指定されたパスにファイルを保存
	if err := f.SaveAs(filePath); err != nil {
		fmt.Println(err)
	}
}
