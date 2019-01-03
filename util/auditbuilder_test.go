package util

import (
	"StaticData/common"
	"StaticData/domain"
	log "github.com/sirupsen/logrus"
	"testing"
)

func TestCreateAudit(t *testing.T) {
	common.Init()
	var umb domain.Umbrella
	common.GetDB().Debug().Where("civ_id = ?", "8058").First(&umb)
	umb.Company = "blub"
	at := CreateAudit(umb)
	if at == nil {
		t.Fail()
	}
	result := common.GetDB().Debug().Create(at).Error
	if result != nil {

	}
	common.CloseDB()
}

func init() {
	log.SetReportCaller(true)
	log.SetFormatter(&log.TextFormatter{
		DisableColors: true,
		FullTimestamp: true,
	})
}
