package handler

import (
	"net/http"
	"sbm-itb/constants"
	"sbm-itb/dto"
	"sbm-itb/dto/dtousecase"
	"sbm-itb/usecase"
	"sbm-itb/util"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type InventoriesHandler struct {
	inventoriesUsecase usecase.InventoriesUsecase
}

func NewInventoriesHandler(inventoriesUsecase usecase.InventoriesUsecase) *InventoriesHandler {
	return &InventoriesHandler{inventoriesUsecase: inventoriesUsecase}
}

func (h *InventoriesHandler) FindAll(c *gin.Context) {
	limit, err := strconv.Atoi(c.DefaultQuery("limit", "30"))
	if err != nil {
		c.Error(err)
		return
	}

	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		c.Error(err)
		return
	}

	sortBy := c.DefaultQuery("sortBy", "created_at")
	sort := c.DefaultQuery("sort", "desc")
	startDate := c.DefaultQuery("startDate", "")
	endDate := c.DefaultQuery("endDate", "")
	search := c.DefaultQuery("search", "")
	tipe := c.DefaultQuery("tipe", "")
	kondisi := c.DefaultQuery("kondisi", "")
	sumberDana := c.DefaultQuery("sumberDana", "")
	namaBarang := c.DefaultQuery("namaBarang", "")
	kodeBarang := c.DefaultQuery("kodeBarang", "")
	kodeInventaris := c.DefaultQuery("kodeInventaris", "")
	namaRuangan := c.DefaultQuery("namaRuangan", "")
	namaPengguna := c.DefaultQuery("namaPengguna", "")
	unit := c.DefaultQuery("unit", "")
	status := c.DefaultQuery("status", "")
	pengembalianLaptopLama := c.DefaultQuery("pengembalianLaptopLama", "")

	if valid := util.IsDateValid(startDate); !valid && startDate != "" {
		c.Error(util.ErrInvalidDateFormat)
		return
	}

	if valid := util.IsDateValid(endDate); !valid && endDate != "" {
		c.Error(util.ErrInvalidDateFormat)
		return
	}

	if kondisi != "" && kondisi != constants.KondisiBaik && kondisi != constants.KondisiRusakRingan && kondisi != constants.KondisiRusakBerat {
		c.Error(util.ErrInvalidKondisi)
		return
	}

	if status != "" && status != constants.StatusDigunakan && status != constants.StatusTidakDigunakan {
		c.Error(util.ErrInvalidStatus)
		return
	}

	if pengembalianLaptopLama != "" && pengembalianLaptopLama != constants.PengembalianLaptopSudah && pengembalianLaptopLama != constants.PengembalianLaptopBelum {
		c.Error(util.ErrInvalidPengembalianLaptopLama)
		return
	}

	switch sortBy {
	case "date":
		sortBy = "created_at"
	}

	reqData := dtousecase.InventoriesParams{
		SortBy:    sortBy,
		Sort:      sort,
		Search:    search,
		Limit:     limit,
		Page:      page,
		StartDate: startDate,
		EndDate:   endDate,
		Inventories: dtousecase.Inventories{
			Tipe:                   tipe,
			Kondisi:                kondisi,
			SumberDana:             sumberDana,
			KodeBarang:             kodeBarang,
			KodeInventaris:         kodeInventaris,
			NamaRuangan:            namaRuangan,
			NamaPengguna:           namaPengguna,
			Unit:                   unit,
			Status:                 status,
			PengembalianLaptopLama: pengembalianLaptopLama,
			NamaBarang:             namaBarang,
		},
	}

	inventaris, pagination, err := h.inventoriesUsecase.FindAll(c.Request.Context(), reqData)

	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, dto.JSONPagination{Data: *inventaris, Pagination: *pagination})
}

