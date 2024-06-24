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

type VendorRepository interface {
	InsertVendor(ctx context.Context, vendor entity.Vendor) (*entity.Vendor, error)
	FindAllVendors(ctx context.Context, limit int, page int) ([]entity.Vendor, *utils.PaginationData, error)
	FindVendor(ctx context.Context, vendorID string) (*entity.Vendor, error)
	UpdateVendor(ctx context.Context, vendorID string, vendor entity.Vendor) (*entity.Vendor, error)
	DeleteVendor(ctx context.Context, vendorID string) error
}

type vendorRepository struct {
	coll *mongo.Collection
}

func (v *vendorRepository) InsertVendor(ctx context.Context, vendor entity.Vendor) (*entity.Vendor, error) {
	res, err := v.coll.InsertOne(ctx, vendor)
	if err != nil {
		return nil, err
	}

	newID, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, err
	}

	newVendor := &entity.Vendor{
		VendorID:     newID,
		NamaVendor:   vendor.NamaVendor,
		TelpVendor:   vendor.TelpVendor,
		AlamatVendor: vendor.AlamatVendor,
		JenisVendor:  vendor.JenisVendor,
	}

	return newVendor, nil
}

func (v *vendorRepository) FindAllVendors(ctx context.Context, limit int, page int) ([]entity.Vendor, *utils.PaginationData, error) {
	total, err := v.coll.CountDocuments(ctx, bson.M{})
	if err != nil {
		return nil, nil, err
	}

	paginate := utils.NewMongoPaginate(limit, page, int(total))
	opts := paginate.GetPaginatedOpts()
	paginationData := paginate.GetPaginationData()

	res, err := v.coll.Find(ctx, bson.M{}, opts)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil, err
		}
		return nil, nil, err
	}

	vendors := make([]entity.Vendor, 0)
	if err := res.All(ctx, &vendors); err != nil {
		return nil, nil, err
	}

	return vendors, paginationData, nil
}

func (v *vendorRepository) FindVendor(ctx context.Context, vendorID string) (*entity.Vendor, error) {
	objectId, _ := primitive.ObjectIDFromHex(vendorID)

	var vendor entity.Vendor
	err := v.coll.FindOne(ctx, bson.M{"_id": objectId}).Decode(&vendor)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, err
		}
		return nil, err
	}

	return &vendor, nil
}

func (v *vendorRepository) UpdateVendor(ctx context.Context, vendorID string, vendor entity.Vendor) (*entity.Vendor, error) {
	objectId, _ := primitive.ObjectIDFromHex(vendorID)

	filter := bson.M{"_id": objectId}
	update := bson.M{"$set": vendor}
	_, err := v.coll.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}

	var updatedVendor entity.Vendor
	err = v.coll.FindOne(ctx, filter).Decode(&updatedVendor)
	if err != nil {
		return nil, err
	}

	return &updatedVendor, nil
}

func (v *vendorRepository) DeleteVendor(ctx context.Context, vendorID string) error {
	objectId, _ := primitive.ObjectIDFromHex(vendorID)

	filter := bson.M{"_id": objectId}
	_, err := v.coll.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	return nil
}

func NewVendorRepository(db *mongo.Database) VendorRepository {
	return &vendorRepository{
		coll: db.Collection("vendor"),
	}
}
