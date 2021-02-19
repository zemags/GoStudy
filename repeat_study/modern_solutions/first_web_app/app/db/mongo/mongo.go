package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
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
	deleteStar("sirius", collection, &ctx)
	errInsert := createStar(&Star{
		ID:        primitive.NewObjectID(),
		StartType: "a1vm",
		StarName:  "sirius",
		Distance:  8.6,
	}, collection, &ctx)
	if errInsert != nil {
		log.Fatalln("cannot insert to db: ", errInsert)
	}
	star := readStar("sun", collection, &ctx)
	fmt.Println(star)
}

func createStar(star *Star, coll *mongo.Collection, ctx *context.Context) error {
	_, err := coll.InsertOne(*ctx, star)
	return err
}

func readStar(starName string, coll *mongo.Collection, ctx *context.Context) *Star {
	star := &Star{}
	filter := bson.D{{"star_name", starName}}
	err := coll.FindOne(*ctx, filter).Decode(star)
	if err != nil {
		log.Fatalln(err)
	}
	return star
}

func deleteStar(starName string, coll *mongo.Collection, ctx *context.Context) {
	filter := bson.D{{"star_name", starName}}
	_, err := coll.DeleteOne(*ctx, filter)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("start %s deleted\n", starName)
}
