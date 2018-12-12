package domain

type UnitShare struct {
	ID int
	Name string
}

type SubFund struct {
	ID int
	name string
	shares []UnitShare
}

type Umbrella struct {
	ID int
	name string
	funds []SubFund
}

type CsvField struct{
	ID int
	key string
	label string
}
