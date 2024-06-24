package usecase

import (
	"context"
	"fmt"

	"github.com/BIC-Final-Project/backend/internal/asset/entity"
	"github.com/BIC-Final-Project/backend/internal/asset/repository"
	"github.com/BIC-Final-Project/backend/internal/storage"
	"github.com/BIC-Final-Project/backend/pkg/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AsetUsecase interface {
	InsertAset(ctx context.Context, req entity.CreateAset) (*entity.Aset, error)
	FindAllAset(ctx context.Context, limit int, page int) ([]entity.Aset, *utils.PaginationData, error)
	FindAset(ctx context.Context, asetID string) (*entity.Aset, error)
	UpdateAset(ctx context.Context, asetID string, req entity.UpdateAset) (*entity.Aset, error)
	DeleteAset(ctx context.Context, asetID string) error
}

type asetUsecase struct {
	repo repository.AsetRepository
	s3   storage.S3Service
}

var AsetFolder = "aset-bic"

func (a *asetUsecase) InsertAset(ctx context.Context, req entity.CreateAset) (*entity.Aset, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	vendorObjectID, _ := primitive.ObjectIDFromHex(req.VendorID)

	aset := entity.Aset{
		VendorID:      vendorObjectID,
		NamaAset:      req.NamaAset,
		KategoriAset:  req.Kategori,
		MerekAset:     req.MerekAset,
		KodeProduksi:  req.Kode,
		TahunProduksi: req.TahunProduksi,
		DeskripsiAset: req.Deskripsi,
		JumlahAset:    req.Jumlah,
		AsetMasuk:     req.AsetMasuk,
	}

	if req.Gambar != nil {
		res, err := a.s3.Upload(req.Gambar, AsetFolder)
		if err != nil {
			return nil, err
		}

		aset.GambarAset = entity.Gambar{
			ImageKey: *res.Key,
			ImageURL: fmt.Sprintf("%s%s", storage.CloudFrontURLBase, *res.Key),
		}
	}

	data, err := a.repo.InsertAset(ctx, aset)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (a *asetUsecase) FindAllAset(ctx context.Context, limit int, page int) ([]entity.Aset, *utils.PaginationData, error) {
	return a.repo.FindAllAset(ctx, limit, page)
}

func (a *asetUsecase) FindAset(ctx context.Context, asetID string) (*entity.Aset, error) {
	aset, err := a.repo.FindAset(ctx, asetID)
	if err != nil {
		return nil, err
	}

	return aset, nil
}

func (a *asetUsecase) UpdateAset(ctx context.Context, asetID string, req entity.UpdateAset) (*entity.Aset, error) {
	aset, err := a.repo.FindAset(ctx, asetID)
	if err != nil {
		return nil, err
	}

	updatedAset := entity.Aset{
		VendorID:      aset.VendorID,
		NamaAset:      req.NamaAset,
		KategoriAset:  req.Kategori,
		GambarAset:    aset.GambarAset,
		MerekAset:     req.MerekAset,
		KodeProduksi:  req.Kode,
		TahunProduksi: req.TahunProduksi,
		DeskripsiAset: req.Deskripsi,
		JumlahAset:    req.Jumlah,
	}

	if req.Gambar != nil {
		res, err := a.s3.Update(aset.GambarAset.ImageKey, req.Gambar, AsetFolder)
		if err != nil {
			return nil, err
		}

		updatedAset.GambarAset = entity.Gambar{
			ImageKey: *res.Key,
			ImageURL: fmt.Sprintf("%s%s", storage.CloudFrontURLBase, *res.Key),
		}
	}

	data, err := a.repo.UpdateAset(ctx, asetID, updatedAset)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (a *asetUsecase) DeleteAset(ctx context.Context, asetID string) error {
	aset, err := a.repo.FindAset(ctx, asetID)
	if err != nil {
		return err
	}

	if aset.GambarAset.ImageKey != "" {
		if err := a.s3.Delete(aset.GambarAset.ImageKey); err != nil {
			return err
		}
	}

	if err := a.repo.DeleteAset(ctx, asetID); err != nil {
		return err
	}

	return nil
}

func NewAsetUsecase(repo repository.AsetRepository, s3 storage.S3Service) AsetUsecase {
	return &asetUsecase{repo, s3}
}
