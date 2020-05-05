package mgo

import (
	"context"
	"errors"
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestCheckParameters(t *testing.T) {
	tests := []struct {
		p      *parameters
		wanted error
	}{
		{&parameters{
			"host",
			10,
			"username",
			"password",
			"name",
			uint64(50),
			2 * time.Second,
			2 * time.Second,
			0 * time.Second}, nil},
		{&parameters{
			"",
			10,
			"username",
			"password",
			"name",
			uint64(50),
			2 * time.Second,
			2 * time.Second,
			0 * time.Second}, errors.New("mongoDB host is empty")},
		{&parameters{
			"host",
			10,
			"",
			"password",
			"name",
			uint64(50),
			2 * time.Second,
			2 * time.Second,
			0 * time.Second}, errors.New("mongoDB username is empty")},
		{&parameters{
			"host",
			10,
			"username",
			"",
			"name",
			uint64(50),
			2 * time.Second,
			2 * time.Second,
			0 * time.Second}, errors.New("mongoDB password is empty")},
		{&parameters{
			"host",
			10,
			"username",
			"password",
			"",
			uint64(50),
			2 * time.Second,
			2 * time.Second,
			0 * time.Second}, errors.New("mongoDB database name is empty")},
		{&parameters{
			"host",
			10,
			"username",
			"password",
			"name",
			0,
			2 * time.Second,
			2 * time.Second,
			0 * time.Second}, errors.New("mongoDB max pool size must be greater than 0")},
	}

	for _, test := range tests {
		assert.Equal(t, checkParameters(test.p), test.wanted)
	}
}

func TestCheckParametersChangeTimeouts(t *testing.T) {
	p := &parameters{
		"host",
		10,
		"username",
		"password",
		"name",
		uint64(50),
		0 * time.Second,
		0 * time.Second,
		0 * time.Second,
	}
	wanted := &parameters{
		"host",
		10,
		"username",
		"password",
		"name",
		uint64(50),
		10 * time.Second,
		10 * time.Second,
		10 * time.Second,
	}
	checkParameters(p)
	assert.Equal(t, p, wanted)
}

func ExampleOpen() {
	// Database connection
	// -------------------
	err := Open("host", 10, "username", "password", "name", uint64(50), 10*time.Second, 10*time.Second, 10*time.Second)
	if err != nil {
		log.Fatalf("Database connection failed: %v\n", err)
	}
	defer func() {
		if err = DBInstance.Client.Disconnect(DBInstance.Context); err != nil {
			log.Fatalf("Database disconnection failed: %v\n", err)
		}
	}()

	// Usage
	// -----
	type User struct {
		ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
		Lastname  string             `json:"lastname"`
		Firstname string             `json:"firstname"`
		Email     string             `json:"email"`
		Password  string             `json:"password,omitempty"`
		CreatedAt time.Time          `json:"createdAt" bson:"created_at"`
	}

	users := make([]User, 0)
	col := DBInstance.Database.Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), DBInstance.QueryTimeout)
	options := options.Find().SetSort(bson.D{primitive.E{Key: "created_at", Value: -1}})
	filter := bson.D{}
	cur, err := col.Find(ctx, filter, options)
	if err != nil {
		cancel()
	}
	defer cur.Close(context.Background())

	for cur.Next(context.Background()) {
		var user User

		err := cur.Decode(&user)
		if err != nil {
			continue
		}

		user.Password = ""

		users = append(users, user)
	}

	cancel()
	fmt.Printf("Users: %v\n", users)
}
