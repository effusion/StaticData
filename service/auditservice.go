package service

import "github.com/jinzhu/gorm"

type AuditService interface {
	CreateAudit(d *interface{})
}

type auditServiceImpl struct {
	DB *gorm.DB
}

func GetAuditService(DB *gorm.DB) AuditService {
	return &auditServiceImpl{DB}
}

func (a *auditServiceImpl) CreateAudit(d *interface{}) {
	panic("implement me")
}
