package util

import "fmt"

type ErrorCode int

const (
	BadRequest ErrorCode = iota
	NotFound
	Unauthorized
)

type CustomError struct {
	Code    ErrorCode
	Message string
}

func (e CustomError) Error() string {
	return fmt.Sprintf("Code: %d, Message: %s", e.Code, e.Message)
}

func New(code ErrorCode, message string) error {
	return &CustomError{
		Code:    code,
		Message: message,
	}
}

var (
	ErrInvalidInput                  = &CustomError{Message: "Invalid input", Code: BadRequest}
	ErrNoRecordFound                 = &CustomError{Message: "not found", Code: NotFound}
	ErrInvalidDateFormat             = &CustomError{Message: "invalid date format. the date format is YYYY-MM-DD, e.g: 2023-09-10", Code: BadRequest}
	ErrInvalidKondisi                = &CustomError{Message: "kondisi harus salah satu dari 'baik', 'rusak ringan' atau 'rusak berat'", Code: BadRequest}
	ErrInvalidStatus                 = &CustomError{Message: "status harus salah satu dari 'digunakan', atau 'tidak digunaan'", Code: BadRequest}
	ErrInvalidPengembalianLaptopLama = &CustomError{Message: "pengembalian laptop lama harus salah satu dari 'sudah', atau 'belum'", Code: BadRequest}
	ErrDuplicateData                 = &CustomError{Message: "duplicate data", Code: BadRequest}
	ErrInvatoryCodeDuplicate         = &CustomError{Message: "kode inventaris sudah digunakan", Code: BadRequest}
)
