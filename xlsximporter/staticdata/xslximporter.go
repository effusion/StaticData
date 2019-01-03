package staticdata

import (
	"StaticData/domain"
	"StaticData/repository"
	"StaticData/xlsximporter"
	"fmt"
	"github.com/tealeg/xlsx"
	"log"
	"reflect"
	"strings"
)

var mapping []string
var propertyColumnMap map[string]int
var newMapping map[int]domain.ClassMemberMapping

type staticDataImporter struct {
	CsvMappingRepository repository.CsvMappingRepository
	ClassRepository      repository.ClassMemberRepository
}

func GetStaticDataImporter(mappingRepository repository.CsvMappingRepository, classRepository repository.ClassMemberRepository) xlsximporter.Importer {
	return &staticDataImporter{mappingRepository, classRepository}
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
	for rowIndex := firstData; rowIndex <= len(sheet.Rows); rowIndex++ {
		row := sheet.Row(rowIndex)
		if row.Cells != nil {
			var unitShare domain.UnitShare
			var subFund domain.SubFund
			var umbrella domain.Umbrella
			unitShare.SubFund = subFund
			subFund.Umbrella = umbrella
			fmt.Printf("Rownumber: %v\n", rowIndex)
			for index, entry := range newMapping {
				fmt.Printf("%v\t%v\n", index, entry)
				cell := row.Cells[index]
				if cell != nil {
					saveContent(index, cell, entry, unitShare)
				}
			}

		} else {
			break
		}
	}
}

func saveContent(columnNumber int, cell *xlsx.Cell, mapping domain.ClassMemberMapping, unitShare domain.UnitShare) {
	fmt.Printf("Mapping | Class: %v\tMember: %v\t\n", mapping.Class, mapping.Member)
	value, err := cell.FormattedValue()
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	fmt.Printf("%v", value)
	var field reflect.Value
	if mapping.Class == "UnitShare" {
		ps := reflect.ValueOf(&unitShare)
		convertValue(ps, mapping)
	} else if mapping.Class == "SubFund" {
		subFund := unitShare.SubFund
		ps := reflect.ValueOf(&subFund)
		convertValue(ps, mapping)
	} else if mapping.Class == "Umbrella" {
		umbrella := unitShare.SubFund.Umbrella
		ps := reflect.ValueOf(&umbrella)
		convertValue(ps, mapping)
	} else {
		fmt.Printf("No matching class for %v found\n", mapping.Class)
	}

	fmt.Printf("FieldType: %v\n", field)

}

func convertValue(ps reflect.Value, mapping domain.ClassMemberMapping) {
	s := ps.Elem()
	fmt.Printf("Elems: %v\n", s)
	f := s.FieldByName(mapping.Member)
	fmt.Printf("Field: %v\n", f)
	t := f.Kind()
	fmt.Printf("Kind: %v\n", t)

}

//Create the header mapping base on the first row
func (imp *staticDataImporter) createMapping(sheet *xlsx.Sheet) int {
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
	mapping = make([]string, len(firstRow.Cells))
	newMapping = make(map[int]domain.ClassMemberMapping)
	if firstRow != nil {
		for index, cell := range firstRow.Cells {
			if headerMap[cell.Value] != "" {
				classmapping := imp.ClassRepository.GetMappingByLabel(cell.Value)
				newMapping[index] = classmapping
				mapping[index] = headerMap[cell.Value]
			} else {
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
	if strings.HasPrefix(value, "OFST") {
		return 2
	}
	return 1
}
