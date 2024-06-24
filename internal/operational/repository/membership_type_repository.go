package repository

import (
	"context"
	"errors"

	"github.com/BIC-Final-Project/backend/internal/operational/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MembershipTypeRepository interface {
	InsertMembershipType(ctx context.Context, membershipType entity.MembershipType) (*entity.MembershipType, error)
	FindAllMembershipType(ctx context.Context, jenisPaket string) ([]entity.MembershipType, error)
	FindMembershipType(ctx context.Context, id string) (*entity.MembershipType, error)
	UpdateMembershipType(ctx context.Context, id string, membershipType entity.MembershipType) (*entity.MembershipType, error)
	DeleteMembershipType(ctx context.Context, id string) error
}

type membershipTypeRepository struct {
	coll *mongo.Collection
}

// FindAllMembershipType implements MembershipTypeRepository.
func (m *membershipTypeRepository) FindAllMembershipType(ctx context.Context, jenisPaket string) ([]entity.MembershipType, error) {
	filter := bson.M{}
	filter["jenis_paket"] = jenisPaket

	res, err := m.coll.Find(ctx, filter)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, err
		}
		return nil, err
	}

	membershipType := make([]entity.MembershipType, 0)
	if err := res.All(ctx, &membershipType); err != nil {
		return nil, err
	}

	return membershipType, nil
}

// DeleteMembershipType implements MembershipTypeRepository.
func (m *membershipTypeRepository) DeleteMembershipType(ctx context.Context, id string) error {
	objectId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": objectId}
	_, err := m.coll.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	return nil
}

// FindMembershipType implements MembershipTypeRepository.
func (m *membershipTypeRepository) FindMembershipType(ctx context.Context, id string) (*entity.MembershipType, error) {
	objectId, _ := primitive.ObjectIDFromHex(id)

	var membershipType entity.MembershipType
	err := m.coll.FindOne(ctx, bson.M{"_id": objectId}).Decode(&membershipType)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, err
		}
		return nil, err
	}

	return &membershipType, nil
}

// InsertMembershipType implements MembershipTypeRepository.
func (m *membershipTypeRepository) InsertMembershipType(ctx context.Context, membershipType entity.MembershipType) (*entity.MembershipType, error) {
	res, err := m.coll.InsertOne(ctx, membershipType)
	if err != nil {
		return nil, err
	}

	newID, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, err
	}

	newMembershipType := &entity.MembershipType{
		ID:                       newID,
		JenisPaket:               membershipType.JenisPaket,
		JenisKeanggotaan:         membershipType.JenisKeanggotaan,
		JumlahAnggotaYangBerlaku: membershipType.JumlahAnggotaYangBerlaku,
		Harga:                    membershipType.Harga,
		FasilitasMembership:      membershipType.FasilitasMembership,
		CreatedAt:                membershipType.CreatedAt,
		UpdatedAt:                membershipType.UpdatedAt,
	}

	return newMembershipType, nil
}

// UpdateMembershipType implements MembershipTypeRepository.
func (m *membershipTypeRepository) UpdateMembershipType(ctx context.Context, id string, membershipType entity.MembershipType) (*entity.MembershipType, error) {
	objectId, _ := primitive.ObjectIDFromHex(id)

	filter := bson.M{"_id": objectId}
	update := bson.M{"$set": membershipType}
	_, err := m.coll.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}

	var updatedMembershipType entity.MembershipType
	err = m.coll.FindOne(ctx, filter).Decode(&updatedMembershipType)
	if err != nil {
		return nil, err
	}

	return &updatedMembershipType, nil
}

func NewMembershipTypeRepository(db *mongo.Database) MembershipTypeRepository {
	return &membershipTypeRepository{
		coll: db.Collection("membership_type"),
	}
}
