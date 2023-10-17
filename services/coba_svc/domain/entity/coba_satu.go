package entity

import (
	"time"

	"coba/pkg/logger"
)

// CobaSatu coba satu entity
type CobaSatu struct {
	id        string
	name      string
	age       int
	isActive  bool
	startDate time.Time
	createdAt time.Time
	updatedAt time.Time
	deletedAt *time.Time
}

// DTOCobaSatu dto for coba satu entity
type DTOCobaSatu struct {
	ID        string
	Name      string
	Age       int
	IsActive  bool
	StartDate time.Time
}

// NewCobaSatu build new entity coba satu
func NewCobaSatu(dto DTOCobaSatu) (*CobaSatu, error) {
	cobaSatu := &CobaSatu{
		id:        dto.ID,
		name:      dto.Name,
		age:       dto.Age,
		isActive:  dto.IsActive,
		startDate: dto.StartDate,
	}

	err := cobaSatu.validate()
	if err != nil {
		logger.Logger.Error(err)
		return nil, err
	}
	return cobaSatu, nil
}

func (data *CobaSatu) validate() error {
	return nil
}

// GetID get id value
func (data *CobaSatu) GetID() string {
	return data.id
}

// SetID set id value
func (data *CobaSatu) SetID(id string) (*CobaSatu, error) {
	data.id = id
	err := data.validate()
	if err != nil {
		logger.Logger.Error(err)
		return nil, err
	}
	return data, nil
}

// GetName get name value
func (data *CobaSatu) GetName() string {
	return data.name
}

// SetName set name value
func (data *CobaSatu) SetName(name string) (*CobaSatu, error) {
	data.name = name
	err := data.validate()
	if err != nil {
		logger.Logger.Error(err)
		return nil, err
	}
	return data, nil
}

// GetAge get age value
func (data *CobaSatu) GetAge() int {
	return data.age
}

// SetAge set age value
func (data *CobaSatu) SetAge(age int) (*CobaSatu, error) {
	data.age = age
	err := data.validate()
	if err != nil {
		logger.Logger.Error(err)
		return nil, err
	}
	return data, nil
}

// GetIsActive get isActive value
func (data *CobaSatu) GetIsActive() bool {
	return data.isActive
}

// SetIsActive set isActive value
func (data *CobaSatu) SetIsActive(isActive bool) (*CobaSatu, error) {
	data.isActive = isActive
	err := data.validate()
	if err != nil {
		logger.Logger.Error(err)
		return nil, err
	}
	return data, nil
}

// GetStartDate get startDate value
func (data *CobaSatu) GetStartDate() time.Time {
	return data.startDate
}

// SetStartDate set startDate value
func (data *CobaSatu) SetStartDate(startDate time.Time) (*CobaSatu, error) {
	data.startDate = startDate
	err := data.validate()
	if err != nil {
		logger.Logger.Error(err)
		return nil, err
	}
	return data, nil
}

// GetCreatedAt get createdAt value
func (data *CobaSatu) GetCreatedAt() time.Time {
	return data.createdAt
}

// SetCreatedAt set createdAt value
func (data *CobaSatu) SetCreatedAt(date time.Time) (*CobaSatu, error) {
	data.createdAt = date
	err := data.validate()
	if err != nil {
		logger.Logger.Error(err)
		return nil, err
	}
	return data, nil
}

// GetUpdatedAt get updatedAt value
func (data *CobaSatu) GetUpdatedAt() time.Time {
	return data.updatedAt
}

// SetUpdatedAt set updatedAt value
func (data *CobaSatu) SetUpdatedAt(date time.Time) (*CobaSatu, error) {
	data.updatedAt = date
	err := data.validate()
	if err != nil {
		logger.Logger.Error(err)
		return nil, err
	}
	return data, nil
}

// GetDeletedAt get deletedAt value
func (data *CobaSatu) GetDeletedAt() *time.Time {
	return data.deletedAt
}

// SetDeletedAt set deletedAt value
func (data *CobaSatu) SetDeletedAt(date *time.Time) (*CobaSatu, error) {
	data.deletedAt = date
	err := data.validate()
	if err != nil {
		logger.Logger.Error(err)
		return nil, err
	}
	return data, nil
}
