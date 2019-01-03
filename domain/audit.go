package domain

import "time"

type AuditTransaction struct {
	ID              uint
	Version         uint
	ArchiveDate     time.Time
	PartnerName     string
	SessionId       string
	Status          string
	TransactionType string
	ContentInId     uint
	ContentOutId    uint
	UserName        string
	DateCreated     time.Time
	LastUpdated     time.Time
	UserCreated     string
	UserUpdated     string
	Audits          []Audit `gorm:"ForeignKey:audit_transaction_id"`
}

func (AuditTransaction) TableName() string {
	return "audit_transaction"
}

type Audit struct {
	ID                 uint
	Version            uint
	AuditTransactionId uint
	AuditTransaction   AuditTransaction
	ClassId            uint
	ClassName          string
	Status             string
	AuditProperties    []AuditProperty `gorm:"ForeignKey:audit_id"`
}

func (Audit) TableName() string {
	return "audit"
}

type AuditProperty struct {
	ID           uint
	Version      uint
	AuditId      uint
	Audit        Audit
	PropertyName string
	NewValue     string
	OldValue     string
}

func (AuditProperty) TableName() string {
	return "audit_property"
}
