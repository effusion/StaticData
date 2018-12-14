package domain

type CsvMapping struct {
	ID int
	Version int
	Field string
	Kind string
	ParticipantId int
	ParticipantKey int
	Label string
}

func (CsvMapping) TableName() string {
	return "csv_mapping"
}

