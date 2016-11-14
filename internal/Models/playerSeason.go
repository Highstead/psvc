package models

import (
    "time"
)


// SELECT (fn).first_name, (fn).last_name, (fn).pos, (fn).image_id, (fn).Age, (fn).End_Team, (fn).GP, (fn).goals, (fn).points, (fn).assists, (fn).PPG, (fn).ESTOI, (fn).Corsi from temp;
type PlayerSeason struct {
    //playerId            int
    PlayerCompute    string  `json:"playerId"`
    FName            string  `json:"firstName"`
    LName            string  `json:"lastName"`
    Pos              string  `json:"pos"`
    ImageId             int  `json:"imageid"`
    DOB           time.Time  `json:"DoB"`

    EndTeam          string  `json:"endTeam"`
    GP                  int  `json:"GP"`
    Goals               int  `json:"goals"`
    Points          float32  `json:"points"`
    Assits              int  `json:"assists"`
    PPG             float32  `json:"PPG"`
    ESTOI           float32  `json:"ESTOI"`
    Corsi           float32  `json:"corsi"`
}
