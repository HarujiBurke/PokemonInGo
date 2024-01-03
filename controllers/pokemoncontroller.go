package controllers

import (
	"PokemonInGo/database"
	"PokemonInGo/entities"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func CreatePokemon(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var pokemon entities.Pokemon
	json.NewDecoder(r.Body).Decode(&pokemon)
	database.Instance.Create(&pokemon)
	json.NewEncoder(w).Encode(pokemon)
}

func GetPokemonById(w http.ResponseWriter, r *http.Request) {
	pokemonId := mux.Vars(r)["id"]
	if !checkIfPokemonExists(pokemonId) {
		json.NewEncoder(w).Encode("Pokemon Not Found!")
		return
	}
	var pokemon entities.Pokemon
	database.Instance.First(&pokemon, pokemonId)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pokemon)
}

func GetPokemons(w http.ResponseWriter, r *http.Request) {
	var pokemons []entities.Pokemon
	database.Instance.Find(&pokemons)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(pokemons)
}

func UpdatePokemon(w http.ResponseWriter, r *http.Request) {
	pokemonId := mux.Vars(r)["id"]
	if !checkIfPokemonExists(pokemonId) {
		json.NewEncoder(w).Encode("Pokemon Not Found!")
		return
	}
	var pokemon entities.Pokemon
	database.Instance.First(&pokemon, pokemonId)
	json.NewDecoder(r.Body).Decode(&pokemon)
	database.Instance.Save(&pokemon)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pokemon)
}

func DeletePokemon(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	pokemonId := mux.Vars(r)["id"]
	if !checkIfPokemonExists(pokemonId) {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Pokemon Not Found!")
		return
	}
	var pokemon entities.Pokemon
	database.Instance.Delete(&pokemon, pokemonId)
	json.NewEncoder(w).Encode("Pokemon Deleted Successfully!")
}

func checkIfPokemonExists(productId string) bool {
	var pokemon entities.Pokemon
	database.Instance.First(&pokemon, productId)
	return pokemon.ID != 0
}
