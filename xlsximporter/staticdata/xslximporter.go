package staticdata

import (
	"StaticData/csvmapping"
	"StaticData/xlsximporter"
	"github.com/tealeg/xlsx"
	"log"
	"strings"
)

var mapping []string
var propertyColumnMap map[string] int

type staticDataImporter struct {
	CsvMappingRepository csvmapping.Repository
}

func GetStaticDataImporter(mappingRepository csvmapping.Repository) xlsximporter.Importer {
	return &staticDataImporter{mappingRepository}
}

func (imp *staticDataImporter) ImportStaticDataXlsx(filePath string) (result *xlsximporter.ImportResult) {
	imp.importXlsx(filePath)
	return nil
}

func (imp *staticDataImporter) importXlsx(filePath string) {

	xlsxFile, err := xlsx.OpenFile(filePath)
	if err != nil {
		log.Fatal(err)
		return
	}
	sheet := xlsxFile.Sheet["Sheet1"]
	firstData := imp.createMapping(sheet)
	for rowIndex := firstData ; rowIndex <= len(sheet.Rows); rowIndex++ {
		row := sheet.Row(rowIndex)
		if row != nil {
			for _, cell := range row.Cells {
				if cell != nil {
					saveContent(cell)
				}
			}
		}
	}
}

func saveContent(cell *xlsx.Cell) {

}

//Create the header mapping base on the first row
func (imp *staticDataImporter) createMapping(sheet *xlsx.Sheet) int{
	defaultMappings := imp.CsvMappingRepository.AllByKind("static")
	headerMap := make(map[string]string)
	for _, entry := range defaultMappings {
		headerMap[entry.Label] = entry.Field
	}

	customerMappings := imp.CsvMappingRepository.AllByKindAndParticipant("static", 197)
	for _, entry := range customerMappings {
		headerMap[entry.Label] = entry.Field
	}
	firstRow := sheet.Row(0)
	mapping = make([]string,len(firstRow.Cells))
	if firstRow != nil{
		for index, cell := range firstRow.Cells {
			if headerMap[cell.Value] != "" {
				mapping[index] = headerMap[cell.Value]
			}else{
				mapping[index] = "noMapping"
			}
		}
	}
	propertyColumnMap = make(map[string]int)
	for index, entry := range mapping {
		if "noMapping" != entry {
			propertyColumnMap[entry] = index
		}
	}

	return getFirstDataRow(sheet)
}

func getFirstDataRow(sheet *xlsx.Sheet) int {
	value := sheet.Row(0).Cells[0].Value
	if strings.HasPrefix(value, "OFST"){
		return 2
	}
	return 1
}
