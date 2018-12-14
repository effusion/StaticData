package domain

type UnitShare struct {
	ID int
	UsId string
	SubFundId int32
	UnitShareId int32
	CtrlStatus string
	Isin string
}

func (UnitShare) TableName() string {
	return "unit_share"
}

