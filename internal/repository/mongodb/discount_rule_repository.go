package mongodb

import (
	"context"
	"play-to-win-api/internal/domain"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type discountRuleRepository struct {
	db   *mongo.Database
	coll *mongo.Collection
}

func NewDiscountRuleRepository(db *mongo.Database) domain.DiscountRuleRepository {
	return &discountRuleRepository{
		db:   db,
		coll: db.Collection("discount_rules"),
	}
}

func (r *discountRuleRepository) Create(ctx context.Context, discountRule *domain.DiscountRule) error {
	discountRule.CreatedAt = time.Now()
	discountRule.UpdatedAt = time.Now()
	result, err := r.coll.InsertOne(ctx, discountRule)
	if err != nil {
		return err
	}
	discountRule.ID = result.InsertedID.(primitive.ObjectID)
	return nil
}

func (r *discountRuleRepository) FindByID(ctx context.Context, id string) (*domain.DiscountRule, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	var discountRule domain.DiscountRule
	err = r.coll.FindOne(ctx, primitive.M{"_id": objectID}).Decode(&discountRule)
	return &discountRule, err
}

func (r *discountRuleRepository) FindAll(ctx context.Context) ([]domain.DiscountRule, error) {
	pipeline := []bson.M{
		{
			"$lookup": bson.M{
				"from":         "campaigns",
				"localField":   "campaign_id",
				"foreignField": "_id",
				"as":           "campaign",
			},
		},
		{
			"$unwind": bson.M{
				"path":                       "$campaign",
				"preserveNullAndEmptyArrays": true,
			},
		},
		{
			"$project": bson.M{
				"_id":                           1,
				"campaign_id":                   1,
				"discount_type":                 1,
				"amount":                        1,
				"percentage":                    1,
				"item_category":                 1,
				"points_ratio":                  1,
				"max_discount_percentage":       1,
				"threshold_amount":              1,
				"discount_percentage_threshold": 1,
				"created_at":                    1,
				"updated_at":                    1,
				"campaign_name":                 "$campaign.name",
			},
		},
		{
			"$sort": bson.M{
				"created_at": -1,
			},
		},
	}

	cursor, err := r.coll.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var discountRules []domain.DiscountRule
	if err = cursor.All(ctx, &discountRules); err != nil {
		return nil, err
	}

	return discountRules, nil
}

func (r *discountRuleRepository) Update(ctx context.Context, discountRule *domain.DiscountRule) error {
	discountRule.UpdatedAt = time.Now()
	_, err := r.coll.UpdateOne(
		ctx,
		primitive.M{"_id": discountRule.ID},
		primitive.M{"$set": discountRule},
	)
	return err
}

func (r *discountRuleRepository) Delete(ctx context.Context, id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = r.coll.DeleteOne(ctx, primitive.M{"_id": objectID})
	return err
}
