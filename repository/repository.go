package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

type Repository struct {
	db         *gorm.DB
	collection *mongo.Collection
	ctx        context.Context
}

func InitiateRepo(db *gorm.DB, collection *mongo.Collection, ctx context.Context) *Repository {
	return &Repository{
		db:         db,
		collection: collection,
		ctx:        ctx,
	}
}
