package service

import (
	"StaticData/common"
	"StaticData/domain"
	"StaticData/repository/mocks"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/suite"
	"testing"
)

type AuditServiceTestSuite struct {
	suite.Suite
	service AuditService
	db      *gorm.DB
}


func TestSuiteAuditService(t *testing.T) {
	suite.Run(t, new(AuditServiceTestSuite))
}

func (suite *AuditServiceTestSuite) SetupSuite() {
	db := common.Init()
	db.LogMode(true)
	db = db.Begin()
	suite.service = GetAuditService(db,&mocks.AuditRepository{})
	suite.db = db

}

func (suite *AuditServiceTestSuite) TearDownSuite() {
	suite.db.Rollback()
	_ = suite.db.Close()
}

func (suite * AuditServiceTestSuite) TestCreateAudit(){
	auditTransaction := domain.AuditTransaction{PartnerName: "Test"}
	suite.service.CreateAudit(&auditTransaction)
}