func (h *InventoriesHandler) Create(c *gin.Context) {
	var req dto.InventoryCreateRequest

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.Error(util.ErrInvalidInput)
		return
	}

	reqData := dtousecase.Inventories{
		NamaBarang:             strings.TrimSpace(req.NamaBarang),
		Tipe:                   strings.TrimSpace(req.Tipe),
		TanggalPerolehan:       strings.TrimSpace(req.TanggalPerolehan),
		Kondisi:                strings.TrimSpace(req.Kondisi),
		SumberDana:             strings.TrimSpace(req.SumberDana),
		KodeBarang:             strings.TrimSpace(req.KodeBarang),
		HargaSatuan:            req.HargaSatuan,
		KodeInventaris:         strings.TrimSpace(req.KodeInventaris),
		NamaRuangan:            strings.TrimSpace(req.NamaRuangan),
		NamaPengguna:           strings.TrimSpace(req.NamaPengguna),
		Unit:                   strings.TrimSpace(req.Unit),
		Keterangan:             strings.TrimSpace(req.Keterangan),
		Status:                 strings.TrimSpace(req.Status),
		PengembalianLaptopLama: strings.TrimSpace(req.PengembalianLaptopLama),
	}

	inventaris, err := h.inventoriesUsecase.Create(c.Request.Context(), reqData)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, dto.JSONResponse{Message: "Successfully create data", Data: inventaris})
}

func (h *InventoriesHandler) FindByID(c *gin.Context) {
	barangId := c.Param("barangId")

	id, err := strconv.Atoi(barangId)
	if err != nil {
		c.Error(util.ErrInvalidInput)
		return
	}

	reqData := dtousecase.Inventories{
		ID: id,
	}

	inventaris, err := h.inventoriesUsecase.FindByID(c.Request.Context(), reqData)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, dto.JSONResponse{Message: "Successfully get data", Data: inventaris})
}

func (h *InventoriesHandler) UpdateByID(c *gin.Context) {
	barangId := c.Param("barangId")

	id, err := strconv.Atoi(barangId)
	if err != nil {
		c.Error(util.ErrInvalidInput)
		return
	}

	var req dto.InventoryCreateRequest

	err = c.ShouldBindJSON(&req)
	if err != nil {
		c.Error(util.ErrInvalidInput)
		return
	}

	reqData := dtousecase.Inventories{
		ID:                     id,
		NamaBarang:             strings.TrimSpace(req.NamaBarang),
		Tipe:                   strings.TrimSpace(req.Tipe),
		TanggalPerolehan:       strings.TrimSpace(req.TanggalPerolehan),
		Kondisi:                strings.TrimSpace(req.Kondisi),
		SumberDana:             strings.TrimSpace(req.SumberDana),
		KodeBarang:             strings.TrimSpace(req.KodeBarang),
		HargaSatuan:            req.HargaSatuan,
		KodeInventaris:         strings.TrimSpace(req.KodeInventaris),
		NamaRuangan:            strings.TrimSpace(req.NamaRuangan),
		NamaPengguna:           strings.TrimSpace(req.NamaPengguna),
		Unit:                   strings.TrimSpace(req.Unit),
		Keterangan:             strings.TrimSpace(req.Keterangan),
		Status:                 strings.TrimSpace(req.Status),
		PengembalianLaptopLama: strings.TrimSpace(req.PengembalianLaptopLama),
	}

	inventaris, err := h.inventoriesUsecase.UpdateByID(c.Request.Context(), reqData)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, dto.JSONResponse{Message: "Successfully update data", Data: inventaris})
}

func (h *InventoriesHandler) DeleteByID(c *gin.Context) {
	barangId := c.Param("barangId")

	id, err := strconv.Atoi(barangId)
	if err != nil {
		c.Error(util.ErrInvalidInput)
		return
	}

	reqData := dtousecase.Inventories{
		ID: id,
	}

	_, err = h.inventoriesUsecase.DeleteByID(c.Request.Context(), reqData)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, dto.JSONResponse{Message: "Successfully delete data"})
}
