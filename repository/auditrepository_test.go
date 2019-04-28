package repository

import (
	"StaticData/common"
	"StaticData/domain"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type AuditRepositoryTestSuite struct {
	suite.Suite
	repo AuditRepository
	db   *gorm.DB
}

func TestSuiteAuditRepository(t *testing.T) {
	suite.Run(t, new(AuditRepositoryTestSuite))
}

func (suite *AuditRepositoryTestSuite) SetupSuite() {
	db := common.Init()
	db.LogMode(true)
	db = db.Begin()
	suite.repo = GetAuditRepository(db)
	suite.db = db
}

func (suite *AuditRepositoryTestSuite) TearDownSuite() {
	suite.db.Rollback()
	_ = suite.db.Close()

}

func (suite *AuditRepositoryTestSuite) TestSaveAudit() {
	var auditProperties []domain.AuditProperty
	auditProperty := domain.AuditProperty{OldValue: "blub", NewValue: "test", PropertyName: "testProperty"}
	auditProperties = append(auditProperties, auditProperty)

	audit := domain.Audit{Status: "success", ClassName: "Test"}
	audit.AuditProperties = auditProperties
	var audits []domain.Audit
	audits = append(audits, audit)

	auditTransaction := domain.AuditTransaction{PartnerName: "Test"}
	auditTransaction.Audits = audits
	errors := suite.repo.SaveAudit(auditTransaction)
	assert.NotNil(suite.T(), errors)
	assert.NotNil(suite.T(),auditTransaction.ID)
}
