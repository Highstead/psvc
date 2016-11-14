package ddd

import (
    "encoding/json"
    "fmt"
    "strconv"
    "net/http"

    "github.com/gorilla/mux"
    "time"
    "app/ddd/internal/logic"
    "app/ddd/internal/models"
)

func Index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Welcome!")
}

const timeConst string = "2006-01-02"
func PlayerIndex(w http.ResponseWriter, r *http.Request) {
    DOBdate := time.Now()

    DOBdate, _ = time.Parse(timeConst, "1988-02-12")
    players := models.Players{
        models.Player{FName: "Sidney", LName: "Crosby", ImageId: 1999, PlayerCompute: "SCROSBY19880212", DOB: DOBdate},
        models.Player{FName: "Alex", LName: "Ovechkin", ImageId: 1999, PlayerCompute: "AOVECHK19880212", DOB: DOBdate},
    }

    if err := json.NewEncoder(w).Encode(players); err != nil {
        panic(err)
    }

}

func PlayerShow(w http.ResponseWriter, r *http.Request) {
    // vars := mux.Vars(r)
    // playerId := vars["playerId"]

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

    fmt.Println("Season request for player" + playerStr )
    ps, err := logic.GetPlayerSeasons(playerId)

    if err != nil {
        panic(err)
    }
    js, err := json.Marshal(ps)

    if err != nil {
        fmt.Println(w, "Bad Arguments")
        w.WriteHeader(http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    w.Write(js)
    // ps, _ := logic.GetPlayerSeasons(playerId)
    //fmt.Fprintln(w, ps)
}