package model

import (
	"time"
)

type Fortune struct {
    Ganji      string    `json:"ganji" bson:"ganji" required:"true"`
    Fortune    string    `json:"fortune" bson:"fortune" required:"true"`
    CreatedAt  time.Time `json:"createdAt" bson:"createdAt" default:"$currentDate"`
}
