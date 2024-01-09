package dto

import (
	"time"

	"github.com/shopspring/decimal"
)

type InventoryCreateRequest struct {
	NamaBarang             string          `json:"nama_barang" binding:"required"`
	Tipe                   string          `json:"tipe" binding:"required"`
	TanggalPerolehan       string          `json:"tanggal_perolehan"`
	Kondisi                string          `json:"kondisi" binding:"required"`
	SumberDana             string          `json:"sumber_dana"`
	KodeBarang             string          `json:"kode_barang"`
	HargaSatuan            decimal.Decimal `json:"harga_satuan"`
	KodeInventaris         string          `json:"kode_inventaris"`
	NamaRuangan            string          `json:"nama_ruangan" binding:"required"`
	NamaPengguna           string          `json:"nama_pengguna"`
	Unit                   string          `json:"unit"`
	Keterangan             string          `json:"keterangan"`
	Status                 string          `json:"status" binding:"required"`
	PengembalianLaptopLama string          `json:"pengembalian_laptop_lama"`
}

type InventoryCreateResponse struct {
	ID                     int             `json:"id"`
	NamaBarang             string          `json:"nama_barang"`
	Tipe                   string          `json:"tipe"`
	TanggalPerolehan       string          `json:"tanggal_perolehan"`
	Kondisi                string          `json:"kondisi"`
	SumberDana             string          `json:"sumber_dana"`
	KodeBarang             string          `json:"kode_barang"`
	HargaSatuan            decimal.Decimal `json:"harga_satuan"`
	KodeInventaris         string          `json:"kode_inventaris"`
	NamaRuangan            string          `json:"nama_ruangan"`
	NamaPengguna           string          `json:"nama_pengguna"`
	Unit                   string          `json:"unit"`
	Keterangan             string          `json:"keterangan"`
	Status                 string          `json:"status"`
	PengembalianLaptopLama string          `json:"pengembalian_laptop_lama"`
	CreatedAt              time.Time       `gorm:"not null;default:CURRENT_TIMESTAMP;type:timestamp"`
	UpdatedAt              time.Time       `gorm:"not null;default:CURRENT_TIMESTAMP;type:timestamp"`
}

type InventoryFindRequest struct {
	ID                     int             `json:"id"`
	NamaBarang             string          `json:"nama_barang"`
	Tipe                   string          `json:"tipe"`
	TanggalPerolehan       string          `json:"tanggal_perolehan"`
	Kondisi                string          `json:"kondisi"`
	SumberDana             string          `json:"sumber_dana"`
	KodeBarang             string          `json:"kode_barang"`
	HargaSatuan            decimal.Decimal `json:"harga_satuan"`
	KodeInventaris         string          `json:"kode_inventaris"`
	NamaRuangan            string          `json:"nama_ruangan"`
	NamaPengguna           string          `json:"nama_pengguna"`
	Unit                   string          `json:"unit"`
	Keterangan             string          `json:"keterangan"`
	Status                 string          `json:"status"`
	PengembalianLaptopLama string          `json:"pengembalian_laptop_lama"`
}

type PaginationData struct {
	TotalPage   int `json:"total_page"`
	TotalItem   int `json:"total_item"`
	CurrentPage int `json:"current_page"`
	Limit       int `json:"limit"`
}

type JSONPagination struct {
	Data       any            `json:"data,omitempty"`
	Message    string         `json:"message,omitempty"`
	Pagination PaginationData `json:"pagination,omitempty"`
}
