package mgo

import (
	"context"
	"errors"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// DB structure
type DB struct {
	Name         string
	Database     *mongo.Database
	Client       *mongo.Client
	Context      context.Context
	QueryTimeout time.Duration
}

// parameters structure
type parameters struct {
	host           string
	port           int
	username       string
	password       string
	name           string
	maxPoolSize    uint64
	connectTimeout time.Duration
	socketTimeout  time.Duration
	queryTimeout   time.Duration
}

// DBInstance is an *DB instance creates by Open().
var DBInstance *DB

// Open initializes a MongoDB database.
func Open(host string,
	port int,
	username, password, name string,
	maxPoolSize uint64,
	connectTimeout, socketTimeout, queryTimeout time.Duration) error {

	// Paramètres
	// ----------
	params := parameters{
		host:           host,
		port:           port,
		username:       username,
		password:       password,
		name:           name,
		maxPoolSize:    maxPoolSize,
		connectTimeout: connectTimeout,
		socketTimeout:  socketTimeout,
		queryTimeout:   queryTimeout,
	}
	err := checkParameters(&params)
	if err != nil {
		return err
	}

	uri := "mongodb://" +
		params.username + ":" + params.password +
		"@" + params.host + ":" + strconv.Itoa(params.port) +
		"/" + params.name

	ctx := context.Background() // or TODO() ?
	client, err := mongo.Connect(ctx,
		options.Client().ApplyURI(uri),
		options.Client().SetMaxPoolSize(params.maxPoolSize),
		options.Client().SetConnectTimeout(params.connectTimeout*time.Second),
		options.Client().SetSocketTimeout(params.socketTimeout*time.Second))
	if err != nil {
		return err
	}

	DBInstance = &DB{
		Name:         params.name,
		Database:     client.Database(params.name),
		Client:       client,
		Context:      ctx,
		QueryTimeout: params.queryTimeout * time.Second,
	}

	return nil
}

// checkParameters checks if database parameters are correct.
func checkParameters(p *parameters) error {
	if p.host == "" {
		return errors.New("mongoDB host is empty")
	}

	if p.port == 0 {
		return errors.New("mongoDB port is empty")
	}

	if p.username == "" {
		return errors.New("mongoDB username is empty")
	}

	if p.password == "" {
		return errors.New("mongoDB password is empty")
	}

	if p.name == "" {
		return errors.New("mongoDB database name is empty")
	}

	if p.maxPoolSize <= 0 {
		return errors.New("mongoDB max pool size must be greater than 0")
	}

	// Valeurs par défaut si incorrect
	// -------------------------------
	if p.connectTimeout <= 0 {
		p.connectTimeout = 10 * time.Second
	}

	if p.socketTimeout <= 0 {
		p.socketTimeout = 10 * time.Second
	}

	if p.queryTimeout <= 0 {
		p.queryTimeout = 10 * time.Second
	}

	return nil
}
