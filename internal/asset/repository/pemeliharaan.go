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

type PemeliharaanRepository interface {
	InsertPelihara(ctx context.Context, pelihara entity.Pemeliharaan) (*entity.Pemeliharaan, error)
	FindAllPelihara(ctx context.Context, limit int, page int) ([]entity.Pemeliharaan, *utils.PaginationData, error)
	FindPelihara(ctx context.Context, id string) (*entity.Pemeliharaan, error)
	UpdatePelihara(ctx context.Context, id string, pelihara entity.Pemeliharaan) (*entity.Pemeliharaan, error)
	DeletePelihara(ctx context.Context, id string) error
}

type pemeliharaanRepository struct {
	coll *mongo.Collection
}

func (p *pemeliharaanRepository) InsertPelihara(ctx context.Context, pelihara entity.Pemeliharaan) (*entity.Pemeliharaan, error) {
	res, err := p.coll.InsertOne(ctx, pelihara)
	if err != nil {
		return nil, err
	}

	newID, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, err
	}

	newPelihara := &entity.Pemeliharaan{
		RencanaID:            pelihara.RencanaID,
		PemeliharaanID:       newID,
		KondisiStlhPerbaikan: pelihara.KondisiStlhPerbaikan,
		StatusPemeliharaan:   pelihara.StatusPemeliharaan,
		PenanggungJawab:      pelihara.PenanggungJawab,
		Deskripsi:            pelihara.Deskripsi,
	}

	return newPelihara, nil
}

func (p *pemeliharaanRepository) FindAllPelihara(ctx context.Context, limit int, page int) ([]entity.Pemeliharaan, *utils.PaginationData, error) {
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

	pelihara := make([]entity.Pemeliharaan, 0)
	if err := res.All(ctx, &pelihara); err != nil {
		return nil, nil, err
	}

	return pelihara, paginationData, nil
}

func (p *pemeliharaanRepository) FindPelihara(ctx context.Context, id string) (*entity.Pemeliharaan, error) {
	objectID, _ := primitive.ObjectIDFromHex(id)

	var pemeliharaan entity.Pemeliharaan
	err := p.coll.FindOne(ctx, bson.M{"_id": objectID}).Decode(&pemeliharaan)
	if err != nil {
		return nil, err
	}
	return &pemeliharaan, nil
}

func (p *pemeliharaanRepository) UpdatePelihara(ctx context.Context, id string, pelihara entity.Pemeliharaan) (*entity.Pemeliharaan, error) {
	objectID, _ := primitive.ObjectIDFromHex(id)

	filter := bson.M{"_id": objectID}
	update := bson.M{"$set": pelihara}
	_, err := p.coll.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}

	var updatedPelihara entity.Pemeliharaan
	err = p.coll.FindOne(ctx, filter).Decode(&updatedPelihara)
	if err != nil {
		return nil, err
	}
	return &updatedPelihara, nil
}

func (p *pemeliharaanRepository) DeletePelihara(ctx context.Context, id string) error {
	objectID, _ := primitive.ObjectIDFromHex(id)
	_, err := p.coll.DeleteOne(ctx, bson.M{"_id": objectID})
	if err != nil {
		return err
	}
	return nil
}

func NewPemeliharaanRepository(db *mongo.Database) PemeliharaanRepository {
	return &pemeliharaanRepository{
		coll: db.Collection("pemeliharaan"),
	}
}
