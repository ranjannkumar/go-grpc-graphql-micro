package main

import (
	"log"
	"net/http"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/kelseyhightower/envconfig"
)

type AppConfig struct {
	Accounturl string `envconfig:"ACCOUNT_SERVICE_URL"`
	CatalogURL string `envconfig:"CATALOG_SERVICE_URL"`
	OrderURL   string `envconfig:"ORDER_SERVICE_URL"`
}

func main() {
	var cfg AppConfig
	err := envconfig.Process("", &cfg)
	if err != nil {
		log.Fatalf("Error processing environment config: %v", err)
	}
	log.Printf("Config loaded: AccountURL=%s, CatalogURL=%s, OrderURL=%s", cfg.Accounturl, cfg.CatalogURL, cfg.OrderURL) 

	s,err:= NewGraphQLServer(cfg.Accounturl,cfg.CatalogURL,cfg.OrderURL)
	if err!=nil{
		log.Fatalf("Error creating GraphQL server: %v", err)
	}
	log.Println("GraphQL server starting and handlers being registered...")


	http.Handle("/graphql",handler.New(s.ToExecutableSchema()))
	http.Handle("/playground",playground.Handler("ranjan","/graphql"))

	log.Fatal(http.ListenAndServe(":8080",nil))
}