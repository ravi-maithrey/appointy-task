package appointytask

import (
	"context" // to take care of tieouts and other signals
	"log"     //to log errors
	"net/http"
	"time" //to have a timeout measure

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type instagram struct {
	user     string
	password string
	post     string
}

func createUser() {

}
func main() {

	mux = http.NewServeMux() //to handle routing

	mux.HandleFunc("/users", createUser)
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
