package service

import (
	"github.com/jinzhu/gorm"
)

type SaveService interface {
	SaveDomainWithAudit(d *interface{}, save func())
	SaveDomainsWithAudit(d []*interface{}, save func())
}

type saveServiceImp struct {
	DB *gorm.DB
}

func GetSaveService(DB *gorm.DB) SaveService {
	return &saveServiceImp{DB}
}

func (s *saveServiceImp) SaveDomainWithAudit(d *interface{}, save func()) {
	panic("implement me")
}

func (s *saveServiceImp) SaveDomainsWithAudit(d []*interface{}, save func()) {
	panic("implement me")
}
