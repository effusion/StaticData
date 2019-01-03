package main

import (
	"StaticData/common"
	"StaticData/domain"
	repos "StaticData/repository"
	"StaticData/xlsximporter/staticdata"
	"fmt"
	"gopkg.in/asaskevich/govalidator.v4"
)

func main() {
	common.Init()
	ar := repos.GetCsvMappingRepository(common.GetDB())
	cr := repos.GetClassMemberRepo(common.GetDB())
	imp := staticdata.GetStaticDataImporter(ar, cr)
	imp.ImportStaticDataXlsx("/home/heuby/go/src/StaticData/files/test_e1.xlsx")
	fmt.Printf("Finished!\n")
	//validatorTest()
	//auditTest()

}

func dbtest() {
	//var umbrella domain.Umbrella
	var unitShare domain.UnitShare
	//common.GetDB().Debug().Preload("SubFunds.UnitShares").Preload("SubFunds").First(&umbrella, 101)
	//fmt.Printf("%v", umbrella)
	common.GetDB().Debug().Where("isin = ?", "LU0070848972").Preload("SubFund.Umbrella").Preload("SubFund.Umbrella.Participant").Preload("SubFund").First(&unitShare)
	fmt.Printf("%v", unitShare)

	var partner domain.Partner
	common.GetDB().Debug().Where("id = ?", 197).First(&partner)
	fmt.Printf("%v", partner)

	var csvmapping domain.CsvMapping
	common.GetDB().Debug().Where("label = ?", "ISIN").Preload("ClassMemberMapping").First(&csvmapping)
	fmt.Printf("%v", csvmapping)
}

func validatorTest() {
	var umbrella domain.Umbrella
	umbrella.HomeCountry = "CHE"
	umbrella.CivId = "1234"
	ok, err := govalidator.ValidateStruct(umbrella)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	fmt.Printf("Validation result: %v\n", ok)
}

func auditTest() {
	var audits []*domain.Audit
	common.GetDB().Debug().Where("class_name = ?", "Umbrella").Preload("AuditProperties").Find(&audits)
	for index := range audits {
		audit := audits[index]
		fmt.Printf("Audit: %v\n", audit)
		ap := audit.AuditProperties
		for index1 := range ap {
			fmt.Printf("AuditProperty: %v\n", ap[index1])
		}

	}
}
