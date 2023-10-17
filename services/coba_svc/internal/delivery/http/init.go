package http

import (
	"net/http"

	"coba/pkg/logger"

	"github.com/gorilla/mux"

	"coba/services/coba_svc/domain/usecase"
	"coba/services/coba_svc/internal/delivery/http/handler"
)

func InitRoutes(router *mux.Router, uc *usecase.Wrapper) {
	logger.Logger.Info("initialize routes...")

	var (
		path    = "/v1"
		handler = handler.NewHandler(uc)
	)
	router = router.PathPrefix(path).Subrouter()

	router.HandleFunc("/save-coba", handler.SaveCobaSatu).Methods(http.MethodPost)
}
