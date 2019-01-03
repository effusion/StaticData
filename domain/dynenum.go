package domain

type DynEnum struct {
	ID       uint
	Category string
	Code     string
	Type     string
}

func (DynEnum) TableName() string {
	return "dyn_enum"
}
