package entity

import (
	"errors"
	"mime/multipart"

	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Aset struct {
	AsetID        primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	VendorID      primitive.ObjectID `bson:"vendor_id,omitempty" json:"vendor_id,omitempty"`
	NamaAset      string             `bson:"nama_aset,omitempty" json:"nama_aset,omitempty"`
	KategoriAset  string             `bson:"kategori_aset,omitempty" json:"kategori_aset,omitempty"`
	MerekAset     string             `bson:"merek_aset,omitempty" json:"merek_aset,omitempty"`
	KodeProduksi  string             `bson:"kode_produksi,omitempty" json:"kode_produksi,omitempty"`
	TahunProduksi string             `bson:"tahun_produksi,omitemtpy" json:"tahun_produksi,omitempty"`
	DeskripsiAset string             `bson:"deskripsi_aset,omitempty" json:"deskripsi_aset,omitempty"`
	GambarAset    Gambar             `bson:"gambar_aset,omitemtpy" json:"gambar_aset,omitemtpy"`
	JumlahAset    string             `bson:"jumlah_aset,omitempty" json:"jumlah_aset,omitempty"`
	AsetMasuk     string             `bson:"aset_masuk,omitempty" json:"aset_masuk,omitempty"`
}

type Gambar struct {
	ImageKey string `bson:"image_key,omitempty" json:"image_key,omitempty"`
	ImageURL string `bson:"image_url,omitempty" json:"image_url,omitempty"`
}

type CreateAset struct {
	VendorID      string                `json:"vendor_id" validate:"required"`
	NamaAset      string                `json:"nama" validate:"required"`
	Kategori      string                `json:"kategori" validate:"required"`
	MerekAset     string                `json:"merek" validate:"required"`
	Kode          string                `json:"kode" validate:"required"`
	TahunProduksi string                `json:"produksi" validate:"required"`
	Deskripsi     string                `json:"deskripsi" validate:"required"`
	Gambar        *multipart.FileHeader `json:"gambar"`
	Jumlah        string                `json:"jumlah" validate:"required"`
	AsetMasuk     string                `json:"aset_masuk" validate:"required"`
}

func (c *CreateAset) Validate() error {
	validate := validator.New()

	err := validate.Struct(c)
	if err != nil {
		errorMessages := map[string]string{
			"Key: 'CreateAset.Nama' Error:Field validation for 'Nama' failed on the 'required' tag":                   "Nama Aset is required",
			"Key: 'CreateAset.Kategori' Error:Field validation for 'Kategori' failed on the 'required' tag":           "Kategori Aset is required",
			"Key: 'CreateAset.MerekAset' Error:Field validation for 'MerekAset' failed on the 'required' tag":         "Merek Aset is required",
			"Key: 'CreateAset.Kode' Error:Field validation for 'Kode' failed on the 'required' tag":                   "Kode Aset is required",
			"Key: 'CreateAset.TahunProduksi' Error:Field validation for 'TahunProduksi' failed on the 'required' tag": "Tahun Produksi Aset is required",
			"Key: 'CreateAset.Deskripsi' Error:Field validation for 'Deskripsi' failed on the 'required' tag":         "Deskripsi Aset is required",
			"Key: 'CreateAset.Jumlah' Error:Field validation for 'Jumlah' failed on the 'required' tag":               "Jumlah Aset is required",
			"Key: 'CreateAset.AsetMasuk' Error:Field validation for 'AsetMasuk' failed on the 'required' tag":         "Aset Masuk is required",
		}

		if message, exists := errorMessages[err.Error()]; exists {
			return errors.New(message)
		} else {
			return err
		}
	}
	return nil
}

type UpdateAset struct {
	NamaAset      string                `json:"nama,omitempty"`
	Kategori      string                `json:"kategori,omitempty"`
	MerekAset     string                `json:"merek,omitempty"`
	Kode          string                `json:"kode,omitempty"`
	TahunProduksi string                `json:"produksi,omitempty"`
	Deskripsi     string                `json:"deskripsi,omitempty"`
	Gambar        *multipart.FileHeader `json:"gambar,omitempty"`
	Jumlah        string                `json:"jumlah,omitempty"`
	AsetMasuk     string                `json:"aset_masuk,omitempty"`
}
