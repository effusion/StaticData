package importer

import (
	"StaticData/domain"
	"fmt"
	"github.com/tealeg/xlsx"
	"log"
)

func ImportXlsx()  {

	xlsxFile, err := xlsx.OpenFile("/home/heuby/go/src/StaticData/files/test.xlsx")
	if err !=nil{
		log.Fatal(err)
		return
	}
	sheet := xlsxFile.Sheet["Sheet1"]
	for _, row := range sheet.Rows{
		if row != nil {
			for _, cell := range row.Cells{
				if cell != nil {
					fmt.Println(cell)
					_ = domain.UnitShare{ID: 0, Name: "test"}
				}
			}
		}
	}
}

func addMapping(){

}


func parseHeader(sheet xlsx.Sheet){
	_ = make(map[string]domain.CsvField)
}
