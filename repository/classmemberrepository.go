package repository

import (
	"StaticData/domain"
	"github.com/jinzhu/gorm"
)

type ClassMemberRepository interface {
	GetMappingByLabel(label string) domain.ClassMemberMapping
}

type classMemberRepoImpl struct {
	DB *gorm.DB
}

func GetClassMemberRepo(DB *gorm.DB) ClassMemberRepository {
	return &classMemberRepoImpl{DB}
}

func (m *classMemberRepoImpl) GetMappingByLabel(label string) domain.ClassMemberMapping {
	var mapping domain.CsvMapping
	//m.DB.Joins("join csv_mapping cm on cm.class_member_mapping_id = class_member_mapping.id").Where("cm.label = ?", label).First(&mapping)
	m.DB.Where("label = ?", label).Preload("ClassMemberMapping").First(&mapping)
	return mapping.ClassMemberMapping
}
