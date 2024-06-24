package entity

import (
	"errors"

	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MembershipType struct {
	ID                       primitive.ObjectID    `bson:"_id,omitempty" json:"_id,omitempty"`
	JenisPaket               string                `bson:"jenis_paket,omitempty" json:"jenis_paket,omitempty"`
	JenisKeanggotaan         string                `bson:"jenis_keanggotaan,omitempty" json:"jenis_keanggotaan,omitempty"`
	JumlahAnggotaYangBerlaku int64                 `bson:"jumlah_anggota_yang_berlaku,omitempty" json:"jumlah_anggota_yang_berlaku,omitempty"`
	Harga                    string                `bson:"harga,omitempty" json:"harga,omitempty"`
	FasilitasMembership      []FasilitasMembership `bson:"fasilitas_membership,omitempty" json:"fasilitas_membership,omitempty"`
	CreatedAt                int64                 `bson:"created_at,omitempty" json:"created_at,omitempty"`
	UpdatedAt                int64                 `bson:"updated_at,omitempty" json:"updated_at,omitempty"`
}

type FasilitasMembership struct {
	NamaFasilitas string `bson:"nama_fasilitas,omitempty" json:"nama_fasilitas,omitempty"`
	IsWithMembers bool   `bson:"is_with_members" json:"is_with_members"`
}

type CreateMembershipType struct {
	JenisPaket               string                `json:"jenis_paket" validate:"required"`
	JenisKeanggotaan         string                `json:"jenis_keanggotaan" validate:"required"`
	JumlahAnggotaYangBerlaku int64                 `json:"jumlah_anggota_yang_berlaku" validate:"required"`
	Harga                    string                `json:"harga" validate:"required"`
	FasilitasMembership      []FasilitasMembership `json:"fasilitas_membership" validate:"required"`
}

func (c *CreateMembershipType) Validate() error {
	validate := validator.New()

	err := validate.Struct(c)
	if err != nil {
		errorMessages := map[string]string{
			"Key: 'CreateMembershipType.JenisPaket' Error:Field validation for 'JenisPaket' failed on the 'required' tag":                             "Jenis Paket is required",
			"Key: 'CreateMembershipType.JenisKeanggotaan' Error:Field validation for 'JenisKeanggotaan' failed on the 'required' tag":                 "Jenis Keanggotaan is required",
			"Key: 'CreateMembershipType.JumlahAnggotaYangBerlaku' Error:Field validation for 'JumlahAnggotaYangBerlaku' failed on the 'required' tag": "Jumlah Anggota Yang Berlaku is required",
			"Key: 'CreateMembershipType.Harga' Error:Field validation for 'Harga' failed on the 'required' tag":                                       "Harga is required",
			"Key: 'CreateMembershipType.FasilitasMembership' Error:Field validation for 'FasilitasMembership' failed on the 'required' tag":           "Fasilitas Membership is required",
		}

		if message, exists := errorMessages[err.Error()]; exists {
			return errors.New(message)
		} else {
			return err
		}
	}

	return nil
}

type UpdateMembershipType struct {
	JenisPaket               string                `json:"jenis_paket"`
	JenisKeanggotaan         string                `json:"jenis_keanggotaan"`
	JumlahAnggotaYangBerlaku int64                 `json:"jumlah_anggota_yang_berlaku"`
	Harga                    string                `json:"harga"`
	FasilitasMembership      []FasilitasMembership `json:"fasilitas_membership"`
}
