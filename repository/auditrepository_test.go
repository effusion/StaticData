package repository

import (
	"StaticData/common"
	"StaticData/domain"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/suite"
	"testing"
)

type AuditRepositoryTestSuite struct {
	suite.Suite
	repo AuditRepository
	db   *gorm.DB
}

func TestSaveAudit(t *testing.T) {
	suite.Run(t, new(AuditRepositoryTestSuite))
}

func (s *AuditRepositoryTestSuite) SetupSuite() {
	db := common.Init()
	db.LogMode(true)
	db = db.Begin()
	s.repo = GetAuditRepository(db)
	s.db = db
}

func (s *AuditRepositoryTestSuite) TearDownSuite() {
	s.db.Rollback()
	_ = s.db.Close()

}

func (s *AuditRepositoryTestSuite) TestSaveAudit() {
	var auditProperties []domain.AuditProperty
	auditProperty := domain.AuditProperty{OldValue: "blub", NewValue: "test", PropertyName: "testProperty"}
	auditProperties = append(auditProperties, auditProperty)

	audit := domain.Audit{Status: "success", ClassName: "Test"}
	audit.AuditProperties = auditProperties
	var audits []domain.Audit
	audits = append(audits, audit)

	auditTransaction := domain.AuditTransaction{PartnerName: "Test"}
	auditTransaction.Audits = audits
	s.repo.SaveAudit(auditTransaction)
}
