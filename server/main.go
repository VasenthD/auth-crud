package main

import (
	"context"
	"fmt"
	"net"
	pro "github.com/VasenthD/auth-crud/proto"
	"github.com/VasenthD/auth-crud/config"
	"github.com/VasenthD/auth-crud/controller"
	"github.com/VasenthD/auth-crud/infos"
	"github.com/VasenthD/auth-crud/services"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
)

func initserver(mongoclient *mongo.Client) {
	collection := config.Getcollection(infos.Dbname, "authcrud")
	controller.Buycontroller = services.InitBuyerservice(collection,context.Background())
}

func main() {
	mongclient, err := config.ConnectDatabase()
	if err != nil {
		fmt.Println("connection problem")
		return 
	}
	initserver(mongclient)
	lis, err := net.Listen("tcp", ":3000")
	if err != nil {
		fmt.Println("listening errosscccr :",err)
		return
	}

	s:=grpc.NewServer()
	pro.RegisterBuyerServiceServer(s, &controller.Rpc{})
	fmt.Println("Server listening on", infos.Port)
	fmt.Println("🥳 🥳 🥳")
	if err := s.Serve(lis); err != nil {
		fmt.Printf("Failed to serve: %v", err)
	}


}
