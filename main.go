package main

import (
	"context" // to take care of tieouts and other signals
	"log"     //to log errors
	"time"    //to have a timeout measure

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func createUser()

func main() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://ravi:<password>@cluster0.bwqsb.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)
}
