package logic

import (
    "app/ddd/internal/models"
	"database/sql"
    "fmt"
    "time"

	_ "github.com/lib/pq"

)

const timeConst = "2006-01-02"
const fnGetPlayers = "WITH temp as (SELECT fnGetPlayers() AS fn)" +
    "SELECT (fn).player_id, (fn).first_name, (fn).last_name, (fn).pos, (fn).image_id from temp;"

/*
p.first_name,
      p.last_name,
      p.pos,
      p.image_id,
      DATE '1984-02-12' as Age,   -- todo: replace with age
      s.End_Team,
      s.GP,
      s.ESG as goals,    -- May have to be Sum( verify with nhl.com)
      s.ESP as points,
      s.ESA as assists,
      (s.ESP::Decimal(5,3) / s.GP) as PPG,
      s.ESTOI,
      s.Corsi,
      s.Adj_Corsi
 */

const fnGetPlayerSeason = "WITH temp as (SELECT fnGetPlayerSeason($1) AS fn)" +
    "SELECT (fn).first_name, (fn).last_name, (fn).pos, (fn).image_id, (fn).Age, (fn).End_Team, (fn).GP, (fn).goals," +
    "(fn).points, (fn).assists, (fn).PPG, (fn).ESTOI, (fn).Corsi from temp;"


func GetPlayers() []models.Player {
    fmt.Println("getting players, no ssl")
    hasRows := false
	// db, err := sql.Open("postgres", "host=192.168.99.100 user=postgres dbname=player sslmode=verify-full")
     db, err := sql.Open("postgres", "host=192.168.99.100 user=postgres dbname=player sslmode=disable")
	if err != nil {
		fmt.Println(err)
	}

	// we do it this way until the potsgres drivers return more than 1 parameter.
    rows, err := db.Query(fnGetPlayers);

    if err != nil {
        fmt.Println(err)
        return nil
    }
    //rows.Close()  // just added
    players := []models.Player {}
    DOBdate, _ := time.Parse(timeConst, "1988-02-12")

    for rows.Next() {
        hasRows = true

        var player_id int
        var f_name string
        var l_name string
        var pos string
        var image_id int

        rowerr := rows.Scan(&player_id, &f_name, &l_name, &pos, &image_id)
        if rowerr != nil {
            fmt.Println(rowerr)
            continue
        }
        players = append(players, models.Player{FName: f_name, LName: l_name, Pos: pos, ImageId: 0, PlayerCompute: f_name, DOB: DOBdate})
    }
    if !hasRows {
        fmt.Println("No Rows")
        return nil
    }

    return players
}

func GetPlayerSeasons(player_id int) (models.PlayerSeason, error) {
    fmt.Println("getting seasons, no ssl")
    var ps models.PlayerSeason

    db, err := sql.Open("postgres", "host=192.168.99.100 user=postgres dbname=player sslmode=disable")
	if err != nil {
		fmt.Println(err)
        return ps, err
	}

    if err != nil {
        fmt.Println(err)
        return ps, err
    }
    //rows.Close()  // just added

    // we do it this way until the potsgres drivers return more than 1 parameter.
    queryErr := db.QueryRow(fnGetPlayerSeason, player_id).Scan(&(ps.FName), &(ps.LName), &(ps.Pos), &(ps.ImageId), &(ps.DOB),
        &(ps.EndTeam), &(ps.GP), &(ps.Goals), &(ps.Points), &(ps.Assits), &(ps.PPG), &(ps.ESTOI), &(ps.Corsi))

    //rowerr := rows.Scan(&player_id, &f_name, &l_name, &pos, &image_id)
    if queryErr != nil {
        fmt.Println(queryErr)
        return ps, queryErr
    }
    return ps, err
}