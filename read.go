package read

import (
	"fmt"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
	xlsx "github.com/tealeg/xlsx/v3"
)

func rowVisitor(r *xlsx.Row) error {
	// fmt.Println()
	return r.ForEachCell(cellVisitor)
}

func cellVisitor(c *xlsx.Cell) error {
	_, err := c.FormattedValue()
	if err != nil {
		fmt.Println(err.Error())
		// } else {
		// 	fmt.Print(value, "\t")
		// }
	}
	return err
}

func Method1() {
	// timeStart := time.Now()
	f, err := excelize.OpenFile("pricelist.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Get all the rows in the data.

	rows, _ := f.Rows("data")
	for rows.Next() {
		row, _ := rows.Columns()
		for i := range row {
			// fmt.Print(colCell, "\t")
			if i >= 0 {
				continue
			}
		}
		// fmt.Println()

	}

	// fmt.Println("1)\t", time.Since(timeStart).Seconds(), "detik")
}

func Method2() {
	// timeStart2 := time.Now()
	f2, err := excelize.OpenFile("pricelist.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}

	rows, _ := f2.GetRows("data")
	for _, row := range rows {
		for i := range row {
			// fmt.Print(colCell, "\t")
			if i >= 0 {
				continue
			}
		}
		// fmt.Println()
	}

	// fmt.Println("2)\t", time.Since(timeStart2).Seconds(), "detik")
}

func XlsxModule() {
	// timeStart3 := time.Now()
	wb, err := xlsx.OpenFile("pricelist.xlsx")
	if err != nil {
		fmt.Println(err)
	}

	sh, ok := wb.Sheet["data"]
	if !ok {
		fmt.Errorf("Sheet not found")
	}
	sh.ForEachRow(rowVisitor)

	// fmt.Println("3)\t", time.Since(timeStart3).Seconds(), "detik")
}
