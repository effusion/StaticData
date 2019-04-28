package repository

import (
	"StaticData/domain"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

type AuditRepository interface {
	SaveAudit(transaction domain.AuditTransaction) gorm.Errors
}

type auditRepositoryImpl struct {
	DB *gorm.DB
}

func GetAuditRepository(DB *gorm.DB) AuditRepository {
	return &auditRepositoryImpl{DB}
}

func (a *auditRepositoryImpl) SaveAudit(transaction domain.AuditTransaction) gorm.Errors{
	err := a.DB.Debug().Create(&transaction).GetErrors()
	if len(err) > 0 {
		log.Error("Error saving audit transaction", err)
	}
	return err
}
