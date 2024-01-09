package model

import (
	"time"

	"github.com/shopspring/decimal"
)

type Inventories struct {
	ID                     int             `gorm:"primaryKey;not null,autoIncrement;serial"`
	NamaBarang             string          `gorm:"type:varchar(255);not null"`
	Tipe                   string          `gorm:"type:varchar(255);not null"`
	TanggalPerolehan       string          `gorm:"type:varchar(255);default:null"`
	Kondisi                string          `gorm:"type:varchar(255);not null"`
	SumberDana             string          `gorm:"type:varchar(255);default:null"`
	KodeBarang             string          `gorm:"type:varchar(255);default:null"`
	HargaSatuan            decimal.Decimal `gorm:"type:decimal"`
	KodeInventaris         string          `gorm:"type:varchar(255);not null;default:-"`
	NamaRuangan            string          `gorm:"type:varchar(255);not null"`
	NamaPengguna           string          `gorm:"type:varchar(255);default:null"`
	Unit                   string          `gorm:"type:varchar(255);default:null"`
	Keterangan             string          `gorm:"type:varchar(255);default:null"`
	Status                 string          `gorm:"type:varchar(255);default:null"`
	PengembalianLaptopLama string          `gorm:"type:varchar(255);default:null"`
	CreatedAt              time.Time       `gorm:"not null;default:CURRENT_TIMESTAMP;type:timestamp"`
	UpdatedAt              time.Time       `gorm:"not null;default:CURRENT_TIMESTAMP;type:timestamp"`
}
