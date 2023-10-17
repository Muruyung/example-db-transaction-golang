package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	"coba/pkg/connection_elastic"
	configmysql "coba/pkg/database/mysql"
	"coba/pkg/dotenv"
	"coba/pkg/logger"
	httpmodule "coba/services/coba_svc/internal/delivery/http"
	mysqltx "coba/services/coba_svc/internal/repository/mysql/mysql_tx"
	"coba/services/coba_svc/internal/service"
	"coba/services/coba_svc/internal/usecase"
)

func main() {
	_ = godotenv.Load()
	elastic, err := connection_elastic.Init()
	if err != nil {
		log.Fatal(err)
	}

	elastic.SetIndex("coba_elastic")

	logger.InitLogger("local", "coba", elastic)
	var (
		port      = dotenv.APPPORT()
		mysqlDB   = configmysql.InitMysqlDB(dotenv.MYSQLCONFIG())
		repo      = mysqltx.NewMysqlTx(mysqlDB)
		svc       = service.Init(repo.Wrapper())
		uc        = usecase.Init(svc)
		originsOk = handlers.AllowedOrigins([]string{"*"})
		headersOk = handlers.AllowedHeaders([]string{"Content-Type", "Accept-Language", "Authorization",
			"X-Requested-With", "Ciam-Type", "X-Device", "X-App-Version", "Channel", "Device-Brand"})
		methodsOk = handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})
		credsOk   = handlers.AllowCredentials()
		router    = mux.NewRouter()
	)

	defer func(sql *sql.DB) {
		_ = mysqlDB.Close()
	}(mysqlDB)

	httpmodule.InitRoutes(router, uc)
	log.Fatal(http.ListenAndServe(port, handlers.CORS(originsOk, headersOk, methodsOk, credsOk)(router)))
}
