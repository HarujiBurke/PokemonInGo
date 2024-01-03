package main

import (
	"PokemonInGo/controllers"
	"PokemonInGo/database"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

var DB *gorm.DB

func main() {

	// Load Configurations from config.json using Viper
	LoadAppConfig()

	// Initialize Database
	database.Connect(AppConfig.ConnectionString)
	database.Migrate()

	// Initialize the router
	router := mux.NewRouter().StrictSlash(true)

	// Register Routes
	RegisterProductRoutes(router)
	RegisterPokemonRoutes(router)
	// Start the server
	log.Printf("Starting Server on port %s", AppConfig.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", AppConfig.Port), router))
}

func RegisterProductRoutes(router *mux.Router) {
	router.HandleFunc("/api/products", controllers.GetProducts).Methods("GET")
	router.HandleFunc("/api/products/{id}", controllers.GetProductById).Methods("GET")
	router.HandleFunc("/api/products", controllers.CreateProduct).Methods("POST")
	router.HandleFunc("/api/products/{id}", controllers.UpdateProduct).Methods("PUT")
	router.HandleFunc("/api/products/{id}", controllers.DeleteProduct).Methods("DELETE")
}

func RegisterPokemonRoutes(router *mux.Router) {
	router.HandleFunc("/api/pokemons", controllers.GetPokemons).Methods("GET")
	router.HandleFunc("/api/pokemons/{id}", controllers.GetPokemonById).Methods("GET")
	router.HandleFunc("/api/pokemons", controllers.CreatePokemon).Methods("POST")
	router.HandleFunc("/api/pokemons/{id}", controllers.UpdatePokemon).Methods("PUT")
	router.HandleFunc("/api/pokemons/{id}", controllers.DeletePokemon).Methods("DELETE")
}
