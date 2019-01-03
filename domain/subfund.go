package domain

type SubFund struct {
	ID                  uint
	SfId                string
	Version             uint
	DispStatus          string
	ShortName           string
	Name                string
	UnitShares          []UnitShare `gorm:"ForeignKey:sub_fund_id"`
	Umbrella            Umbrella
	UmbrellaId          uint
	InvestmentManagers  string
	ChPayingAgent       string
	Benchmark           string
	IsFundOfFund        bool
	InvestmentObjective string
	AccountingCurrency  string
}

func (SubFund) TableName() string {
	return "sub_fund"
}
