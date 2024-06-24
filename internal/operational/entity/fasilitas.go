package entity

import (
	"errors"
	"mime/multipart"

	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Fasilitas struct {
	ID                 primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	NamaFasilitas      string             `bson:"nama_fasilitas,omitempty" json:"nama_fasilitas,omitempty"`
	GambarFasilitas    GambarFasilitas    `bson:"gambar_fasilitas,omitempty" json:"gambar_fasilitas,omitempty"`
	DeskripsiFasilitas string             `bson:"deskripsi_fasilitas,omitempty" json:"deskripsi_fasilitas,omitempty"`
	HargaFasilitas     []HargaFasilitas   `bson:"harga_fasilitas,omitempty" json:"harga_fasilitas,omitempty"`
	CreatedAt          int64              `bson:"created_at,omitempty" json:"created_at,omitempty"`
	UpdatedAt          int64              `bson:"updated_at,omitempty" json:"updated_at,omitempty"`
}
type GambarFasilitas struct {
	ImageKey string `bson:"image_key,omitempty" json:"image_key,omitempty"`
	ImageURL string `bson:"image_url,omitempty" json:"image_url,omitempty"`
}

type HargaFasilitas struct {
	Hari []string `bson:"hari,omitempty" json:"hari,omitempty"`
	Jam  []Jam    `bson:"jam,omitempty" json:"jam,omitempty"`
}

type Jam struct {
	JamAwal  string `bson:"jam_awal,omitempty" json:"jam_awal,omitempty"`
	JamAkhir string `bson:"jam_akhir,omitempty" json:"jam_akhir,omitempty"`
	Harga    string `bson:"harga,omitempty" json:"harga,omitempty"`
}

type CreateFasilitas struct {
	Nama      string                `json:"nama" validate:"required"`
	Gambar    *multipart.FileHeader `json:"gambar"`
	Deskripsi string                `json:"deskripsi" validate:"required"`
	Harga     string                `json:"harga" validate:"required"`
}

func (c *CreateFasilitas) Validate() error {
	validate := validator.New()

	err := validate.Struct(c)
	if err != nil {
		errorMessages := map[string]string{
			"Key: 'CreateFasilitas.Nama' Error:Field validation for 'Nama' failed on the 'required' tag":           "Nama is required",
			"Key: 'CreateFasilitas.Deskripsi' Error:Field validation for 'Deskripsi' failed on the 'required' tag": "Deskripsi is required",
			"Key: 'CreateFasilitas.Harga' Error:Field validation for 'Harga' failed on the 'required' tag":         "Harga is required",
		}

		if message, exists := errorMessages[err.Error()]; exists {
			return errors.New(message)
		} else {
			return err
		}
	}

	return nil
}

type UpdateFasilitas struct {
	Nama      string                `json:"nama,omitempty"`
	Gambar    *multipart.FileHeader `json:"gambar",omitempty`
	Deskripsi string                `json:"deskripsi",omitempty`
	Harga     string                `json:"harga",omitempty`
}
