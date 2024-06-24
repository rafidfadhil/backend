package repository

import (
	"context"
	"errors"

	"github.com/BIC-Final-Project/backend/internal/operational/entity"
	"github.com/BIC-Final-Project/backend/pkg/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type FasilitasRepository interface {
	InsertFasilitas(ctx context.Context, fasilitas entity.Fasilitas) (*entity.Fasilitas, error)
	FindAllFasilitas(ctx context.Context, limit int, page int) ([]entity.Fasilitas, *utils.PaginationData, error)
	FindFasilitas(ctx context.Context, id string) (*entity.Fasilitas, error)
	UpdateFasilitas(ctx context.Context, id string, fasilitas entity.Fasilitas) (*entity.Fasilitas, error)
	DeleteFasilitas(ctx context.Context, id string) error
	FindAllFasilitasName(ctx context.Context) ([]string, error)
}

type fasilitasRepository struct {
	coll *mongo.Collection
}

// FindAllFasilitasName implements FasilitasRepository.
func (f *fasilitasRepository) FindAllFasilitasName(ctx context.Context) ([]string, error) {
    var fasilitasNames []string
    cursor, err := f.coll.Find(ctx, bson.M{})
    if err != nil {
        return nil, err
    }
    defer cursor.Close(ctx)

    for cursor.Next(ctx) {
        var fasilitas entity.Fasilitas
        err := cursor.Decode(&fasilitas)
        if err != nil {
            return nil, err
        }
        fasilitasNames = append(fasilitasNames, fasilitas.NamaFasilitas)
    }

    if err := cursor.Err(); err != nil {
        return nil, err
    }

    return fasilitasNames, nil
}

// FindAllFasilitas implements FasilitasRepository.
func (f *fasilitasRepository) FindAllFasilitas(ctx context.Context, limit int, page int) ([]entity.Fasilitas, *utils.PaginationData, error) {
	total, err := f.coll.CountDocuments(ctx, bson.M{})
	if err != nil {
		return nil, nil, err
	}

	paginate := utils.NewMongoPaginate(limit, page, int(total))
	opts := paginate.GetPaginatedOpts()
	paginationData := paginate.GetPaginationData()

	res, err := f.coll.Find(ctx, bson.M{}, opts)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil, err
		}
		return nil, nil, err
	}

	fasilitas := make([]entity.Fasilitas, 0)
	if err := res.All(ctx, &fasilitas); err != nil {
		return nil, nil, err
	}

	return fasilitas, paginationData, nil
}

// DeleteFasilitas implements FasilitasRepository.
func (f *fasilitasRepository) DeleteFasilitas(ctx context.Context, id string) error {
	objectId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": objectId}
	_, err := f.coll.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	return nil
}

// FindFasilitas implements FasilitasRepository.
func (f *fasilitasRepository) FindFasilitas(ctx context.Context, id string) (*entity.Fasilitas, error) {
	objectId, _ := primitive.ObjectIDFromHex(id)

	var fasilitas entity.Fasilitas
	err := f.coll.FindOne(ctx, bson.M{"_id": objectId}).Decode(&fasilitas)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, err
		}
		return nil, err
	}

	return &fasilitas, nil
}

// InsertFasilitas implements FasilitasRepository.
func (f *fasilitasRepository) InsertFasilitas(ctx context.Context, fasilitas entity.Fasilitas) (*entity.Fasilitas, error) {
	res, err := f.coll.InsertOne(ctx, fasilitas)
	if err != nil {
		return nil, err
	}

	newID, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, err
	}

	newFasilitas := &entity.Fasilitas{
		ID:                 newID,
		NamaFasilitas:      fasilitas.NamaFasilitas,
		GambarFasilitas:    fasilitas.GambarFasilitas,
		DeskripsiFasilitas: fasilitas.DeskripsiFasilitas,
		HargaFasilitas:     fasilitas.HargaFasilitas,
		CreatedAt:          fasilitas.CreatedAt,
		UpdatedAt:          fasilitas.UpdatedAt,
	}

	return newFasilitas, nil
}

// UpdateFasilitas implements FasilitasRepository.
func (f *fasilitasRepository) UpdateFasilitas(ctx context.Context, id string, fasilitas entity.Fasilitas) (*entity.Fasilitas, error) {
	objectId, _ := primitive.ObjectIDFromHex(id)

	filter := bson.M{"_id": objectId}
	update := bson.M{"$set": fasilitas}
	_, err := f.coll.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}

	var updatedFasilitas entity.Fasilitas
	err = f.coll.FindOne(ctx, filter).Decode(&updatedFasilitas)
	if err != nil {
		return nil, err
	}

	return &updatedFasilitas, nil
}

func NewFasilitasRepository(db *mongo.Database) FasilitasRepository {
	return &fasilitasRepository{
		coll: db.Collection("fasilitas"),
	}
}
