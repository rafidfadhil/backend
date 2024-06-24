package usecase

import (
	"context"

	"github.com/BIC-Final-Project/backend/internal/operational/entity"
	"github.com/BIC-Final-Project/backend/internal/operational/repository"
	"github.com/BIC-Final-Project/backend/pkg/utils"
)

type MembershipTypeUsecase interface {
	InsertMembershipType(ctx context.Context, req entity.CreateMembershipType) (*entity.MembershipType, error)
	FindAllMembershipType(ctx context.Context, jenisPaket string) ([]entity.MembershipType, error)
	FindMembershipType(ctx context.Context, id string) (*entity.MembershipType, error)
	UpdateMembershipType(ctx context.Context, id string, req entity.UpdateMembershipType) (*entity.MembershipType, error)
	DeleteMembershipType(ctx context.Context, id string) error
}

type membershipTypeUsecase struct {
	repo repository.MembershipTypeRepository
}

// DeleteMembershipType implements FasilitasUsecase.
func (m *membershipTypeUsecase) DeleteMembershipType(ctx context.Context, id string) error {
	_, err := m.repo.FindMembershipType(ctx, id)
	if err != nil {
		return err
	}

	if err := m.repo.DeleteMembershipType(ctx, id); err != nil {
		return err
	}

	return nil
}

// FindAllMembershipType implements FasilitasUsecase.
func (m *membershipTypeUsecase) FindAllMembershipType(ctx context.Context, jenisPaket string) ([]entity.MembershipType, error) {
	return m.repo.FindAllMembershipType(ctx, jenisPaket)
}

// FindMembershipType implements FasilitasUsecase.
func (m *membershipTypeUsecase) FindMembershipType(ctx context.Context, id string) (*entity.MembershipType, error) {
	fasilitas, err := m.repo.FindMembershipType(ctx, id)
	if err != nil {
		return nil, err
	}

	return fasilitas, nil
}

// InsertMembershipType implements FasilitasUsecase.
func (m *membershipTypeUsecase) InsertMembershipType(ctx context.Context, req entity.CreateMembershipType) (*entity.MembershipType, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	membershipType := entity.MembershipType{
		JenisPaket:               req.JenisPaket,
		JenisKeanggotaan:         req.JenisKeanggotaan,
		JumlahAnggotaYangBerlaku: req.JumlahAnggotaYangBerlaku,
		Harga:                    req.Harga,
		FasilitasMembership:      req.FasilitasMembership,
		CreatedAt:                utils.GetTimeNow(),
		UpdatedAt:                utils.GetTimeNow(),
	}

	data, err := m.repo.InsertMembershipType(ctx, membershipType)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// UpdateMembershipType implements FasilitasUsecase.
func (m *membershipTypeUsecase) UpdateMembershipType(ctx context.Context, id string, req entity.UpdateMembershipType) (*entity.MembershipType, error) {
	_, err := m.repo.FindMembershipType(ctx, id)
	if err != nil {
		return nil, err
	}

	updatedMembershipType := entity.MembershipType{
		JenisPaket:               req.JenisPaket,
		JenisKeanggotaan:         req.JenisKeanggotaan,
		JumlahAnggotaYangBerlaku: req.JumlahAnggotaYangBerlaku,
		Harga:                    req.Harga,
		FasilitasMembership:      req.FasilitasMembership,
		UpdatedAt:                utils.GetTimeNow(),
	}

	data, err := m.repo.UpdateMembershipType(ctx, id, updatedMembershipType)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func NewMembershipTypeUsecase(repo repository.MembershipTypeRepository) MembershipTypeUsecase {
	return &membershipTypeUsecase{repo}
}
