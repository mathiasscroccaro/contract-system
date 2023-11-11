package main

import (
	"fmt"

	"github.com/joho/godotenv"

	"errors"

	"github.com/mathiasscroccaro/contract-system/internal/domain"
	"github.com/mathiasscroccaro/contract-system/internal/repository/mongo_adapter"
)

func main() {
	godotenv.Load()
	mongoInstance := mongo_adapter.GetDBInstance()
	mongoInstance.Init(mongo_adapter.BuildUrlFromEnv())
	// _, err := mongoInstance.UpdateTenant(domain.Tenant{
	// 	Name:                "123",
	// 	CPF:                 "123",
	// 	RG:                  "123",
	// 	Role:                "test",
	// 	PersonalPhone:       "testPhone",
	// 	ProfessionalAddress: "test",
	// 	ProfessionalPhone:   "test",
	// 	SpouseName:          "test",
	// 	SpousePhone:         "test",
	// })
	value, err := mongoInstance.ListTenants()
	if errors.Is(err, domain.ErrorDuplicatedPrimaryKey) {
		fmt.Println("Duplicate key error")
	} else if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Result: ", value)
}
