package main

import (
	"StaticData/common"
	csvMappingRepository "StaticData/csvmapping/repository"
	"StaticData/xlsximporter/staticdata"
)

func main() {
	common.Init()
	ar := csvMappingRepository.GetCsvMappingRepository(common.GetDB())
	imp := staticdata.GetStaticDataImporter(ar)
	imp.ImportStaticDataXlsx("/home/heuby/go/src/StaticData/files/test.xlsx")

}
