package usecase

import (
	"context"

	"github.com/BIC-Final-Project/backend/internal/asset/entity"
	"github.com/BIC-Final-Project/backend/internal/asset/repository"
	"github.com/BIC-Final-Project/backend/pkg/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PerencanaanUsecase interface {
	InsertRencana(ctx context.Context, req entity.CreatePerencanaan) (*entity.Perencanaan, error)
	FindAllRencana(ctx context.Context, limit int, page int) ([]entity.Perencanaan, *utils.PaginationData, error)
	FindRencana(ctx context.Context, id string) (*entity.Perencanaan, error)
	UpdateRencana(ctx context.Context, id string, req entity.UpdatePerencanaan) (*entity.Perencanaan, error)
	DeleteRencana(ctx context.Context, id string) error
}

type perencanaanUsecase struct {
	repo repository.PerencanaanRepository
}

var FolderRencana = "perencanaan-bic"

func (p *perencanaanUsecase) InsertRencana(ctx context.Context, req entity.CreatePerencanaan) (*entity.Perencanaan, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	var perencanaan entity.Perencanaan

	asetObjectID, _ := primitive.ObjectIDFromHex(req.AsetID)
	vendorObjectID, _ := primitive.ObjectIDFromHex(req.VendorID)

	perencanaan = entity.Perencanaan{
		AsetID:         asetObjectID,
		VendorID:       vendorObjectID,
		KondisiAset:    req.KondisiAset,
		StatusAset:     req.StatusAset,
		TglPerencanaan: req.TglPerencanaan,
		UsiaAset:       req.UsiaAset,
		MaksUsiaAset:   req.MaksUsiaAset,
		Deskripsi:      req.Deskripsi,
	}

	data, err := p.repo.InsertRencana(ctx, perencanaan)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (p *perencanaanUsecase) FindAllRencana(ctx context.Context, limit int, page int) ([]entity.Perencanaan, *utils.PaginationData, error) {
	return p.repo.FindAllRencana(ctx, limit, page)
}

func (p *perencanaanUsecase) FindRencana(ctx context.Context, id string) (*entity.Perencanaan, error) {
	rencana, err := p.repo.FindRencana(ctx, id)
	if err != nil {
		return nil, err
	}

	return rencana, nil
}

func (p *perencanaanUsecase) UpdateRencana(ctx context.Context, id string, req entity.UpdatePerencanaan) (*entity.Perencanaan, error) {
	_, err := p.repo.FindRencana(ctx, id)
	if err != nil {
		return nil, err
	}

	updatedRencana := entity.Perencanaan{
		KondisiAset:    req.KondisiAset,
		StatusAset:     req.StatusAset,
		TglPerencanaan: req.TglPerencanaan,
		UsiaAset:       req.UsiaAset,
		MaksUsiaAset:   req.MaksUsiaAset,
		Deskripsi:      req.Deskripsi,
	}

	data, err := p.repo.UpdateRencana(ctx, id, updatedRencana)
	if err != nil {
		return nil, err
	}

	return data, nil

}

func (p *perencanaanUsecase) DeleteRencana(ctx context.Context, id string) error {
	_, err := p.repo.FindRencana(ctx, id)
	if err != nil {
		return err
	}

	if err = p.repo.DeleteRencana(ctx, id); err != nil {
		return err
	}

	return nil
}

func NewPerencanaanUsecase(repo repository.PerencanaanRepository) PerencanaanUsecase {
	return &perencanaanUsecase{repo}
}
