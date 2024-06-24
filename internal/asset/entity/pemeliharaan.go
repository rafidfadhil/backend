package entity

import (
	"errors"

	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Pemeliharaan struct {
	PemeliharaanID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	RencanaID            primitive.ObjectID `json:"rencana_id,omitempty" bson:"rencana_id,omitempty"`
	KondisiStlhPerbaikan string             `json:"kondisi_stlh_perbaikan,omitempty" bson:"kondisi_stlh_perbaikan,omitempty"`
	StatusPemeliharaan   string             `json:"status_pemeliharaan,omitempty" bson:"status_pemeliharaan,omitempty"`
	PenanggungJawab      string             `json:"penanggung_jawab,omitempty" bson:"penanggung_jawab,omitempty"`
	Deskripsi            string             `json:"deskripsi,omitempty" bson:"deskripsi,omitempty"`
}

// type PemeliharaanDarurat struct {
// 	ID                   primitive.ObjectID `json:"id" bson:"_id,omitempty"`
// 	UsiaAset             string             `json:"usia_aset" bson:"usia_aset,omitempty"`
// 	MaksUsiaAset         string             `json:"maks_usia_aset" bson:"maks_usia_aset,omitempty"`
// 	ThnProduksi          string             `json:"thn_produksi" bson:"thn_produksi,omitempty"`
// 	TglPemeliharaan      string             `json:"tgl_pemeliharaan" bson:"tgl_pemeliharaan,omitempty"`
// 	KondisiStlhPerbaikan string             `json:"kondisi_stlh_perbaikan" bson:"kondisi_stlh_perbaikan,omitempty"`
// 	StatusPemeliharaan   string             `json:"status_pemeliharaan" bson:"status_pemeliharaan,omitempty"`
// 	PenanggungJawab      string             `json:"penanggung_jawab" bson:"penanggung_jawab,omitempty"`
// 	DeskripsiRusak       string             `json:"deskripsi_rusak" bson:"deskripsi_rusak,omitempty"`
// }

type CreatePelihara struct {
	RencanaID            string `json:"rencana_id" validate:"required"`
	KondisiStlhPerbaikan string `json:"kondisi_stlh_perbaikan" validate:"required"`
	StatusPemeliharaan   string `json:"status_pemeliharaan" validate:"required"`
	PenanggungJawab      string `json:"penanggung_jawab" validate:"required"`
	Deskripsi            string `json:"deskripsi" validate:"required"`
}

func (c *CreatePelihara) Validate() error {
	validate := validator.New()

	err := validate.Struct(c)
	if err != nil {
		errorMessages := map[string]string{
			"Key: 'Pemeliharaan.KondisiStlhPerbaikan' Error:Field validation for 'KondisiStlhPerbaikan' failed on the 'required' tag": "KondisiStlhPerbaikan is required",
			"Key: 'Pemeliharaan.StatusPemeliharaan' Error:Field validation for 'StatusPemeliharaan' failed on the 'required' tag":     "StatusPemeliharaan is required",
			"Key: 'Pemeliharaan.PenanggungJawab' Error:Field validation for 'PenanggungJawab' failed on the 'required' tag":           "PenanggungJawab is required",
			"Key: 'Pemeliharaan.Deskripsi' Error:Field validation for 'Deskripsi' failed on the 'required' tag":                       "Deskripsi is required",
		}

		if message, exists := errorMessages[err.Error()]; exists {
			return errors.New(message)
		} else {
			return err
		}
	}

	return nil
}

type UpdatePelihara struct {
	KondisiStlhPerbaikan string `json:"kondisi_stlh_perbaikan,omitempty"`
	StatusPemeliharaan   string `json:"status_pemeliharaan,omitempty"`
	PenanggungJawab      string `json:"penanggung_jawab,omitempty"`
	Deskripsi            string `json:"deskripsi,omitempty"`
}
