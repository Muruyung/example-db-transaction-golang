package mapper

import (
	"time"

	"coba/pkg/logger"

	"coba/services/coba_svc/domain/entity"
	"coba/services/coba_svc/domain/repository"
	"coba/services/coba_svc/internal/repository/models"
)

type cobaSatuMapperInteractor struct {
	cobaSatuModels *models.CobaSatuModels
	cobaSatuEntity *entity.CobaSatu
}

// NewCobaSatuMapper create new coba satu models mapper
func NewCobaSatuMapper(cobaSatuEntity *entity.CobaSatu, cobaSatuModels *models.CobaSatuModels) repository.MapperCommon {
	return &cobaSatuMapperInteractor{
		cobaSatuEntity: cobaSatuEntity,
		cobaSatuModels: cobaSatuModels,
	}
}

// MapDomainToModels map domain to models coba satu
func (mapper *cobaSatuMapperInteractor) MapDomainToModels() repository.ModelsCommon {
	repoModels := models.CobaSatuModels{
		ID:        mapper.cobaSatuEntity.GetID(),
		Name:      mapper.cobaSatuEntity.GetName(),
		Age:       mapper.cobaSatuEntity.GetAge(),
		IsActive:  mapper.cobaSatuEntity.GetIsActive(),
		StartDate: mapper.cobaSatuEntity.GetStartDate(),
		CreatedAt: mapper.cobaSatuEntity.GetCreatedAt(),
		UpdatedAt: time.Now(),
	}
	return repoModels
}

// MapModelsToDomain map models to domain coba satu
func (mapper *cobaSatuMapperInteractor) MapModelsToDomain(entityStruct interface{}) {
	domain, err := entity.NewCobaSatu(entity.DTOCobaSatu{
		ID:        mapper.cobaSatuModels.ID,
		Name:      mapper.cobaSatuModels.Name,
		Age:       mapper.cobaSatuModels.Age,
		IsActive:  mapper.cobaSatuModels.IsActive,
		StartDate: mapper.cobaSatuModels.StartDate,
	})
	if err != nil {
		logger.Logger.Warnf("Error: %v", err)
	}

	domain, err = domain.SetCreatedAt(mapper.cobaSatuModels.CreatedAt)
	if err != nil {
		logger.Logger.Warnf("Error: %v", err)
	}

	domain, err = domain.SetUpdatedAt(mapper.cobaSatuModels.UpdatedAt)
	if err != nil {
		logger.Logger.Warnf("Error: %v", err)
	}

	domain, err = domain.SetDeletedAt(mapper.cobaSatuModels.DeletedAt)
	if err != nil {
		logger.Logger.Warnf("Error: %v", err)
	}

	entityDomain := entityStruct.(*entity.CobaSatu)
	*entityDomain = *domain
}
