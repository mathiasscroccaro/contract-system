package domain

import (
	"time"
)

type Building struct {
	ID         string 
	Address    string
	Reference  string
	Properties []Property
}

type Property struct {
	ID        string
	Reference string
}

type Tenant struct {
	Name                string `bson:"name",omitempty`
	CPF                 string `bson:"_id",omitempty`
	RG                  string `bson:"rg",omitempty`
	PersonalPhone       string `bson:"personal_phone",omitempty`
	Role                string `bson:"role",omitempty`
	ProfessionalAddress string `bson:"professional_address",omitempty`
	ProfessionalPhone   string `bson:"professional_phone",omitempty`
	SpouseName          string `bson:"spouse_name",omitempty`
	SpousePhone         string `bson:"spouse_phone",omitempty`
}

type RentContract struct {
	ID             string
	Property       Property
	Value          uint
	ExtraBills     []ExtraBill
	Tenant         Tenant
	StartDate      time.Time
	EndDate        time.Time
	GuaranteeValue uint
	ExtraNotes     string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type ExtraBill struct {
	ID			string
	Value       uint
	Description string
	PaymentType PaymentType
}

type PaymentType string

const (
	RENT     PaymentType = "ALUGUEL"
	IPTU     PaymentType = "IPTU"
	WATER    PaymentType = "ÁGUA"
	POWER    PaymentType = "LUZ"
	GARANTEE PaymentType = "CAUÇÃO"
	OTHER    PaymentType = "OUTRO"
)

type PaymentHistory struct {
	ID                 string
	ContractID         string
	ReferenceDate      time.Time
	PaymentValue       uint
	PaymentType        PaymentType
	PaymentDescription string
	PropertyID         string
	ValidationUUID     string
	CreatedAt          time.Time
	UpdatedAt          time.Time
}
