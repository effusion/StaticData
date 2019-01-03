package domain

import (
	. "StaticData/common"
	"fmt"
	"gopkg.in/asaskevich/govalidator.v4"
	"time"
)

type Umbrella struct {
	ID                      uint
	Version                 uint
	CivId                   string `valid:"matches([0-9]{4}),required"`
	CivType                 string
	LegaStatus              string
	LegalForm               string
	SubFunds                []SubFund `gorm:"ForeignKey:umbrella_id"`
	Participant             Partner   `valid:"partnerValidator~Partner has wrong role"`
	ParticipantId           uint
	DispStatus              string
	HomeCountry             string `valid:"countryValidator~Country not allowed"`
	DepositaryBank          string
	GenericName             string
	MaturityType            string
	IsSelfManaged           BitBool
	OrganisationalStructure string
	EndOfFiscalYear         time.Time
	Auditor                 string
	Company                 string
	UserCreated             string
	UserUpdate              string
	DateCreated             time.Time
	DateUpdate              time.Time
}

func (Umbrella) TableName() string {
	return "umbrella"
}

func init() {
	govalidator.CustomTypeTagMap.Set("countryValidator", func(i interface{}, context interface{}) bool {
		var enum DynEnum
		GetDB().Where("type = ? and code = ?", "Country", i).First(&enum)
		if len(enum.Code) != 0 {
			return true
		}
		return false
	})

	govalidator.CustomTypeTagMap.Set("partnerValidator", func(i interface{}, context interface{}) bool {
		switch v := i.(type) { // type switch on the struct field being validated
		case Partner:
			if v.Role == "Participant" {
				return true
			}
		}
		return false
	})
}

func (Umbrella) CreateAudit() {
	fmt.Printf("")
}
