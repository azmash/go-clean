package main

import (
	// "github.com/azmash/go-clean/repository/grpc"
	"log"
	"net/http"

	httpRepo "github.com/azmash/go-clean/repository/http"
	"github.com/azmash/go-clean/usecase/shop"
	"github.com/julienschmidt/httprouter"

	"github.com/azmash/go-clean/config"
	deliveryHTTP "github.com/azmash/go-clean/delivery/http"
	"github.com/azmash/go-clean/repository/postgre"
)

func main() {

	// read config
	cfg := &config.Configuration{}
	config.ReadModuleConfig(cfg, "etc/")

	// repoGRPC := grpc.NewGRPC()
	repoHTTP := httpRepo.NewHTTP(cfg.URL)
	repoPostgre := postgre.NewDB(cfg.DB)

	router := httprouter.New()

	uc := shop.NewUsecase(repoHTTP.Kyc, repoHTTP.ShopScore, repoHTTP.Order, repoPostgre)
	h := deliveryHTTP.NewHTTP(router, uc)
	h.SetEndpoint()

	log.Fatal(http.ListenAndServe(":7070", router))

}
