package appointytask

import (
	"context"       // to take care of tieouts and other signals
	"encoding/json" //to decode and encode incoming and outgoing data
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

type instagram struct { //structure of the total data
	User []user
	Post []post
}
type user struct { //structure of user data
	userid    int    `json:"userid", bson:"userid"`
	username  string `json:"username", bson:"username"`
	useremail string `json:useremail", bson:"useremail"`
	password  string `json:"password", bson:"password"`
}
type post struct { //structure of posts data
	postid      int       `json:"postid", bson:"postid"`
	postcaption string    `json:"postcaption", bson:"postcaption"`
	imageurl    string    `json:"imageurl", bson:"imageurl"`
	timestamp   time.Time `json:"timestamp", bson:"timestamp"`
}

func createUser(w http.ResponseWriter, r *http.Request) {
	details := &user{}
	_ = json.NewDecoder(r.Body).Decode(details)
	//hashing to store passwords
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
	var details int
	_ = json.NewDecoder(r.Body).Decode(details)

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err := client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	usersDatabase := client.Database("users")
	usersDetails := usersDatabase.Collection("userdetails")
	userResult, err := usersDetails.Find(ctx, bson.M{"userid": details})
	json.NewEncoder(w).Encode(userResult)
}

func createPost(w http.ResponseWriter, r *http.Request) {
	details := &post{}
	_ = json.NewDecoder(r.Body).Decode(details)

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err := client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	postsDatabase := client.Database("posts")
	postDetails := postsDatabase.Collection("postdetails")
	postResult, err := postDetails.InsertOne(ctx, bson.D{
		{Key: "postid", Value: details.postid},
		{Key: "postcaption", Value: details.postcaption},
		{Key: "imageurl", Value: details.imageurl},
		{Key: "timestamp", Value: details.timestamp},
	},
	)

	json.NewEncoder(w).Encode(postResult)
}

func getPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	var details int
	_ = json.NewDecoder(r.Body).Decode(details)

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err := client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	postsDatabase := client.Database("posts")
	postDetails := postsDatabase.Collection("postdetails")
	postResult, err := postDetails.InsertOne(ctx, bson.M{"postid": details})
	json.NewEncoder(w).Encode(postResult)
}

func getAllPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	var details int
	_ = json.NewDecoder(r.Body).Decode(details)

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err := client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	postsDatabase := client.Database("posts")
	postDetails := postsDatabase.Collection("postdetails")
	postResult, err := postDetails.InsertOne(ctx, bson.M{"postid": details})
	json.NewEncoder(w).Encode(postResult)
}

func routing(r *http.Request) {
	details := &instagram{}
	_ = json.NewDecoder(r.Body).Decode(details)
	mux := http.NewServeMux() //to handle routing

	mux.HandleFunc("/users", createUser)
	mux.HandleFunc(fmt.Sprintf("/users/%d", details.User[0].userid), getUser)
	mux.HandleFunc("/posts", createPost)
	mux.HandleFunc(fmt.Sprintf("/posts/%d", details.Post[0].postid), getPost)
	mux.HandleFunc(fmt.Sprintf("/posts/users/%d", details.User[0].userid), getAllPosts)
}
func main() {
	var url http.Request
	_, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://ravi:ravi_password@cluster0.bwqsb.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Scanln(&url)
	routing(&url)

}
