package appointytask

import (
	"context" // to take care of tieouts and other signals
	"encoding/json"
	"fmt"
	"log"      //to log errors
	"net/http" //routing
	"time"     //to have a timeout measure

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt" //password hashing
)

var client *mongo.Client

type instagram struct {
	User []user
	post string
}
type user struct {
	userid    int    `json:"userid", db:"userid"`
	username  string `json:"username", db:"username"`
	useremail string `json:useremail", db:"useremail"`
	password  string `json:"password", db:"password"`
}

func createUser(w http.ResponseWriter, r *http.Request) {
	details := &user{}
	_ = json.NewDecoder(r.Body).Decode(details)

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(details.password), 8)

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	usersDatabase := client.Database("user")
	usersDetails := usersDatabase.Collection("user")
	userResult, err := usersDetails.InsertOne(ctx, bson.D{
		{Key: "userid", Value: details.userid},
		{Key: "username", Value: details.username},
		{Key: "useremail", Value: details.useremail},
		{Key: "password", Value: hashedPassword},
	},
	)

	json.NewEncoder(w).Encode(userResult)

}

func getUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	details := &user.userid
	_ = json.NewDecoder(r.Body).Decode(details)
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

}
func main() {

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://ravi:<password>@cluster0.bwqsb.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
	}

	var mux = http.NewServeMux() //to handle routing

	mux.HandleFunc("/users", createUser)
	mux.HandleFunc(fmt.Sprintf("/users/%d", userid), getUser)

}
