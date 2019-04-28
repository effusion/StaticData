package service

import (
	"StaticData/repository"
	"github.com/jinzhu/gorm"
)

type AuditService interface {
	CreateAudit(d interface{})
}

type auditServiceImpl struct {
	DB *gorm.DB
	repository.AuditRepository
}

func GetAuditService(DB *gorm.DB, auditRepository repository.AuditRepository) AuditService {
	return &auditServiceImpl{DB, auditRepository}
}

func (a *auditServiceImpl) CreateAudit(d interface{}) {
	panic("implement me")
}
