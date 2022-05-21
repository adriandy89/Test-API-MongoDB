package models

import "time"

type Word struct {
	Word string `bson:"word" json:"word"`
}

type At struct {
	At *time.Time `bson:"at" json:"at"`
}
