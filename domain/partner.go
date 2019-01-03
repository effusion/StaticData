package domain

import "time"

type Partner struct {
	ID                         uint
	Role                       string
	Name                       string
	Street                     string
	ZipCode                    string
	City                       string
	Country                    string
	Phone                      string
	Mail                       string
	Url                        string
	StandardPricePublication   BitBool
	TradeAssociationMembership string
	RecipientsGroup            string
	PriceImportMail            string
	DocumentReportMail         string
	PricePublicationMail       string
	PricePublicationTime       string
	//Umbrella                   []Umbrella `gorm:"ForeignKey:participant_id"`

	UserCreate string
	DateCreate time.Time
	UserUpdate string
	LastUpdate time.Time
}

func (Partner) TableName() string {
	return "partner"
}
