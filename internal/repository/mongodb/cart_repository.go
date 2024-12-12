package mongodb

import (
	"context"
	"play-to-win-api/internal/domain"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type cartRepository struct {
	db   *mongo.Database
	coll *mongo.Collection
}

func NewCartRepository(db *mongo.Database) domain.CartRepository {
	return &cartRepository{
		db:   db,
		coll: db.Collection("carts"),
	}
}

func (r *cartRepository) Create(cart *domain.Cart) error {
	cart.CreatedAt = time.Now()
	cart.UpdatedAt = time.Now()
	result, err := r.coll.InsertOne(context.Background(), cart)
	if err != nil {
		return err
	}
	cart.ID = result.InsertedID.(primitive.ObjectID)
	return nil
}

func (r *cartRepository) FindByUserID(ctx context.Context, userID string) ([]domain.Cart, error) {
	objectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, domain.ErrInvalidCartID
	}

	// Pipeline for aggregating cart with user data
	pipeline := []bson.M{
		{
			"$match": bson.M{
				"user._id": objectID,
			},
		},
		{
			"$lookup": bson.M{
				"from":         "users",
				"localField":   "user._id",
				"foreignField": "_id",
				"as":           "user_data",
			},
		},
		{
			"$unwind": "$user_data",
		},
		{
			"$addFields": bson.M{
				"user": "$user_data",
			},
		},
		{
			"$project": bson.M{
				"user_data": 0,
			},
		},
	}

	cursor, err := r.coll.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var carts []domain.Cart
	if err := cursor.All(ctx, &carts); err != nil {
		return nil, err
	}

	return carts, nil
}

func (r *cartRepository) FindByID(ctx context.Context, id string) (*domain.Cart, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	var cart domain.Cart
	err = r.coll.FindOne(ctx, primitive.M{"_id": objectID}).Decode(&cart)
	return &cart, err
}

func (r *cartRepository) FindAll(ctx context.Context) ([]domain.Cart, error) {
	cursor, err := r.coll.Find(ctx, primitive.M{})
	if err != nil {
		return nil, err
	}
	var carts []domain.Cart
	err = cursor.All(ctx, &carts)
	return carts, err
}

func (r *cartRepository) Update(ctx context.Context, cart *domain.Cart) error {
	cart.UpdatedAt = time.Now()
	_, err := r.coll.UpdateOne(
		ctx,
		primitive.M{"_id": cart.ID},
		primitive.M{"$set": cart},
	)
	return err
}

func (r *cartRepository) Delete(ctx context.Context, id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = r.coll.DeleteOne(ctx, primitive.M{"_id": objectID})
	return err
}
