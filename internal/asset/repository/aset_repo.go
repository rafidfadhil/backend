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

type AsetRepository interface {
	InsertAset(ctx context.Context, aset entity.Aset) (*entity.Aset, error)
	FindAllAset(ctx context.Context, limit int, page int) ([]entity.Aset, *utils.PaginationData, error)
	FindAset(ctx context.Context, asetID string) (*entity.Aset, error)
	UpdateAset(ctx context.Context, asetID string, Aset entity.Aset) (*entity.Aset, error)
	DeleteAset(ctx context.Context, asetID string) error
}

type asetRepository struct {
	coll *mongo.Collection
}

// FindAllAset implements AsetRepository.
func (a *asetRepository) FindAllAset(ctx context.Context, limit int, page int) ([]entity.Aset, *utils.PaginationData, error) {
	total, err := a.coll.CountDocuments(ctx, bson.M{})
	if err != nil {
		return nil, nil, err
	}

	paginate := utils.NewMongoPaginate(limit, page, int(total))
	opts := paginate.GetPaginatedOpts()
	paginationData := paginate.GetPaginationData()

	res, err := a.coll.Find(ctx, bson.M{}, opts)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil, err
		}
		return nil, nil, err
	}

	aset := make([]entity.Aset, 0)
	if err := res.All(ctx, &aset); err != nil {
		return nil, nil, err
	}

	return aset, paginationData, nil
}

// FindAset implements AsetRepository.
func (a *asetRepository) FindAset(ctx context.Context, asetID string) (*entity.Aset, error) {
	objectId, _ := primitive.ObjectIDFromHex(asetID)

	var aset entity.Aset
	err := a.coll.FindOne(ctx, bson.M{"_id": objectId}).Decode(&aset)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, err
		}
		return nil, err
	}

	return &aset, nil
}

// InsertAset implements AsetRepository.
func (a *asetRepository) InsertAset(ctx context.Context, aset entity.Aset) (*entity.Aset, error) {
	res, err := a.coll.InsertOne(ctx, aset)
	if err != nil {
		return nil, err
	}

	newID, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, err
	}

	newAset := &entity.Aset{
		AsetID:        newID,
		VendorID:      aset.AsetID,
		NamaAset:      aset.NamaAset,
		KategoriAset:  aset.KategoriAset,
		MerekAset:     aset.MerekAset,
		KodeProduksi:  aset.KodeProduksi,
		TahunProduksi: aset.TahunProduksi,
		DeskripsiAset: aset.DeskripsiAset,
		GambarAset:    aset.GambarAset,
		JumlahAset:    aset.JumlahAset,
		AsetMasuk:     aset.AsetMasuk,
	}

	return newAset, nil
}

// UpdateAset implements AsetRepository.
func (a *asetRepository) UpdateAset(ctx context.Context, asetID string, aset entity.Aset) (*entity.Aset, error) {
	objectId, _ := primitive.ObjectIDFromHex(asetID)

	filter := bson.M{"_id": objectId}
	update := bson.M{"$set": aset}
	_, err := a.coll.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}

	var updatedAset entity.Aset
	err = a.coll.FindOne(ctx, filter).Decode(&updatedAset)
	if err != nil {
		return nil, err
	}

	return &updatedAset, nil
}

// DeleteAset implements AsetRepository.
func (a *asetRepository) DeleteAset(ctx context.Context, asetID string) error {
	objectId, _ := primitive.ObjectIDFromHex(asetID)
	filter := bson.M{"_id": objectId}
	_, err := a.coll.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	return nil
}

func NewAsetRepository(db *mongo.Database) AsetRepository {
	return &asetRepository{
		coll: db.Collection("aset"),
	}
}
