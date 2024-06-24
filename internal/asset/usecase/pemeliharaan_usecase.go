package usecase

import (
	"context"

	"github.com/BIC-Final-Project/backend/internal/asset/entity"
	"github.com/BIC-Final-Project/backend/internal/asset/repository"
	"github.com/BIC-Final-Project/backend/pkg/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PemeliharaanUsecase interface {
	InsertPelihara(ctx context.Context, pemeliharaan entity.CreatePelihara) (*entity.Pemeliharaan, error)
	FindAllPelihara(ctx context.Context, limit int, page int) ([]entity.Pemeliharaan, *utils.PaginationData, error)
	FindPelihara(ctx context.Context, id string) (*entity.Pemeliharaan, error)
	UpdatePelihara(ctx context.Context, id string, pemeliharaan entity.UpdatePelihara) (*entity.Pemeliharaan, error)
	DeletePelihara(ctx context.Context, id string) error
}

type pemeliharaanUsecase struct {
	repo repository.PemeliharaanRepository
}

func (p *pemeliharaanUsecase) InsertPelihara(ctx context.Context, req entity.CreatePelihara) (*entity.Pemeliharaan, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	rencanaidObjectID, _ := primitive.ObjectIDFromHex(req.RencanaID)

	pelihara_aset := &entity.Pemeliharaan{
		RencanaID:            rencanaidObjectID,
		KondisiStlhPerbaikan: req.KondisiStlhPerbaikan,
		StatusPemeliharaan:   req.StatusPemeliharaan,
		PenanggungJawab:      req.PenanggungJawab,
		Deskripsi:            req.Deskripsi,
	}

	data, err := p.repo.InsertPelihara(ctx, *pelihara_aset)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (p *pemeliharaanUsecase) FindAllPelihara(ctx context.Context, limit int, page int) ([]entity.Pemeliharaan, *utils.PaginationData, error) {
	return p.repo.FindAllPelihara(ctx, limit, page)
}

func (p *pemeliharaanUsecase) FindPelihara(ctx context.Context, id string) (*entity.Pemeliharaan, error) {
	pelihara, err := p.repo.FindPelihara(ctx, id)
	if err != nil {
		return nil, err
	}

	return pelihara, nil
}

func (p *pemeliharaanUsecase) UpdatePelihara(ctx context.Context, id string, req entity.UpdatePelihara) (*entity.Pemeliharaan, error) {
	_, err := p.repo.FindPelihara(ctx, id)
	if err != nil {
		return nil, err
	}

	updatedPelihara := entity.Pemeliharaan{
		KondisiStlhPerbaikan: req.KondisiStlhPerbaikan,
		StatusPemeliharaan:   req.StatusPemeliharaan,
		PenanggungJawab:      req.PenanggungJawab,
		Deskripsi:            req.Deskripsi,
	}

	data, err := p.repo.UpdatePelihara(ctx, id, updatedPelihara)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (p *pemeliharaanUsecase) DeletePelihara(ctx context.Context, id string) error {
	err := p.repo.DeletePelihara(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func NewPemeliharaanUsecase(repo repository.PemeliharaanRepository) PemeliharaanUsecase {
	return &pemeliharaanUsecase{repo}
}
