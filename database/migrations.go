package database

import (
	"log"
	"sbm-itb/constants"
	"sbm-itb/model"
	"time"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	err := db.Migrator().DropTable(&model.Inventories{})
	if err != nil {
		panic(err)
	}

	inventories := []model.Inventories{
		{
			NamaBarang:     "Access Point",
			Tipe:           "Ruckus",
			Kondisi:        constants.KondisiBaik,
			KodeInventaris: "SBM/IT/01/AP/2022",
			NamaRuangan:    "Ruang IT Lt.5 (Freeport)",
			Status:         constants.StatusTidakDigunakan,
			CreatedAt:      time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			NamaBarang:     "Access Point",
			Tipe:           "Ruckus R720",
			Kondisi:        constants.KondisiBaik,
			KodeInventaris: "SBM/IT/02/AP/2022",
			NamaRuangan:    "Ruang IT Lt.5 (Freeport)",
			Status:         constants.StatusTidakDigunakan,
			CreatedAt:      time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			NamaBarang:     "Access Point",
			Tipe:           "Ruckus R720",
			Kondisi:        constants.KondisiBaik,
			KodeInventaris: "SBM/IT/03/AP/2022",
			NamaRuangan:    "Ruang IT Lt.5 (Freeport)",
			Status:         constants.StatusTidakDigunakan,
			CreatedAt:      time.Date(2023, 2, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			NamaBarang:     "Access Point",
			Tipe:           "Ruckus R720",
			Kondisi:        constants.KondisiBaik,
			KodeInventaris: "SBM/IT/04/AP/2022",
			NamaRuangan:    "Ruang IT Lt.5 (Freeport)",
			Status:         constants.StatusTidakDigunakan,
			CreatedAt:      time.Date(2023, 12, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			NamaBarang:       "Access Point",
			Tipe:             "Asus RT-AC59U",
			TanggalPerolehan: "2022",
			Kondisi:          constants.KondisiBaik,
			SumberDana:       "RKA  IT 2022",
			HargaSatuan:      decimal.NewFromInt(900000),
			NamaRuangan:      "Ruang IT Lt.5 (Freeport)",
			KodeInventaris:   "SBM/IT/05/AP/2022",
			Status:           constants.StatusTidakDigunakan,
			CreatedAt:        time.Date(2023, 12, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			NamaBarang:       "Access Point",
			Tipe:             "Asus RT-AC59U",
			TanggalPerolehan: "2022",
			Kondisi:          constants.KondisiBaik,
			SumberDana:       "RKA  IT 2022",
			HargaSatuan:      decimal.NewFromInt(900000),
			KodeInventaris:   "SBM/IT/06/AP/2022",
			NamaRuangan:      "Ruang Dekan Lt. 1",
			Status:           constants.StatusDigunakan,
			CreatedAt:        time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			NamaBarang:       "Access Point",
			Tipe:             "Asus RT-AC59U",
			TanggalPerolehan: "2022",
			Kondisi:          constants.KondisiBaik,
			SumberDana:       "RKA  IT 2022",
			HargaSatuan:      decimal.NewFromInt(900000),
			KodeInventaris:   "SBM/IT/07/AP/2022",
			NamaRuangan:      "Meeting Room Lt.6 (Freeport)",
			Status:           constants.StatusDigunakan,
			CreatedAt:        time.Date(2024, 1, 2, 0, 0, 0, 0, time.UTC),
		},
		{
			NamaBarang:     "Adapter",
			Tipe:           "Lenovo 170 W",
			Kondisi:        constants.KondisiBaik,
			KodeInventaris: "SBM/IT/01/ADT/2022",
			NamaRuangan:    "Ruang IT Lt.5 (Freeport)",
			Status:         constants.StatusTidakDigunakan,
			CreatedAt:      time.Date(2023, 8, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			NamaBarang:     "Camera",
			Tipe:           "Aver 650 EX",
			Kondisi:        constants.KondisiBaik,
			KodeInventaris: "-",
			NamaRuangan:    "Auditorium Lt.2",
			Status:         constants.StatusDigunakan,
			CreatedAt:      time.Date(2022, 5, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			NamaBarang:     "Camera",
			Tipe:           "Logitech",
			Kondisi:        constants.KondisiBaik,
			KodeInventaris: "50044-606020458001-12",
			NamaRuangan:    "2302 Lt.3 (Freeport)",
			Status:         constants.StatusDigunakan,
			CreatedAt:      time.Date(2023, 11, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			NamaBarang:     "CPU",
			Tipe:           "Nimitz TR5000",
			Kondisi:        constants.KondisiBaik,
			KodeInventaris: "SBM/IT/01/CPU/2022",
			NamaRuangan:    "Ruang IT Lt.5 (Freeport)",
			Status:         constants.StatusTidakDigunakan,
			CreatedAt:      time.Date(2024, 1, 5, 0, 0, 0, 0, time.UTC),
		},
		{
			NamaBarang:     "CPU",
			Tipe:           "Nimitz TR10",
			Kondisi:        constants.KondisiRusakBerat,
			KodeInventaris: "SBM/IT/02/CPU/2022",
			NamaRuangan:    "Ruang IT Lt.5 (Freeport)",
			Status:         constants.StatusDigunakan,
			CreatedAt:      time.Date(2023, 7, 1, 0, 0, 0, 0, time.UTC),
		},
	}
	err = db.AutoMigrate(&model.Inventories{})
	if err != nil {
		panic(err)
	}

	db.Create(&inventories)

	log.Println("Successfully migrated database")
}
