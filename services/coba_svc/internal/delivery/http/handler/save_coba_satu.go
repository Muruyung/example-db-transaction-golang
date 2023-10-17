package handler

import (
	"context"
	"net/http"

	"github.com/Muruyung/go-utilities/converter"
	"github.com/docker/distribution/uuid"

	"coba/pkg/logger"
	"coba/pkg/utils"
	"coba/services/coba_svc/domain/usecase"
	"coba/services/coba_svc/internal/delivery/request"
)

func (handler *Handler) SaveCobaSatu(w http.ResponseWriter, r *http.Request) {
	var (
		ctx = context.WithValue(r.Context(), logger.ActivityID, uuid.Generate().String())
		req request.RequestCoba
	)
	r.Close = true

	err := utils.DecodeJSONBodyRequest(r, &req)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err)
		return
	}

	start, _ := converter.ConvertStringToDate(req.StartDate)

	err = handler.cobaSatuUC.CreateCobaSatu(ctx, usecase.DTOCobaSatu{
		Name:      req.Name,
		Age:       req.Age,
		IsActive:  req.IsActive,
		StartDate: start,
	})
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err)
		return
	}

	utils.RespondWithSuccess(w, http.StatusOK, nil)
}
