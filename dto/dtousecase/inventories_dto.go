package dtousecase

import (
	"time"

	"github.com/shopspring/decimal"
)

type Inventories struct {
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
	CreatedAt              time.Time       `json:"created_at"`
	UpdatedAt              time.Time       `json:"updated_at"`
}

type InventoriesParams struct {
	Inventories
	SortBy    string
	Search    string
	Sort      string
	Limit     int
	Page      int
	StartDate string
	EndDate   string
}
