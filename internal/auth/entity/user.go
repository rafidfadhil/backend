package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	NamaLengkap string             `bson:"nama_lengkap,omitempty" json:"nama_lengkap,omitempty"`
	NoHandphone string             `bson:"no_handphone,omitempty" json:"no_handphone,omitempty"`
	Email       string             `bson:"email,omitempty" json:"email,omitempty"`
	Password    string             `bson:"password,omitempty" json:"password,omitempty"`
	Role        string             `bson:"role,omitempty" json:"role,omitempty"`
}

type UserResponse struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	NamaLengkap string             `bson:"nama_lengkap,omitempty" json:"nama_lengkap,omitempty"`
	NoHandphone string             `bson:"no_handphone,omitempty" json:"no_handphone,omitempty"`
	Email       string             `bson:"email,omitempty" json:"email,omitempty"`
	Role        string             `bson:"role,omitempty" json:"role,omitempty"`
}