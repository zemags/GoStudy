package main

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// mongod --dbpath db
	dsn := "mongodb://localhost:27017/"
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dsn))
	if err != nil {
		log.Fatalln("cannot connect to db: ", err)
	}
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			log.Fatalln("connection accidentally closed: ", err)
		}
	}()
	if err := client.Ping(ctx, nil); err != nil {
		log.Fatalln(err)
	}

	collection := client.Database("universe").Collection("stars")
	errInsert := createStar(&Star{
		ID:        primitive.NewObjectID(),
		StartType: "g2v",
		StarName:  "sun",
		Distance:  0.0000158,
	}, collection, &ctx)
	if errInsert != nil {
		log.Fatalln("cannot insert to db: ", errInsert)
	}
}

func createStar(star *Star, coll *mongo.Collection, ctx *context.Context) error {
	_, err := coll.InsertOne(*ctx, star)
	return err
}
