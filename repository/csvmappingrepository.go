package repository

import (
	"StaticData/domain"
	"github.com/jinzhu/gorm"
)

type CsvMappingRepository interface {
	// blbubb
	AllByKind(kind string) []*domain.CsvMapping

	//tjrejejej
	AllByKindAndParticipant(kind string, participant int) []*domain.CsvMapping
}

type csvMappingRepositoryImpl struct {
	DB *gorm.DB
}
// Creates a new
func GetCsvMappingRepository(DB *gorm.DB) CsvMappingRepository {
	return &csvMappingRepositoryImpl{DB}
}

func (m *csvMappingRepositoryImpl) AllByKind(kind string) []*domain.CsvMapping {
	var mappings []*domain.CsvMapping
	m.DB.Where("kind = ?", kind).Find(&mappings)
	return mappings
}

func (m *csvMappingRepositoryImpl) AllByKindAndParticipant(kind string, participant int) []*domain.CsvMapping {
	var mappings []*domain.CsvMapping
	m.DB.Where("kind = ? and participant_id = ? ", kind, participant).Find(&mappings)
	return mappings
}
