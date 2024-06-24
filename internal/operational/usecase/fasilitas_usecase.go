package usecase

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/BIC-Final-Project/backend/internal/operational/entity"
	"github.com/BIC-Final-Project/backend/internal/operational/repository"
	"github.com/BIC-Final-Project/backend/internal/storage"
	"github.com/BIC-Final-Project/backend/pkg/utils"
)

type FasilitasUsecase interface {
	InsertFasilitas(ctx context.Context, req entity.CreateFasilitas) (*entity.Fasilitas, error)
	FindAllFasilitas(ctx context.Context, limit int, page int) ([]entity.Fasilitas, *utils.PaginationData, error)
	FindFasilitas(ctx context.Context, id string) (*entity.Fasilitas, error)
	UpdateFasilitas(ctx context.Context, id string, req entity.UpdateFasilitas) (*entity.Fasilitas, error)
	DeleteFasilitas(ctx context.Context, id string) error
	FindAllFasilitasName(ctx context.Context) ([]string, error)
}

type fasilitasUsecase struct {
	repo repository.FasilitasRepository
	s3   storage.S3Service
}

// FindAllFasilitasName implements FasilitasUsecase.
func (f *fasilitasUsecase) FindAllFasilitasName(ctx context.Context) ([]string, error) {
	return f.repo.FindAllFasilitasName(ctx)
}

var FasilitasFolder = "bic-fasilitas"

// DeleteFasilitas implements FasilitasUsecase.
func (f *fasilitasUsecase) DeleteFasilitas(ctx context.Context, id string) error {
	fasilitas, err := f.repo.FindFasilitas(ctx, id)
	if err != nil {
		return err
	}

	if fasilitas.GambarFasilitas.ImageKey != "" {
		if err := f.s3.Delete(fasilitas.GambarFasilitas.ImageKey); err != nil {
			return err
		}
	}

	if err := f.repo.DeleteFasilitas(ctx, id); err != nil {
		return err
	}

	return nil
}

// FindAllFasilitas implements FasilitasUsecase.
func (f *fasilitasUsecase) FindAllFasilitas(ctx context.Context, limit int, page int) ([]entity.Fasilitas, *utils.PaginationData, error) {
	return f.repo.FindAllFasilitas(ctx, limit, page)
}

// FindFasilitas implements FasilitasUsecase.
func (f *fasilitasUsecase) FindFasilitas(ctx context.Context, id string) (*entity.Fasilitas, error) {
	fasilitas, err := f.repo.FindFasilitas(ctx, id)
	if err != nil {
		return nil, err
	}

	return fasilitas, nil
}

// InsertFasilitas implements FasilitasUsecase.
func (f *fasilitasUsecase) InsertFasilitas(ctx context.Context, req entity.CreateFasilitas) (*entity.Fasilitas, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	var hargaFasilitas []entity.HargaFasilitas
	err := json.Unmarshal([]byte(req.Harga), &hargaFasilitas)
	if err != nil {
		return nil, err
	}

	fasilitas := entity.Fasilitas{
		NamaFasilitas:      req.Nama,
		DeskripsiFasilitas: req.Deskripsi,
		HargaFasilitas:     hargaFasilitas,
		CreatedAt:          utils.GetTimeNow(),
		UpdatedAt:          utils.GetTimeNow(),
	}

	if req.Gambar != nil {
		res, err := f.s3.Upload(req.Gambar, FasilitasFolder)
		if err != nil {
			return nil, err
		}

		fasilitas.GambarFasilitas = entity.GambarFasilitas{
			ImageKey: *res.Key,
			ImageURL: fmt.Sprintf("%s%s", storage.CloudFrontURLBase, *res.Key),
		}
	}

	data, err := f.repo.InsertFasilitas(ctx, fasilitas)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// UpdateFasilitas implements FasilitasUsecase.
func (f *fasilitasUsecase) UpdateFasilitas(ctx context.Context, id string, req entity.UpdateFasilitas) (*entity.Fasilitas, error) {
	fasilitas, err := f.repo.FindFasilitas(ctx, id)
	if err != nil {
		return nil, err
	}

	var hargaFasilitas []entity.HargaFasilitas
	err = json.Unmarshal([]byte(req.Harga), &hargaFasilitas)
	if err != nil {
		return nil, err
	}

	updatedFasilitas := entity.Fasilitas{
		NamaFasilitas:      req.Nama,
		GambarFasilitas:    fasilitas.GambarFasilitas,
		DeskripsiFasilitas: req.Deskripsi,
		HargaFasilitas:     hargaFasilitas,
		UpdatedAt:          utils.GetTimeNow(),
	}

	if req.Gambar != nil {
		var res *storage.S3Response

		if fasilitas.GambarFasilitas.ImageKey != "" {
			oldKey := fasilitas.GambarFasilitas.ImageKey
			res, err = f.s3.Update(oldKey, req.Gambar, FasilitasFolder)
		} else {
			res, err = f.s3.Upload(req.Gambar, FasilitasFolder)
		}

		if err != nil {
			return nil, err
		}

		updatedFasilitas.GambarFasilitas = entity.GambarFasilitas{
			ImageKey: *res.Key,
			ImageURL: fmt.Sprintf("%s%s", storage.CloudFrontURLBase, *res.Key),
		}
	} else {
		updatedFasilitas.GambarFasilitas = fasilitas.GambarFasilitas
	}

	data, err := f.repo.UpdateFasilitas(ctx, id, updatedFasilitas)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func NewFasilitasUsecase(repo repository.FasilitasRepository, s3 storage.S3Service) FasilitasUsecase {
	return &fasilitasUsecase{repo, s3}
}
