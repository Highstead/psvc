package models

import (
    "time"
)


type Player struct {
    //playerId            int
    PlayerCompute    string  `json:"playerId"`
    FName            string  `json:"firstName"`
    LName            string  `json:"lastName"`
    Pos              string  `json:"pos"`
    ImageId             int  `json:"imageid"`

    DOB           time.Time  `json:"DoB"`
}

type Players []Player
