package main

import "go.mongodb.org/mongo-driver/bson/primitive"

type Star struct {
	ID        primitive.ObjectID `bson:"_id"`
	StartType string             `bson:"star_type"`
	StarName  string             `bson:"star_name"`
	Distance  float64            `bson:"distance"`
}
