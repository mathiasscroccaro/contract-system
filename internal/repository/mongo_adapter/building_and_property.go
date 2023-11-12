package mongo_adapter

import (
	"errors"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"os"
	"github.com/mathiasscroccaro/contract-system/internal/domain"
)

func (db *DB) ListBuildings() ([]domain.Building, error) {
	var buildingList []domain.Building

	cursor, err := db.MongoInstance.Database(os.Getenv("MONGO_DATABASE")).Collection("buildings").Find(context.Background(), bson.M{})

	if err != nil {
		return []domain.Building{}, domain.ErrorInternalServerError
	}

	for cursor.Next(context.Background()) {
		var building domain.Building
		if err := cursor.Decode(&building); err != nil {
			return []domain.Building{}, domain.ErrorInternalServerError
		}
		buildingList = append(buildingList, building)
	}

	return buildingList, nil
}

func (db *DB) CreateBuilding(building domain.Building) (domain.Building, error) {
	building.ID = primitive.NewObjectID().Hex()
	
	_, err := db.MongoInstance.Database(os.Getenv("MONGO_DATABASE")).Collection("buildings").InsertOne(context.Background(), building)
	
	if mongo.IsDuplicateKeyError(err) {
		return domain.Building{}, domain.ErrorDuplicatedPrimaryKey
	} else if err != nil {
		return domain.Building{}, domain.ErrorInternalServerError
	}
	
	return building, nil
}

func (db *DB) GetBuilding(buildingID string) (domain.Building, error) {
	var building domain.Building

	err := db.MongoInstance.Database(os.Getenv("MONGO_DATABASE")).Collection("buildings").FindOne(context.Background(), bson.M{"_id": buildingID}).Decode(&building)

	if errors.Is(err, mongo.ErrNoDocuments) {
		return domain.Building{}, domain.ErrorEntryNotFound
	}
	if err != nil {
		return domain.Building{}, domain.ErrorInternalServerError
	}

	return building, nil
}

func (db *DB) UpdateBuilding(building domain.Building) (domain.Building, error) {
	document, err := db.MongoInstance.Database(os.Getenv("MONGO_DATABASE")).Collection("buildings").UpdateOne(context.Background(), bson.M{"_id": building.ID}, bson.M{"$set": building})

	if document.MatchedCount == 0 {
		return domain.Building{}, domain.ErrorEntryNotFound
	}
	if err != nil {
		return domain.Building{}, domain.ErrorInternalServerError
	}

	return building, nil
}

func (db *DB) DeleteBuilding(buildingID string) error {
	count, err := db.MongoInstance.Database(os.Getenv("MONGO_DATABASE")).Collection("buildings").DeleteOne(context.Background(), bson.M{"_id": buildingID})

	if count.DeletedCount == 0 {
		return domain.ErrorEntryNotFound
	}
	if err != nil {
		return domain.ErrorInternalServerError
	}

	return nil
}
