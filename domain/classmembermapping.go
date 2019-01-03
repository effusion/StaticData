package domain

type ClassMemberMapping struct {
	ID     uint
	Class  string
	Member string
}

func (ClassMemberMapping) TableName() string {
	return "class_member_mapping"
}
