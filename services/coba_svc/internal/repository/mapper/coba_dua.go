package mapper

import (
	"time"

	"coba/pkg/logger"

	"coba/services/coba_svc/domain/entity"
	"coba/services/coba_svc/domain/repository"
	"coba/services/coba_svc/internal/repository/models"
)

type cobaDuaMapperInteractor struct {
	cobaDuaModels *models.CobaDuaModels
	cobaDuaEntity *entity.CobaDua
}

// NewCobaDuaMapper create new coba dua models mapper
func NewCobaDuaMapper(cobaDuaEntity *entity.CobaDua, cobaDuaModels *models.CobaDuaModels) repository.MapperCommon {
	return &cobaDuaMapperInteractor{
		cobaDuaEntity: cobaDuaEntity,
		cobaDuaModels: cobaDuaModels,
	}
}

// MapDomainToModels map domain to models coba dua
func (mapper *cobaDuaMapperInteractor) MapDomainToModels() repository.ModelsCommon {
	repoModels := models.CobaDuaModels{
		ID:        mapper.cobaDuaEntity.GetID(),
		Name:      mapper.cobaDuaEntity.GetName(),
		Age:       mapper.cobaDuaEntity.GetAge(),
		IsActive:  mapper.cobaDuaEntity.GetIsActive(),
		StartDate: mapper.cobaDuaEntity.GetStartDate(),
		CreatedAt: mapper.cobaDuaEntity.GetCreatedAt(),
		UpdatedAt: time.Now(),
	}
	return repoModels
}

// MapModelsToDomain map models to domain coba dua
func (mapper *cobaDuaMapperInteractor) MapModelsToDomain(entityStruct interface{}) {
	domain, err := entity.NewCobaDua(entity.DTOCobaDua{
		ID:        mapper.cobaDuaModels.ID,
		Name:      mapper.cobaDuaModels.Name,
		Age:       mapper.cobaDuaModels.Age,
		IsActive:  mapper.cobaDuaModels.IsActive,
		StartDate: mapper.cobaDuaModels.StartDate,
	})
	if err != nil {
		logger.Logger.Warnf("Error: %v", err)
	}

	domain, err = domain.SetCreatedAt(mapper.cobaDuaModels.CreatedAt)
	if err != nil {
		logger.Logger.Warnf("Error: %v", err)
	}

	domain, err = domain.SetUpdatedAt(mapper.cobaDuaModels.UpdatedAt)
	if err != nil {
		logger.Logger.Warnf("Error: %v", err)
	}

	domain, err = domain.SetDeletedAt(mapper.cobaDuaModels.DeletedAt)
	if err != nil {
		logger.Logger.Warnf("Error: %v", err)
	}

	entityDomain := entityStruct.(*entity.CobaDua)
	*entityDomain = *domain
}
