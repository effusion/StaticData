package repository

import (
	"StaticData/domain"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

type AuditRepository interface {
	SaveAudit(transaction domain.AuditTransaction)
}

type auditRepositoryImpl struct {
	DB *gorm.DB
}

func GetAuditRepository(DB *gorm.DB) AuditRepository {
	return &auditRepositoryImpl{DB}
}

func (a *auditRepositoryImpl) SaveAudit(transaction domain.AuditTransaction) {
	err := a.DB.Debug().Create(&transaction).Error
	if err != nil {
		log.Error("Error saving audit transaction", err)
	}
}
