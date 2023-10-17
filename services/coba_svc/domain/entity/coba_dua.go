package entity

import (
	"time"

	"coba/pkg/logger"
)

// CobaDua coba dua entity
type CobaDua struct {
	id        string
	name      string
	age       int
	isActive  bool
	startDate time.Time
	createdAt time.Time
	updatedAt time.Time
	deletedAt *time.Time
}

// DTOCobaDua dto for coba dua entity
type DTOCobaDua struct {
	ID        string
	Name      string
	Age       int
	IsActive  bool
	StartDate time.Time
}

// NewCobaDua build new entity coba dua
func NewCobaDua(dto DTOCobaDua) (*CobaDua, error) {
	cobaDua := &CobaDua{
		id:        dto.ID,
		name:      dto.Name,
		age:       dto.Age,
		isActive:  dto.IsActive,
		startDate: dto.StartDate,
	}

	err := cobaDua.validate()
	if err != nil {
		logger.Logger.Error(err)
		return nil, err
	}
	return cobaDua, nil
}

func (data *CobaDua) validate() error {
	return nil
}

// GetID get id value
func (data *CobaDua) GetID() string {
	return data.id
}

// SetID set id value
func (data *CobaDua) SetID(id string) (*CobaDua, error) {
	data.id = id
	err := data.validate()
	if err != nil {
		logger.Logger.Error(err)
		return nil, err
	}
	return data, nil
}

// GetName get name value
func (data *CobaDua) GetName() string {
	return data.name
}

// SetName set name value
func (data *CobaDua) SetName(name string) (*CobaDua, error) {
	data.name = name
	err := data.validate()
	if err != nil {
		logger.Logger.Error(err)
		return nil, err
	}
	return data, nil
}

// GetAge get age value
func (data *CobaDua) GetAge() int {
	return data.age
}

// SetAge set age value
func (data *CobaDua) SetAge(age int) (*CobaDua, error) {
	data.age = age
	err := data.validate()
	if err != nil {
		logger.Logger.Error(err)
		return nil, err
	}
	return data, nil
}

// GetIsActive get isActive value
func (data *CobaDua) GetIsActive() bool {
	return data.isActive
}

// SetIsActive set isActive value
func (data *CobaDua) SetIsActive(isActive bool) (*CobaDua, error) {
	data.isActive = isActive
	err := data.validate()
	if err != nil {
		logger.Logger.Error(err)
		return nil, err
	}
	return data, nil
}

// GetStartDate get startDate value
func (data *CobaDua) GetStartDate() time.Time {
	return data.startDate
}

// SetStartDate set startDate value
func (data *CobaDua) SetStartDate(startDate time.Time) (*CobaDua, error) {
	data.startDate = startDate
	err := data.validate()
	if err != nil {
		logger.Logger.Error(err)
		return nil, err
	}
	return data, nil
}

// GetCreatedAt get createdAt value
func (data *CobaDua) GetCreatedAt() time.Time {
	return data.createdAt
}

// SetCreatedAt set createdAt value
func (data *CobaDua) SetCreatedAt(date time.Time) (*CobaDua, error) {
	data.createdAt = date
	err := data.validate()
	if err != nil {
		logger.Logger.Error(err)
		return nil, err
	}
	return data, nil
}

// GetUpdatedAt get updatedAt value
func (data *CobaDua) GetUpdatedAt() time.Time {
	return data.updatedAt
}

// SetUpdatedAt set updatedAt value
func (data *CobaDua) SetUpdatedAt(date time.Time) (*CobaDua, error) {
	data.updatedAt = date
	err := data.validate()
	if err != nil {
		logger.Logger.Error(err)
		return nil, err
	}
	return data, nil
}

// GetDeletedAt get deletedAt value
func (data *CobaDua) GetDeletedAt() *time.Time {
	return data.deletedAt
}

// SetDeletedAt set deletedAt value
func (data *CobaDua) SetDeletedAt(date *time.Time) (*CobaDua, error) {
	data.deletedAt = date
	err := data.validate()
	if err != nil {
		logger.Logger.Error(err)
		return nil, err
	}
	return data, nil
}
