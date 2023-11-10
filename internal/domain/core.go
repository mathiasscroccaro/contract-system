package domain

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
	ID                  string
	Name                string
	CPF                 string
	RG                  string
	PersonalPhone       string
	Role                string
	ProfessionalAddress string
	ProfessionalPhone   string
	SpouseName          string
	SpousePhone         string
}

type RentContract struct {
	ID             string
	PropertyID     string
	Value          uint
	ExtraBills     []ExtraBill
	TenantID       string
	StartDate      time.Time
	EndDate        time.Time
	GuaranteeValue uint
	ExtraNotes     string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type ExtraBill struct {
	ID          string
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
