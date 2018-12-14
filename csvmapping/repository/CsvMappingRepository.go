package repository

import (
	"StaticData/csvmapping"
	"StaticData/domain"
	"github.com/jinzhu/gorm"
)

type csvMappingRepository struct {
	DB *gorm.DB
}

func GetCsvMappingRepository(DB *gorm.DB) csvmapping.Repository {
	return &csvMappingRepository{DB}
}

func (m *csvMappingRepository) AllByKind(kind string) []*domain.CsvMapping {
	var mappings [] *domain.CsvMapping
	m.DB.Where("kind = ?", kind).Find(&mappings)
	return mappings
}

func (m *csvMappingRepository) AllByKindAndParticipant(kind string, participant int) []*domain.CsvMapping {
	var mappings [] *domain.CsvMapping
	m.DB.Where("kind = ? and participant_id = ? ", kind, participant).Find(&mappings)
	return mappings
}
