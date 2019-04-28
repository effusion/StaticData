package main

import (
	"StaticData/common"
	repos "StaticData/repository"
	"StaticData/xlsximporter/staticdata"
	"fmt"
)

func main() {
	common.Init()
	ar := repos.GetCsvMappingRepository(common.GetDB())
	cr := repos.GetClassMemberRepo(common.GetDB())
	imp := staticdata.GetStaticDataImporter(ar, cr)
	imp.ImportStaticDataXlsx("/home/heuby/go/src/StaticData/files/test_e1.xlsx")
	fmt.Printf("Finished!\n")
}

