package repository

import (
	"context"
	"errors"
	"sbm-itb/dto/dtousecase"
	"sbm-itb/model"
	"sbm-itb/util"

	"gorm.io/gorm"
)

type InventoriesRepository interface {
	FindAll(ctx context.Context, req dtousecase.InventoriesParams) (*[]dtousecase.Inventories, int64, error)
	FindByID(ctx context.Context, req dtousecase.Inventories) (dtousecase.Inventories, error)
	Create(ctx context.Context, req dtousecase.Inventories) (dtousecase.Inventories, error)
	UpdateByID(ctx context.Context, req dtousecase.Inventories) (dtousecase.Inventories, error)
	DeleteByID(ctx context.Context, req dtousecase.Inventories) (dtousecase.Inventories, error)
	FindByName(ctx context.Context, req dtousecase.Inventories) (dtousecase.Inventories, error)
	FindByKodeInventaris(ctx context.Context, req dtousecase.Inventories) (dtousecase.Inventories, error)
}

type inventoriesRepository struct {
	db *gorm.DB
}

func NewInventoriesRepository(db *gorm.DB) InventoriesRepository {
	return &inventoriesRepository{db: db}
}

func (r *inventoriesRepository) FindAll(ctx context.Context, req dtousecase.InventoriesParams) (*[]dtousecase.Inventories, int64, error) {
	inventories := []dtousecase.Inventories{}
	var totalItems int64

	q := `
	select * from inventories i
	`

	query := r.db.WithContext(ctx).Table("(?) as t", gorm.Expr(q))
	if req.StartDate != "" {
		query = query.Where("created_at >= ?", req.StartDate)
	}

	if req.EndDate != "" {
		req.EndDate += " 23:59:59"
		query = query.Where("created_at <= ?", req.EndDate)
	}

	if req.Search != "" {
		query = query.Where("LOWER(nama_barang) LIKE LOWER(?) OR LOWER(tipe) LIKE LOWER(?)", "%"+req.Search+"%", "%"+req.Search+"%")
	}

	if req.Inventories.NamaBarang != "" {
		query = query.Where("LOWER(nama_barang) LIKE LOWER(?)", "%"+req.Inventories.NamaBarang+"%")
	}

	if req.Inventories.Tipe != "" {
		query = query.Where("LOWER(tipe) LIKE LOWER(?)", "%"+req.Inventories.Tipe+"%")
	}

	if req.Inventories.Kondisi != "" {
		query = query.Where("LOWER(kondisi) LIKE LOWER(?)", "%"+req.Inventories.Kondisi+"%")
	}

	if req.Inventories.SumberDana != "" {
		query = query.Where("LOWER(sumber_dana) LIKE LOWER(?)", "%"+req.Inventories.SumberDana+"%")
	}

	if req.Inventories.KodeBarang != "" {
		query = query.Where("kode_barang = ?", req.Inventories.KodeBarang)
	}

	if req.Inventories.KodeInventaris != "" {
		query = query.Where("kode_inventaris = ?", req.Inventories.KodeInventaris)
	}

	if req.Inventories.NamaRuangan != "" {
		query = query.Where("LOWER(nama_ruangan) LIKE LOWER(?)", "%"+req.Inventories.NamaRuangan+"%")
	}

	if req.Inventories.NamaPengguna != "" {
		query = query.Where("LOWER(nama_pengguna) LIKE LOWER(?)", "%"+req.Inventories.NamaPengguna+"%")
	}

	if req.Inventories.Unit != "" {
		query = query.Where("unit = ?", req.Inventories.Unit)
	}

	if req.Inventories.Status != "" {
		query = query.Where("status = ?", req.Inventories.Status)
	}

	if req.Inventories.PengembalianLaptopLama != "" {
		query = query.Where("pengembalian_laptop_lama = ?", req.Inventories.PengembalianLaptopLama)
	}

	if err := query.Count(&totalItems).Error; err != nil {
		return nil, 0, err
	}

	query = query.Order(req.SortBy + " " + req.Sort)
	offset := (req.Page - 1) * req.Limit
	query = query.Offset(offset).Limit(req.Limit)

	if err := query.Find(&inventories).Error; err != nil {
		return nil, 0, err
	}

	return &inventories, totalItems, nil
}

func (r *inventoriesRepository) FindByID(ctx context.Context, req dtousecase.Inventories) (dtousecase.Inventories, error) {
	inventories := dtousecase.Inventories{}

	err := r.db.WithContext(ctx).Where("id = ?", req.ID).First(&inventories).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return inventories, util.ErrNoRecordFound
	}

	if err != nil {
		return inventories, err
	}

	return inventories, nil
}

func (r *inventoriesRepository) FindByName(ctx context.Context, req dtousecase.Inventories) (dtousecase.Inventories, error) {
	inventories := dtousecase.Inventories{}

	err := r.db.WithContext(ctx).Where("nama_barang = ?", req.NamaBarang).First(&inventories).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return inventories, util.ErrNoRecordFound
	}

	if err != nil {
		return inventories, err
	}

	return inventories, nil
}

func (r *inventoriesRepository) FindByKodeInventaris(ctx context.Context, req dtousecase.Inventories) (dtousecase.Inventories, error) {
	inventories := dtousecase.Inventories{}

	err := r.db.WithContext(ctx).Where("kode_inventaris = ?", req.KodeInventaris).First(&inventories).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return inventories, util.ErrNoRecordFound
	}

	if err != nil {
		return inventories, err
	}

	return inventories, nil
}

func (r *inventoriesRepository) Create(ctx context.Context, req dtousecase.Inventories) (dtousecase.Inventories, error) {
	inventories := dtousecase.Inventories{}
	err := r.db.WithContext(ctx).Create(&req).Scan(&inventories).Error
	if err != nil {
		return req, err
	}

	return inventories, nil
}

func (r *inventoriesRepository) UpdateByID(ctx context.Context, req dtousecase.Inventories) (dtousecase.Inventories, error) {
	inventories := dtousecase.Inventories{}

	err := r.db.WithContext(ctx).
		Model(&model.Inventories{}).
		Where("LOWER(id) = LOWER(?)", req.ID).
		Updates(&req).
		Scan(&inventories).Error

	if err != nil {
		return inventories, err
	}

	return inventories, nil
}

func (r *inventoriesRepository) DeleteByID(ctx context.Context, req dtousecase.Inventories) (dtousecase.Inventories, error) {
	inventories := dtousecase.Inventories{}

	err := r.db.WithContext(ctx).Model(&model.Inventories{}).Where("id = ?", req.ID).Delete(&inventories).Error
	if err != nil {
		return inventories, err
	}

	return inventories, nil
}
