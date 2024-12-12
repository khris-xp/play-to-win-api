package mongodb

import (
	"context"
	"play-to-win-api/internal/domain"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type cartItemRepository struct {
	db   *mongo.Database
	coll *mongo.Collection
}

func NewCartItemRepository(db *mongo.Database) domain.CartItemRepository {
	return &cartItemRepository{
		db:   db,
		coll: db.Collection("cart_items"),
	}
}

func (r *cartItemRepository) Create(ctx context.Context, cartItem *domain.CartItem) error {
	cartItem.CreatedAt = time.Now()
	cartItem.UpdatedAt = time.Now()
	result, err := r.coll.InsertOne(ctx, cartItem)
	if err != nil {
		return err
	}
	cartItem.ID = result.InsertedID.(primitive.ObjectID)
	return nil
}

func (r *cartItemRepository) FindByCartID(ctx context.Context, cartID string) ([]domain.CartItem, error) {
	objectID, err := primitive.ObjectIDFromHex(cartID)
	if err != nil {
		return nil, domain.ErrInvalidCartItemID
	}

	pipeline := []bson.M{
		{
			"$match": bson.M{
				"cart_id": objectID,
			},
		},
		{
			"$lookup": bson.M{
				"from":         "products",
				"localField":   "product_id",
				"foreignField": "_id",
				"as":           "product",
			},
		},
		{
			"$unwind": bson.M{
				"path":                       "$product",
				"preserveNullAndEmptyArrays": true,
			},
		},
		{
			"$addFields": bson.M{
				"product_name":        "$product.name",
				"product_description": "$product.description",
				"product_image":       "$product.image",
				"product_price":       "$product.price",
			},
		},
		{
			"$project": bson.M{
				"product": 0, // Remove the nested product object
			},
		},
	}

	cursor, err := r.coll.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var cartItems []domain.CartItem
	if err = cursor.All(ctx, &cartItems); err != nil {
		return nil, err
	}

	return cartItems, nil
}

func (r *cartItemRepository) FindByID(ctx context.Context, id string) (*domain.CartItem, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, domain.ErrInvalidCartItemID
	}

	pipeline := []bson.M{
		{
			"$match": bson.M{
				"_id": objectID,
			},
		},
		{
			"$lookup": bson.M{
				"from":         "products",
				"localField":   "product_id",
				"foreignField": "_id",
				"as":           "product",
			},
		},
		{
			"$unwind": bson.M{
				"path":                       "$product",
				"preserveNullAndEmptyArrays": true,
			},
		},
		{
			"$addFields": bson.M{
				"product_name":        "$product.name",
				"product_description": "$product.description",
				"product_image":       "$product.image",
				"product_price":       "$product.price",
			},
		},
		{
			"$project": bson.M{
				"product": 0,
			},
		},
	}

	var cartItems []domain.CartItem
	cursor, err := r.coll.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	if err = cursor.All(ctx, &cartItems); err != nil {
		return nil, err
	}

	if len(cartItems) == 0 {
		return nil, domain.ErrCartItemNotFound
	}

	return &cartItems[0], nil
}

func (r *cartItemRepository) FindAll(ctx context.Context) ([]domain.CartItem, error) {
	pipeline := []bson.M{
		{
			"$lookup": bson.M{
				"from":         "products",
				"localField":   "product_id",
				"foreignField": "_id",
				"as":           "product",
			},
		},
		{
			"$unwind": bson.M{
				"path":                       "$product",
				"preserveNullAndEmptyArrays": true,
			},
		},
		{
			"$addFields": bson.M{
				"product_name":        "$product.name",
				"product_description": "$product.description",
				"product_image":       "$product.image",
				"product_price":       "$product.price",
			},
		},
		{
			"$project": bson.M{
				"product": 0,
			},
		},
	}

	cursor, err := r.coll.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var cartItems []domain.CartItem
	if err = cursor.All(ctx, &cartItems); err != nil {
		return nil, err
	}

	return cartItems, nil
}

func (r *cartItemRepository) Update(ctx context.Context, cartItem *domain.CartItem) error {
	cartItem.UpdatedAt = time.Now()
	_, err := r.coll.UpdateOne(
		ctx,
		bson.M{"_id": cartItem.ID},
		bson.M{"$set": cartItem},
	)
	return err
}

func (r *cartItemRepository) Delete(ctx context.Context, id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return domain.ErrInvalidCartItemID
	}

	_, err = r.coll.DeleteOne(ctx, bson.M{"_id": objectID})
	return err
}
