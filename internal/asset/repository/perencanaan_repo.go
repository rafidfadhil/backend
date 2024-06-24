package repository

import (
	"context"
	"errors"

	"github.com/BIC-Final-Project/backend/internal/asset/entity"
	"github.com/BIC-Final-Project/backend/pkg/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type PerencanaanRepository interface {
	InsertRencana(ctx context.Context, rencana entity.Perencanaan) (*entity.Perencanaan, error)
	FindAllRencana(ctx context.Context, limit int, page int) ([]entity.Perencanaan, *utils.PaginationData, error)
	FindRencana(ctx context.Context, id string) (*entity.Perencanaan, error)
	UpdateRencana(ctx context.Context, id string, rencana entity.Perencanaan) (*entity.Perencanaan, error)
	DeleteRencana(ctx context.Context, id string) error
}

type perencanaanRepository struct {
	coll *mongo.Collection
}

func (p *perencanaanRepository) InsertRencana(ctx context.Context, rencana entity.Perencanaan) (*entity.Perencanaan, error) {
	res, err := p.coll.InsertOne(ctx, rencana)
	if err != nil {
		return nil, err
	}

	newID, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, err
	}

	newRencana := &entity.Perencanaan{
		RencanaID:      newID,
		AsetID:         rencana.AsetID,
		VendorID:       rencana.VendorID,
		KondisiAset:    rencana.KondisiAset,
		StatusAset:     rencana.StatusAset,
		TglPerencanaan: rencana.TglPerencanaan,
		UsiaAset:       rencana.UsiaAset,
		MaksUsiaAset:   rencana.MaksUsiaAset,
		Deskripsi:      rencana.Deskripsi,
	}
	return newRencana, nil
}

func (p *perencanaanRepository) FindAllRencana(ctx context.Context, limit int, page int) ([]entity.Perencanaan, *utils.PaginationData, error) {
	total, err := p.coll.CountDocuments(ctx, bson.M{})
	if err != nil {
		return nil, nil, err
	}

	paginate := utils.NewMongoPaginate(limit, page, int(total))
	opts := paginate.GetPaginatedOpts()
	paginationData := paginate.GetPaginationData()

	res, err := p.coll.Find(ctx, bson.M{}, opts)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil, err
		}
		return nil, nil, err
	}

	perencanaan := make([]entity.Perencanaan, 0)
	if err := res.All(ctx, &perencanaan); err != nil {
		return nil, nil, err
	}
	return perencanaan, paginationData, nil
}

func (p *perencanaanRepository) FindRencana(ctx context.Context, id string) (*entity.Perencanaan, error) {
	objectId, _ := primitive.ObjectIDFromHex(id)

	var perencanaan entity.Perencanaan
	err := p.coll.FindOne(ctx, bson.M{"_id": objectId}).Decode(&perencanaan)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, err
		}
		return nil, err
	}

	return &perencanaan, nil
}

func (p *perencanaanRepository) UpdateRencana(ctx context.Context, id string, rencana entity.Perencanaan) (*entity.Perencanaan, error) {
	objectId, _ := primitive.ObjectIDFromHex(id)

	filter := bson.M{"_id": objectId}
	update := bson.M{"$set": rencana}
	_, err := p.coll.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}

	var updatedRencana entity.Perencanaan
	err = p.coll.FindOne(ctx, filter).Decode(&updatedRencana)
	if err != nil {
		return nil, err
	}
	return &updatedRencana, nil
}

func (p *perencanaanRepository) DeleteRencana(ctx context.Context, id string) error {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = p.coll.DeleteOne(ctx, bson.M{"_id": objectId})
	if err != nil {
		return err
	}
	return nil
}

func NewPerencanaanRepository(db *mongo.Database) PerencanaanRepository {
	return &perencanaanRepository{
		coll: db.Collection("perencanaan"),
	}
}
