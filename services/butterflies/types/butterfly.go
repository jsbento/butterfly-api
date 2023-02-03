package types

import (
	"time"
)

type Butterfly struct {
	Id             string    `bson:"_id" json:"id"`
	ScientificName string    `bson:"scientificName" json:"scientificName"`
	CommonName     string    `bson:"commonName" json:"commonName"`
	Image          string    `bson:"image" json:"image"`
	Family         string    `bson:"family" json:"family"`
	Subfamily      string    `bson:"subfamily" json:"subfamily"`
	LifeSpan       string    `bson:"lifeSpan" json:"lifeSpan"`
	Range          string    `bson:"range" json:"range"`
	Hosts          string    `bson:"hosts" json:"hosts"`
	Food           string    `bson:"food" json:"food"`
	Habitat        string    `bson:"habitat" json:"habitat"`
	Flights        string    `bson:"flights" json:"flights"`
	History        string    `bson:"history" json:"history"`
	FunFact        string    `bson:"funFact" json:"funFact"`
	Etymology      string    `bson:"etymology" json:"etymology"`
	CreatedAt      time.Time `bson:"createdAt" json:"createdAt"`
	UpdatedAt      time.Time `bson:"updatedAt" json:"updatedAt"`
}
