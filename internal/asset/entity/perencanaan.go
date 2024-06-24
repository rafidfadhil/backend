package entity

import (
	"errors"

	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Perencanaan struct {
	RencanaID      primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	AsetID         primitive.ObjectID `bson:"aset_id,omitempty" json:"aset_id,omitempty"`
	VendorID       primitive.ObjectID `bson:"vendor_id,omitempty" json:"vendor_id,omitempty"`
	KondisiAset    string             `bson:"kondisi_aset,omitempty" json:"kondisi_aset,omitempty"`
	TglPerencanaan string             `bson:"tgl_perencanaan,omitempty" json:"tgl_perencanaan,omitempty"`
	StatusAset     string             `bson:"status_aset,omitempty" json:"status_aset,omitempty"`
	UsiaAset       string             `bson:"usia_aset,omitempty" json:"usia_aset,omitempty"`
	MaksUsiaAset   string             `bson:"maks_usia_aset,omitempty" json:"maks_usia_aset,omitempty"`
	Deskripsi      string             `bson:"deskripsi,omitempty" json:"deskripsi,omitempty"`
}

type CreatePerencanaan struct {
	AsetID         string `bson:"aset_id" validate:"required"`
	VendorID       string `bson:"vendor_id" validate:"required"`
	KondisiAset    string `bson:"kondisi_aset" validate:"required"`
	TglPerencanaan string `bson:"tgl_perencanaan" validate:"required"`
	StatusAset     string `bson:"status_aset" validate:"required"`
	UsiaAset       string `bson:"usia_aset" validate:"required"`
	MaksUsiaAset   string `bson:"maks_usia_aset" validate:"required"`
	Deskripsi      string `bson:"deskripsi" validate:"required"`
}

func (p *CreatePerencanaan) Validate() error {
	validate := validator.New()

	err := validate.Struct(p)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			switch err.Field() {
			case "AsetID":
				return errors.New("AsetID is required")
			case "VendorID":
				return errors.New("VendorID is required")
			case "KondisiAset":
				return errors.New("KondisiAset is required")
			case "TglPerencanaan":
				return errors.New("TglPerencanaan is required")
			case "StatusAset":
				return errors.New("StatusAset is required")
			case "UsiaAset":
				return errors.New("UsiaAset is required")
			case "MaksUsiaAset":
				return errors.New("MaksUsiaAset is required")
			case "Deskripsi":
				return errors.New("Deskripsi is required")
			default:
				return err
			}
		}
	}

	return nil
}

type UpdatePerencanaan struct {
	KondisiAset    string `json:"kondisi_aset,omitempty"`
	TglPerencanaan string `json:"tgl_perencanaan,omitempty"`
	StatusAset     string `json:"status_aset,omitempty"`
	UsiaAset       string `json:"usia_aset,omitempty"`
	MaksUsiaAset   string `json:"maks_usia_aset,omitempty"`
	Deskripsi      string `json:"deskripsi,omitempty"`
}
