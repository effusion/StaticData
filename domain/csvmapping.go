package domain

type CsvMapping struct {
	ID                   int
	Version              int
	Field                string
	Kind                 string
	ParticipantId        int
	ParticipantKey       int
	Label                string
	ClassMemberMapping   ClassMemberMapping
	ClassMemberMappingId uint
}

func (CsvMapping) TableName() string {
	return "csv_mapping"
}
