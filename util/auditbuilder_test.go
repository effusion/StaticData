package util

import (
	"StaticData/common"
	"StaticData/domain"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type AuditBuilderTestSuite struct {
	suite.Suite
	db *gorm.DB
}

func TestSuiteAuditBuilder(t *testing.T) {
	suite.Run(t, new(AuditBuilderTestSuite))

}

func (suite * AuditBuilderTestSuite) SetupSuite(){
	db := common.Init()
	db.LogMode(true)
	suite.db = db
}

func (suite * AuditBuilderTestSuite) TearDownSuite(){
	_ = suite.db.Close()
}

func (suite *AuditBuilderTestSuite) TestCreateAudit() {
	var umb domain.Umbrella
	suite.db.Debug().Where("civ_id = ?", "8058").First(&umb)
	newValue := "blub"
	umb.Company = newValue
	at := CreateAudit(umb)
	assert.NotNil(suite.T(), at)
	audits := at.Audits
	assert.Equal(suite.T(),1,len(audits))
	audit := audits[0]
	assert.NotNil(suite.T(),audits[0])
	auditProperty := audit.AuditProperties[0]
	assert.Equal(suite.T(), newValue, auditProperty.NewValue,"New value does not match")
}
