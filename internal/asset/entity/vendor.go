package entity

import (
	"errors"

	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Vendor struct {
	VendorID     primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	NamaVendor   string             `bson:"nama_vendor,omitempty" json:"nama_vendor,omitempty"`
	TelpVendor   string             `bson:"telp_vendor,omitempty" json:"telp_vendor,omitempty"`
	AlamatVendor string             `bson:"alamat_vendor,omitempty" json:"alamat_vendor,omitempty"`
	JenisVendor  string             `bson:"jenis_vendor,omitempty" json:"jenis_vendor,omitempty"`
}

type CreateVendor struct {
	NamaVendor   string `json:"nama_vendor" validate:"required"`
	TelpVendor   string `json:"telp_vendor" validate:"required"`
	AlamatVendor string `json:"alamat_vendor" validate:"required"`
	JenisVendor  string `json:"jenis_vendor" validate:"required"`
}

func (v *CreateVendor) Validate() error {
	validate := validator.New()

	err := validate.Struct(v)
	if err != nil {
		errorMessages := map[string]string{
			"Key: 'CreateVendor.Nama' Error:Field validation for 'Nama' failed on the 'required' tag":                 "Nama is required",
			"Key: 'CreateVendor.TelpVendor' Error:Field validation for 'TelpVendor' failed on the 'required' tag":     "TelpVendor is required",
			"Key: 'CreateVendor.AlamatVendor' Error:Field validation for 'AlamatVendor' failed on the 'required' tag": "AlamatVendor is required",
			"Key: 'CreateVendor.JenisVendor' Error:Field validation for 'JenisVendor' failed on the 'required' tag":   "JenisVendor is required",
		}

		if message, exists := errorMessages[err.Error()]; exists {
			return errors.New(message)
		} else {
			return err
		}
	}

	return nil
}

type UpdateVendor struct {
	NamaVendor   string `json:"nama_vendor,omitempty"`
	TelpVendor   string `json:"telp_vendor,omitempty"`
	AlamatVendor string `json:"alamat_vendor,omitempty"`
	JenisVendor  string `json:"jenis_vendor,omitempty"`
}
