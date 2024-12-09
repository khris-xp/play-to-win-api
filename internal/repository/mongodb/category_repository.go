package mongodb

import (
	"context"
	"play-to-win-api/internal/domain"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type categoryRepository struct {
	db   *mongo.Database
	coll *mongo.Collection
}

func NewCategoryRepository(db *mongo.Database) domain.CategoryRepository {
	return &categoryRepository{
		db:   db,
		coll: db.Collection("categories"),
	}
}

func (r *categoryRepository) Create(ctx context.Context, c *domain.Category) error {
	c.CreatedAt = time.Now()
	c.UpdatedAt = time.Now()
	result, err := r.coll.InsertOne(ctx, c)
	if err != nil {
		return err
	}
	c.ID = result.InsertedID.(primitive.ObjectID)
	return nil
}

func (r *categoryRepository) FindByID(ctx context.Context, id string) (*domain.Category, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	var category domain.Category
	err = r.coll.FindOne(ctx, primitive.M{"_id": objectID}).Decode(&category)
	return &category, err
}

func (r *categoryRepository) FindAll(ctx context.Context) ([]domain.Category, error) {
	cursor, err := r.coll.Find(ctx, primitive.M{})
	if err != nil {
		return nil, err
	}
	var categories []domain.Category
	err = cursor.All(ctx, &categories)
	return categories, err
}

func (r *categoryRepository) Update(ctx context.Context, c *domain.Category) error {
	c.UpdatedAt = time.Now()
	_, err := r.coll.UpdateOne(
		ctx,
		primitive.M{"_id": c.ID},
		primitive.M{"$set": c},
	)
	return err
}

func (r *categoryRepository) Delete(ctx context.Context, id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = r.coll.DeleteOne(ctx, primitive.M{"_id": objectID})
	return err
}
