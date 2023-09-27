package config

import (
	"context"
	"fmt"
	"time"

	"github.com/VasenthD/auth-crud/infos"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func ConnectDatabase() (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	mongoConnection := options.Client().ApplyURI(infos.Connectionstring)
	mongoClient, err := mongo.Connect(ctx, mongoConnection)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	if err := mongoClient.Ping(ctx, readpref.Primary()); err != nil {
		return nil, err
	}
	return mongoClient, nil
}


func Getcollection(dbname string, collectionname string) (*mongo.Collection,){
	client,err := ConnectDatabase()
	if err != nil {
		fmt.Println("connection issues---!!!")
		return nil
	}
	collection:=client.Database(dbname).Collection(collectionname)
	return collection
}