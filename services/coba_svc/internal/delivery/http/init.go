package http

import (
	"net/http"

	"github.com/gorilla/mux"

	"coba/pkg/logger"
	"coba/services/coba_svc/domain/usecase/command"
	"coba/services/coba_svc/internal/delivery/http/handler"
)

func InitRoutes(router *mux.Router, c *command.Wrapper) {
	logger.Logger.Info("initialize routes...")

	var (
		path    = "/v1"
		handler = handler.NewHandler(c)
	)
	router = router.PathPrefix(path).Subrouter()

	router.HandleFunc("/save-coba", handler.SaveCobaAja).Methods(http.MethodPost)
}
