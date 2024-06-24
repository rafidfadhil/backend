package repository

import (
	"context"
	"errors"

	"github.com/BIC-Final-Project/backend/internal/auth/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuthRepository interface {
	SaveUser(ctx context.Context, user entity.User) (*entity.User, error)
	FindUserByEmail(ctx context.Context, email string) (*entity.User, error)
}
type authRepository struct {
	coll *mongo.Collection
}

// SaveUser implements AuthRepository.
func (a *authRepository) SaveUser(ctx context.Context, user entity.User) (*entity.User, error) {	
	res, err := a.coll.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}
	
	newID, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, errors.New("failed to get inserted ID")
	}
	
	newUser := &entity.User{
		ID:          newID,
		NamaLengkap: user.NamaLengkap,
		NoHandphone: user.NoHandphone,
		Email:       user.Email,
		Password:    user.Password,
		Role:        user.Role,
	}
	
	return newUser, nil
}

// FindAdminByUsername implements AuthRepository.
func (a *authRepository) FindUserByEmail(ctx context.Context, email string) (*entity.User, error) {	
	var user entity.User
	err := a.coll.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, err
		}
	}
	
	return &user, nil
}

func NewAuthRepository(db *mongo.Database) AuthRepository {
	return &authRepository{
		coll: db.Collection("users"),
	}
}