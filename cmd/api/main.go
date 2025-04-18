package main

import (
	"fmt"
	"net/http"

	"github.com/luizfelipe94/billing-prices/internal"
	"github.com/luizfelipe94/billing-prices/internal/domain/repositories"
	"github.com/luizfelipe94/billing-prices/internal/infra"
	"github.com/luizfelipe94/billing-prices/internal/infra/http_router"
	"github.com/luizfelipe94/billing-prices/internal/infra/persistence"
)

func main() {

	connStr := "postgres://postgres:postgres@localhost:5432/pricesdb?sslmode=disable"
	internal.ConnectToDB(connStr)
	defer internal.DB.Close()

	kafkaProducerPrices := infra.NewKafkaProducer([]string{"localhost:9092"}, "billing-usage-pricing")
	defer kafkaProducerPrices.Close()

	var priceRepository repositories.PriceRepository = persistence.NewPostgresPriceRepository(internal.DB)

	priceRouter := http_router.NewPriceRouter(priceRepository, internal.DB, kafkaProducerPrices)

	port := 8081
	router := http.NewServeMux()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "ok")
	})

	router.HandleFunc("POST /api/v1/prices", priceRouter.CreatePrice)

	fmt.Printf("Server running on port %d\n", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), router); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
