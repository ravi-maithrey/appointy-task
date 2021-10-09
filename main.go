package appointytask

import (
	"context" // to take care of tieouts and other signals
	"fmt"
	"log"      //to log errors
	"net/http" //routing
	"time"     //to have a timeout measure

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

type instagram struct {
	User     user
	password string
	post     string
}
type user struct {
	userid    int    `json:"userid"`
	username  string `json:"username"`
	useremail string `json:useremail"`
}

func createUser() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
}

func getUser(w http.ResponseWriter, r *http.Request) {

}
func main() {

	mux = http.NewServeMux() //to handle routing

	mux.HandleFunc("/users", createUser)
	mux.HandleFunc(fmt.Sprintf("/users/%d", userid), getUser)

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://ravi:<password>@cluster0.bwqsb.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
	}

	defer client.Disconnect(ctx)
}
