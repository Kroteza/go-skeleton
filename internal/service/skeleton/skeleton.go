package skeleton

import (
	"context"

	"github.com/pkg/errors"
	testingEntity "go-skeleton/internal/entity/testing"
)

// Data ...
// Masukkan function dari package data ke dalam interface ini
type Data interface {
	GetAllData(ctx context.Context) ([]testingEntity.Testing, error)
	GetDataByID(ctx context.Context, ID string) (testingEntity.Testing, error)
}

// Service ...
// Tambahkan variable sesuai banyak data layer yang dibutuhkan
type Service struct {
	data Data
}

// New ...
// Tambahkan parameter sesuai banyak data layer yang dibutuhkan
func New(data Data) Service {
	// Assign variable dari parameter ke object
	return Service{
		data: data,
	}
}

// GetAllData ...
func (s Service) GetAllData(ctx context.Context) ([]testingEntity.Testing, error) {
	hasil, err := s.data.GetAllData(ctx)
	if err != nil {
		return hasil, errors.Wrap(err, "[SERVICE][SKELETON][GetAllData]")
	}
	return hasil, err
}

// GetDataByID ...
func (s Service) GetDataByID(ctx context.Context, ID string) (testingEntity.Testing, error)  {

	hasil, err := s.data.GetDataByID(ctx, ID)
	if err != nil{
		return hasil, errors.Wrap(err, "[SERVICE][SKELETON][GetDataByID]")
	}
	return hasil, err
}
