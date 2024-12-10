package mongodb

import (
	"context"
	"play-to-win-api/internal/domain"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type productRepository struct {
	db   *mongo.Database
	coll *mongo.Collection
}

func NewProductRepository(db *mongo.Database) domain.ProductRepository {
	return &productRepository{
		db:   db,
		coll: db.Collection("products"),
	}
}

func (r *productRepository) Create(ctx context.Context, p *domain.Product) error {
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()
	result, err := r.coll.InsertOne(ctx, p)
	if err != nil {
		return err
	}
	p.ID = result.InsertedID.(primitive.ObjectID)
	return nil
}

func (r *productRepository) FindByID(ctx context.Context, id string) (*domain.Product, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	var product domain.Product
	err = r.coll.FindOne(ctx, primitive.M{"_id": objectID}).Decode(&product)
	return &product, err
}

func (r *productRepository) FindAll(ctx context.Context) ([]domain.Product, error) {
	cursor, err := r.coll.Find(ctx, primitive.M{})
	if err != nil {
		return nil, err
	}
	var products []domain.Product
	err = cursor.All(ctx, &products)
	return products, err
}

func (r *productRepository) Update(ctx context.Context, p *domain.Product) error {
	p.UpdatedAt = time.Now()
	_, err := r.coll.UpdateOne(
		ctx,
		primitive.M{"_id": p.ID},
		primitive.M{"$set": p},
	)
	return err
}

func (r *productRepository) Delete(ctx context.Context, id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = r.coll.DeleteOne(ctx, primitive.M{"_id": objectID})
	return err
}
