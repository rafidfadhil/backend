package usecase

import (
	"context"

	"github.com/BIC-Final-Project/backend/internal/asset/entity"
	"github.com/BIC-Final-Project/backend/internal/asset/repository"
	"github.com/BIC-Final-Project/backend/pkg/utils"
)

type VendorUsecase interface {
	InsertVendor(ctx context.Context, req entity.CreateVendor) (*entity.Vendor, error)
	FindAllVendors(ctx context.Context, limit int, page int) ([]entity.Vendor, *utils.PaginationData, error)
	FindVendor(ctx context.Context, vendorID string) (*entity.Vendor, error)
	UpdateVendor(ctx context.Context, vendorID string, req entity.UpdateVendor) (*entity.Vendor, error)
	DeleteVendor(ctx context.Context, vendorID string) error
}

type vendorUsecase struct {
	repo repository.VendorRepository
}

var VendorFolder = "vendors-bic"

func (v *vendorUsecase) InsertVendor(ctx context.Context, req entity.CreateVendor) (*entity.Vendor, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	vendor := entity.Vendor{
		NamaVendor:   req.NamaVendor,
		TelpVendor:   req.TelpVendor,
		AlamatVendor: req.AlamatVendor,
		JenisVendor:  req.JenisVendor,
	}

	data, err := v.repo.InsertVendor(ctx, vendor)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (v *vendorUsecase) FindAllVendors(ctx context.Context, limit int, page int) ([]entity.Vendor, *utils.PaginationData, error) {
	return v.repo.FindAllVendors(ctx, limit, page)
}

func (v *vendorUsecase) FindVendor(ctx context.Context, vendorID string) (*entity.Vendor, error) {
	vendor, err := v.repo.FindVendor(ctx, vendorID)

	if err != nil {
		return nil, err
	}

	return vendor, nil
}

func (v *vendorUsecase) UpdateVendor(ctx context.Context, vendorID string, req entity.UpdateVendor) (*entity.Vendor, error) {
	_, err := v.repo.FindVendor(ctx, vendorID)
	if err != nil {
		return nil, err
	}

	updatedVendor := entity.Vendor{
		NamaVendor:   req.NamaVendor,
		TelpVendor:   req.TelpVendor,
		AlamatVendor: req.AlamatVendor,
		JenisVendor:  req.JenisVendor,
	}

	data, err := v.repo.UpdateVendor(ctx, vendorID, updatedVendor)

	if err != nil {
		return nil, err

	}

	return data, nil
}

func (v *vendorUsecase) DeleteVendor(ctx context.Context, vendorID string) error {
	_, err := v.repo.FindVendor(ctx, vendorID)

	if err != nil {
		return err
	}

	if err := v.repo.DeleteVendor(ctx, vendorID); err != nil {
		return err
	}

	return nil
}

func NewVendorUsecase(repo repository.VendorRepository) VendorUsecase {
	return &vendorUsecase{repo}
}
