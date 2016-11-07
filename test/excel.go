package main

import (
	"fmt"
	"github.com/tealeg/xlsx"
)

type box struct {
	Phone  string
	level1 string
	level2 string
	level3 string
	level4 string
	level5 string
}

func main() {
	excelFileName := "1468463612.xlsx"
	xlFile, err := xlsx.OpenFile(excelFileName)
	if err != nil {
		fmt.Println(err)
	}

	for _, sheet := range xlFile.Sheets {
		for _, row := range sheet.Rows {
			m1 := make(map[int]string)
			for k, cell := range row.Cells {
				aa, err := cell.String()
				if err != nil {
					fmt.Println(err)
				}
				if aa == "" {
					continue
				}

				m1[k] = aa

			}
			fmt.Println(m1)
		}
	}
}
