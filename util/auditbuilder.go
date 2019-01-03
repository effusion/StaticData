package util

import (
	"StaticData/common"
	"StaticData/domain"
	"fmt"
	log "github.com/sirupsen/logrus"
	"reflect"
	"strings"
)

func CreateAudit(new interface{}) *domain.AuditTransaction {
	at := domain.AuditTransaction{PartnerName: "Test", Version: 0}
	var audits []domain.Audit
	var properties []domain.AuditProperty
	e := reflect.ValueOf(new)
	structType := e.Type()
	if structType.Kind() != reflect.Struct {
		log.Fatal("Audit Type is not a struct.")
		return nil
	}
	id := e.FieldByName("ID").Interface()
	determineTransactionType(id, &at)
	c := getCurrentInstance(structType.Name(), id.(uint))
	cv := reflect.ValueOf(c)
	className := fmt.Sprintf("%v", structType)
	className = className[7:]
	audit := domain.Audit{ClassId: 0, ClassName: className, Status: "success", Version: 0}

	for i := 0; i < e.NumField(); i++ {
		fieldType := e.Type().Field(i).Type
		fieldName := e.Type().Field(i).Name
		k := fieldType.Kind()
		fn := fieldType.String()
		if k != reflect.Struct && !(k == reflect.Slice && strings.Contains(fn, "domain.")) {
			newValue := fmt.Sprintf("%v", e.Field(i).Interface())
			oldValue := fmt.Sprintf("%v", cv.Field(i).Interface())
			if newValue != oldValue {
				properties = append(properties, domain.AuditProperty{PropertyName: fieldName, NewValue: newValue, OldValue: oldValue, Version: 0})
			}
		}
	}
	audit.AuditProperties = properties
	audits = append(audits, audit)
	at.Audits = audits
	return &at
}

func determineTransactionType(id interface{}, at *domain.AuditTransaction) {
	if id == nil {
		at.TransactionType = "create"
	} else {
		at.TransactionType = "modify"
	}
}

func CreateAudits(new []interface{}) {

}

func getCurrentInstance(name string, id uint) interface{} {

	switch name {
	case "Umbrella":
		var umb domain.Umbrella
		common.GetDB().Debug().Where("id = ?", id).First(&umb)
		return umb
	}

	return nil
}
