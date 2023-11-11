package mongo_adapter

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"os"
	"errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"github.com/mathiasscroccaro/contract-system/internal/domain"
)

func (db *DB) CreateTenant(tenant domain.Tenant) (domain.Tenant, error) {
	_, err := db.MongoInstance.Database(os.Getenv("MONGO_DATABASE")).Collection("tenants").InsertOne(context.Background(), tenant)
	if mongo.IsDuplicateKeyError(err) {
		return domain.Tenant{}, domain.ErrorDuplicatedPrimaryKey
	} else if err != nil {
		return domain.Tenant{}, domain.ErrorInternalServerError
	}
	return tenant, nil
}

func (db *DB) UpdateTenant(tenant domain.Tenant) (domain.Tenant, error) {
	_, err := db.MongoInstance.Database(os.Getenv("MONGO_DATABASE")).Collection("tenants").UpdateOne(context.Background(), bson.M{"_id": tenant.CPF}, bson.M{"$set": tenant})
	if err != nil {
		return domain.Tenant{}, domain.ErrorInternalServerError
	}
	return tenant, nil
}

func (db *DB) GetTenant(cpf string) (domain.Tenant, error) {
	var tenant domain.Tenant
	
	err := db.MongoInstance.Database(os.Getenv("MONGO_DATABASE")).Collection("tenants").FindOne(context.Background(), bson.M{"_id": cpf}).Decode(&tenant)

	if errors.Is(err, mongo.ErrNoDocuments) {
		return domain.Tenant{}, domain.ErrorEntryNotFound
	}

	if err != nil {
		return domain.Tenant{}, domain.ErrorInternalServerError
	}

	return tenant, nil
}

func (db *DB) SearchTenant(key string) ([]domain.Tenant, error) {
	var tenantList []domain.Tenant

	cpfFilter := bson.M{"_id": bson.M{"$regex": primitive.Regex{Pattern: key, Options: ""}}}
	nameFilter := bson.M{"name": bson.M{"$regex": primitive.Regex{Pattern: key, Options: "i"}}}
	rgFilter := bson.M{"rg": bson.M{"$regex": primitive.Regex{Pattern: key, Options: ""}}}

	filter := bson.M{"$or": []bson.M{cpfFilter, nameFilter, rgFilter}}

	cursor, err := db.MongoInstance.Database(os.Getenv("MONGO_DATABASE")).Collection("tenants").Find(context.Background(), filter)

	if err != nil {
		return []domain.Tenant{}, domain.ErrorInternalServerError
	}

	for cursor.Next(context.Background()) {
		var tenant domain.Tenant
		if err := cursor.Decode(&tenant); err != nil {
			return []domain.Tenant{}, domain.ErrorInternalServerError
		}
		tenantList = append(tenantList, tenant)
	}

	return tenantList, nil
}

func (db *DB) ListTenants() ([]domain.Tenant, error) {
	var tenantList []domain.Tenant

	cursor, err := db.MongoInstance.Database(os.Getenv("MONGO_DATABASE")).Collection("tenants").Find(context.Background(), bson.M{})

	if err != nil {
		return []domain.Tenant{}, domain.ErrorInternalServerError
	}

	for cursor.Next(context.Background()) {
		var tenant domain.Tenant
		if err := cursor.Decode(&tenant); err != nil {
			return []domain.Tenant{}, domain.ErrorInternalServerError
		}
		tenantList = append(tenantList, tenant)
	}

	return tenantList, nil
}

func (db *DB) DeleteTenant(cpf string) error {
	count, err := db.MongoInstance.Database(os.Getenv("MONGO_DATABASE")).Collection("tenants").DeleteOne(context.Background(), bson.M{"_id": cpf})

	if count.DeletedCount == 0 {
		return domain.ErrorEntryNotFound
	}

	if err != nil {
		return domain.ErrorInternalServerError
	}

	return nil
}
