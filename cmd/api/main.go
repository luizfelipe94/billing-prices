package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/luizfelipe94/billing-prices/internal"
	"github.com/luizfelipe94/billing-prices/internal/domain/repositories"
	"github.com/luizfelipe94/billing-prices/internal/infra"
	"github.com/luizfelipe94/billing-prices/internal/infra/http_router"
	"github.com/luizfelipe94/billing-prices/internal/infra/persistence"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	requiredVars := []string{"DATABASE_URL", "KAFKA_BROKER", "PORT"}
	for _, v := range requiredVars {
		if os.Getenv(v) == "" {
			log.Fatalf("Environment variable %s is not set", v)
		}
	}

	fmt.Println("Environment variables loaded successfully")
}

func main() {

	connStr := os.Getenv("DATABASE_URL")
	internal.ConnectToDB(connStr)
	defer internal.DB.Close()

	kafkaProducerPrices := infra.NewKafkaProducer([]string{os.Getenv("KAFKA_BROKER")}, "billing-usage-pricing")
	defer kafkaProducerPrices.Close()

	var priceRepository repositories.PriceRepository = persistence.NewPostgresPriceRepository(internal.DB)
	priceRouter := http_router.NewPriceRouter(priceRepository, internal.DB, kafkaProducerPrices)

	port := os.Getenv("PORT")
	router := http.NewServeMux()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "ok")
	})

	router.HandleFunc("POST /api/v1/prices", priceRouter.CreatePrice)
	router.HandleFunc("GET /api/v1/prices", priceRouter.ListPrices)

	fmt.Printf("Server running on port %s\n", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), router); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
