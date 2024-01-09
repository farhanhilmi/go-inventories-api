package usecase

import (
	"context"
	"errors"
	"sbm-itb/constants"
	"sbm-itb/dto"
	"sbm-itb/dto/dtousecase"
	"sbm-itb/repository"
	"sbm-itb/util"
)

type InventoriesUsecase interface {
	FindAll(ctx context.Context, req dtousecase.InventoriesParams) (*[]dtousecase.Inventories, *dto.PaginationData, error)
	FindByID(ctx context.Context, req dtousecase.Inventories) (dtousecase.Inventories, error)
	Create(ctx context.Context, req dtousecase.Inventories) (dtousecase.Inventories, error)
	UpdateByID(ctx context.Context, req dtousecase.Inventories) (dtousecase.Inventories, error)
	DeleteByID(ctx context.Context, req dtousecase.Inventories) (dtousecase.Inventories, error)
}

type inventoriesUsecase struct {
	inventoriesRepo repository.InventoriesRepository
}

func NewInventoriesUsecase(inventoriesRepo repository.InventoriesRepository) InventoriesUsecase {
	return &inventoriesUsecase{inventoriesRepo: inventoriesRepo}
}

func (u *inventoriesUsecase) FindAll(ctx context.Context, req dtousecase.InventoriesParams) (*[]dtousecase.Inventories, *dto.PaginationData, error) {
	items, totalItems, err := u.inventoriesRepo.FindAll(ctx, req)
	if err != nil {
		return nil, nil, err
	}

	pagination := dto.PaginationData{
		TotalItem:   int(totalItems),
		TotalPage:   (int(totalItems) + req.Limit - 1) / req.Limit,
		CurrentPage: req.Page,
		Limit:       req.Limit,
	}

	return items, &pagination, nil
}

func (u *inventoriesUsecase) FindByID(ctx context.Context, req dtousecase.Inventories) (dtousecase.Inventories, error) {
	res := dtousecase.Inventories{}
	inventory, err := u.inventoriesRepo.FindByID(ctx, req)

	if errors.Is(err, util.ErrNoRecordFound) {
		return res, util.ErrNoRecordFound
	}
	if err != nil {
		return res, err
	}

	return inventory, nil
}

func (u *inventoriesUsecase) Create(ctx context.Context, req dtousecase.Inventories) (dtousecase.Inventories, error) {
	res := dtousecase.Inventories{}

	if req.Status != constants.StatusDigunakan && req.Status != constants.StatusTidakDigunakan {
		return res, util.ErrInvalidStatus
	}

	if req.Kondisi != constants.KondisiBaik && req.Kondisi != constants.KondisiRusakRingan && req.Kondisi != constants.KondisiRusakBerat {
		return res, util.ErrInvalidKondisi
	}

	if req.PengembalianLaptopLama != "" && req.PengembalianLaptopLama != constants.PengembalianLaptopSudah && req.PengembalianLaptopLama != constants.PengembalianLaptopBelum {
		return res, util.ErrInvalidPengembalianLaptopLama
	}

	item, err := u.inventoriesRepo.FindByName(ctx, req)
	if err != nil && !errors.Is(err, util.ErrNoRecordFound) {
		return res, err
	}

	if item.ID != 0 {
		return res, util.ErrDuplicateData
	}

	itemInventaris, err := u.inventoriesRepo.FindByKodeInventaris(ctx, req)
	if err != nil && !errors.Is(err, util.ErrNoRecordFound) {
		return res, err
	}

	if itemInventaris.ID != 0 {
		return res, util.ErrInvatoryCodeDuplicate
	}

	inventory, err := u.inventoriesRepo.Create(ctx, req)
	if err != nil {
		return res, err
	}

	return inventory, nil
}

func (u *inventoriesUsecase) UpdateByID(ctx context.Context, req dtousecase.Inventories) (dtousecase.Inventories, error) {
	res := dtousecase.Inventories{}
	_, err := u.inventoriesRepo.FindByID(ctx, req)

	if errors.Is(err, util.ErrNoRecordFound) {
		return res, util.ErrNoRecordFound
	}
	if err != nil {
		return res, err
	}

	if req.Status != constants.StatusDigunakan && req.Status != constants.StatusTidakDigunakan {
		return res, util.ErrInvalidStatus
	}

	if req.Kondisi != constants.KondisiBaik && req.Kondisi != constants.KondisiRusakRingan && req.Kondisi != constants.KondisiRusakBerat {
		return res, util.ErrInvalidKondisi
	}

	if req.PengembalianLaptopLama != "" && req.PengembalianLaptopLama != constants.PengembalianLaptopSudah && req.PengembalianLaptopLama != constants.PengembalianLaptopBelum {
		return res, util.ErrInvalidPengembalianLaptopLama
	}

	item, err := u.inventoriesRepo.FindByName(ctx, req)
	if err != nil && !errors.Is(err, util.ErrNoRecordFound) {
		return res, err
	}

	if item.ID != 0 && item.NamaBarang != req.NamaBarang {
		return res, util.ErrDuplicateData
	}

	inventory, err := u.inventoriesRepo.UpdateByID(ctx, req)
	if err != nil {
		return res, err
	}

	return inventory, nil
}

func (u *inventoriesUsecase) DeleteByID(ctx context.Context, req dtousecase.Inventories) (dtousecase.Inventories, error) {
	res := dtousecase.Inventories{}
	_, err := u.inventoriesRepo.FindByID(ctx, req)

	if errors.Is(err, util.ErrNoRecordFound) {
		return res, util.ErrNoRecordFound
	}
	if err != nil {
		return res, err
	}

	inventory, err := u.inventoriesRepo.DeleteByID(ctx, req)
	if err != nil {
		return res, err
	}

	return inventory, nil
}
