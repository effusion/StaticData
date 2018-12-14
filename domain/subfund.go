package domain

type SubFund struct{

}

func (SubFund) TableName() string {
	return "sub_fund"
}