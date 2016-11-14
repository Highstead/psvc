package ddd

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	logic "github.com/highstead/psvc/internal/logic"
	"net/http"
	"strconv"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

const timeConst string = "2006-01-02"

func PlayersAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	js, err := json.Marshal(logic.GetPlayers())
	if err != nil {
		fmt.Println(w, "Bad Arguments")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(js)
}
func Players(w http.ResponseWriter, r *http.Request) {
	fmt.Println("HTTPGET Players")
	w.Header().Set("Content-Type", "application/json")
	players, err := logic.GetAllPlayersSeasons()
	if err != nil {
		fmt.Println(w, "Internal Exception")
		w.WriteHeader(http.StatusInternalServerError)
	}
	js, err := json.Marshal(*players)
	if err != nil {
		fmt.Println(w, "Bad Arguments")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(js)
}
func Goalies(w http.ResponseWriter, r *http.Request) {
	fmt.Println("HTTPGET Goalies")
	w.Header().Set("Content-Type", "application/json")
	players, err := logic.GetAllGoaliesSeasons()
	if err != nil {
		fmt.Println(w, "Internal Exception")
		w.WriteHeader(http.StatusInternalServerError)
	}
	js, err := json.Marshal(*players)
	if err != nil {
		fmt.Println(w, "Bad Arguments")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(js)
}
func PlayerSeason(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	//seasonId := vars["seasonId"]
	playerStr := vars["playerId"]
	playerId, err := strconv.Atoi(playerStr)
	if err != nil {
		fmt.Println(w, "Bad Arguments")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// TODO, cleanup this garbage code
	ps, err := logic.GetPlayerSeasons(playerId)
	if err == nil {
		fmt.Println("    Season request for player" + playerStr)
		js, err := json.Marshal(*ps)
		if err != nil {
			panic(err)
		} //Panic, couldn't serialize
		w.WriteHeader(http.StatusOK)
		w.Write(js)
		return
	}
	if err != sql.ErrNoRows {
		fmt.Println("    Error getting player seasons: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(w, err) // todo, remove when we go to prod
		return
	}
	// Goalie Time, player id returned nothing... maybe its a goalie.
	gs, err := logic.GetGoalieSeasons(playerId)
	if err == nil {
		fmt.Println("    Season request for goalie" + playerStr)
		js, err := json.Marshal(*gs)
		if err != nil {
			panic(err)
		} //Panic, couldn't serialize
		w.WriteHeader(http.StatusOK)
		w.Write(js)
		return
	}
	if err == sql.ErrNoRows {
		http.NotFound(w, r)
		return
	}
	w.WriteHeader(http.StatusInternalServerError)
	fmt.Println(w, err) // todo, remove when we go to prod
}
func PlayerSeasonFull(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	playerStr := vars["playerId"]
	playerId, err := strconv.Atoi(playerStr)
	if err != nil {
		fmt.Println(w, "Bad Arguments")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// TODO, cleanup this garbage code
	ps, err := logic.GetPlayerSeasonsFull(playerId)
	fmt.Println("getting playas.")
	if err == nil {
		fmt.Println("    Season request for player: " + playerStr)
		js, err := json.Marshal(*ps)
		if err != nil {
			panic(err)
		} //Panic, couldn't serialize
		w.WriteHeader(http.StatusOK)
		w.Write(js)
		return
	}
	fmt.Println("getting playas again?.")
	if err != sql.ErrNoRows {
		fmt.Println("    Error getting player seasons: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(w, err) // todo, remove when we go to prod
		return
	}
	fmt.Println("getting goalies.")
	// Goalie Time, player id returned nothing... maybe its a goalie.
	gs, err := logic.GetGoalieSeasonsFull(playerId)
	if err == nil {
		fmt.Println("    Season request for goalie: " + playerStr)
		js, err := json.Marshal(*gs)
		if err != nil {
			panic(err)
		} //Panic, couldn't serialize
		w.WriteHeader(http.StatusOK)
		w.Write(js)
		return
	}
	if err == sql.ErrNoRows {
		http.NotFound(w, r)
		return
	}
	w.WriteHeader(http.StatusInternalServerError)
	fmt.Println(w, err) // todo, remove when we go to prod
}
