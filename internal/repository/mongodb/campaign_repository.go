package mongodb

import (
	"context"
	"play-to-win-api/internal/domain"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type campaignRepository struct {
	db   *mongo.Database
	coll *mongo.Collection
}

func NewCampaignRepository(db *mongo.Database) domain.CampaignRepository {
	return &campaignRepository{
		db:   db,
		coll: db.Collection("campaigns"),
	}
}

func (r *campaignRepository) Create(ctx context.Context, c *domain.Campaign) error {
	c.CreatedAt = time.Now()
	c.UpdatedAt = time.Now()
	result, err := r.coll.InsertOne(ctx, c)
	if err != nil {
		return err
	}
	c.ID = result.InsertedID.(primitive.ObjectID)
	return nil
}

func (r *campaignRepository) FindByID(ctx context.Context, id string) (*domain.Campaign, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	var campaign domain.Campaign
	err = r.coll.FindOne(ctx, primitive.M{"_id": objectID}).Decode(&campaign)
	return &campaign, err
}

func (r *campaignRepository) FindAll(ctx context.Context) ([]domain.Campaign, error) {
	cursor, err := r.coll.Find(ctx, primitive.M{})
	if err != nil {
		return nil, err
	}
	var campaigns []domain.Campaign
	err = cursor.All(ctx, &campaigns)
	return campaigns, err
}

func (r *campaignRepository) Update(ctx context.Context, c *domain.Campaign) error {
	c.UpdatedAt = time.Now()
	_, err := r.coll.UpdateOne(
		ctx,
		primitive.M{"_id": c.ID},
		primitive.M{"$set": c},
	)
	return err
}

func (r *campaignRepository) Delete(ctx context.Context, id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = r.coll.DeleteOne(ctx, primitive.M{"_id": objectID})
	return err
}
