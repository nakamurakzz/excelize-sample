package examples

import (
	"fmt"

	"github.com/xuri/excelize/v2"
)

func createSheet(filePath string) (*excelize.File, error) {
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
		return nil, err
	}

	// セルの値を設定
	f.SetCellValue("Sheet1", "A1", "Hello world.")

	// ワークブックのデフォルトワークシートを設定
	f.SetActiveSheet(index)

	// 指定されたパスにファイルを保存
	if err := f.SaveAs(filePath); err != nil {
		fmt.Println(err)
	}
	return f, nil
}
