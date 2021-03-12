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
	GetDataByAge(ctx context.Context, Age string) (testingEntity.Testing, error)
	GetDataByBalance(ctx context.Context, Balance string) (testingEntity.Testing, error)
	InsertDataUser(ctx context.Context, singleTesting testingEntity.Testing) error
	UpdateDataUser(ctx context.Context, singleTesting testingEntity.Testing, ID string) (testingEntity.Testing, error)
	DeleteDataUser(ctx context.Context, singleTesting testingEntity.Testing, ID string) error
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

// GetDataByAge ...
func (s Service) GetDataByAge(ctx context.Context, Age string) (testingEntity.Testing, error) {

	hasil, err := s.data.GetDataByAge(ctx, Age)
	if err != nil{
		return hasil, errors.Wrap(err, "[SERVICE][SKELETON][GetDataByAge]")
	}
	return hasil, err
}

// GetDataByBalance ...
func (s Service) GetDataByBalance(ctx context.Context, Balance string) (testingEntity.Testing, error) {

	hasil, err := s.data.GetDataByBalance(ctx, Balance)
	if err != nil{
		return hasil, errors.Wrap(err, "[SERVICE][SKELETON][GetDataByBalance]")
	}
	return hasil, err
}

// InsertDataUser ...
func(s Service) InsertDataUser(ctx context.Context, singleTesting testingEntity.Testing) error{
	err:= s.data.InsertDataUser(ctx, singleTesting)
	if err != nil{
		return errors.Wrap(err, "[SERVICE][SKELETON][InsertDataUser]")
	}
	return err
}

// UpdateDataUser ...
func(s Service) UpdateDataUser(ctx context.Context, singleTesting testingEntity.Testing, ID string) (testingEntity.Testing, error){
	hasil, err:= s.data.UpdateDataUser(ctx, singleTesting, ID)
	if err != nil{
		return hasil, errors.Wrap(err, "[SERVICE][SKELETON][UpdateDataUser]")
	}
	return hasil, err
}

// DeleteDataUser ...
func (s Service) DeleteDataUser(ctx context.Context, singleTesting testingEntity.Testing, ID string) error {
	err:= s.data.DeleteDataUser(ctx, singleTesting, ID)
	if err != nil{
		return errors.Wrap(err, "[SERVICE][SKELETON][DeleteDataUser]")
	}
	return err
}