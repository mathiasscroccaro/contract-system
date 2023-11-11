package repository

type IRepository interface {
	Init(dbUrl string)
	DropAll() error

	ListTenants() ([]domain.Tenant, error)
	GetTenant(key string) (domain.Tenant, error)
	SearchTenant(key string) ([]domain.Tenant, error)
	CreateTenant(tenant domain.Tenant) (domain.Tenant, error)
	UpdateTenant(tenant domain.Tenant) (domain.Tenant, error)
	DeleteTenant(tenantID string) error
	
	ListBuildings() ([]domain.Building, error)
	GetBuilding(buildingID string) (domain.Building, error)
	CreateBuilding(building domain.Building) (domain.Building, error)
	UpdateBuilding(building domain.Building) (domain.Building, error)
	DeleteBuilding(buildingID string) error

	ListProperties(buildingID string) ([]domain.Property, error)
	GetProperty(propertyID string) (domain.Property, error
	CreateProperty(property domain.Property) (domain.Property, error)
	UpdateProperty(property domain.Property) (domain.Property, error)
	DeleteProperty(propertyID string) error

	ListActiveContracts() ([]domain.RentContract, error)
	ListAllContracts() ([]domain.RentContract, error)
	GetContract(contractID string) (domain.RentContract, error)
	CreateContract(contract domain.RentContract) (domain.RentContract, error)
	UpdateContract(contract domain.RentContract) (domain.RentContract, error)
	DeleteContract(contractID string) error

	LisAllPayments() ([]domain.Payment, error)
	GetPayment(paymentID string) (domain.Payment, error)
	CreatePaymentHistory(paymentHistory domain.PaymentHistory) (domain.PaymentHistory, error)
	UpdatePaymentHistory(paymentHistory domain.PaymentHistory) (domain.PaymentHistory, error)
	DeletePaymentHistory(paymentHistoryID string) error
}