package domain

type Umbrella struct{

}

func (Umbrella) TableName() string {
	return "sub_fund"
}
