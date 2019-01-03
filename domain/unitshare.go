package domain

type UnitShare struct {
	ID         uint
	UsId       string
	SubFundId  uint
	CtrlStatus string
	Isin       string
	SubFund    SubFund
	Valor      uint
	Name       string
}

func (UnitShare) TableName() string {
	return "unit_share"
}
