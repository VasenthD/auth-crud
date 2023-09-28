package services

import (
	"context"
	"fmt"

	"github.com/VasenthD/auth-crud/interfaces"
	"github.com/VasenthD/auth-crud/models"
	"go.mongodb.org/mongo-driver/mongo"
)


type Buyerservice struct {
	collection *mongo.Collection
	ctx context.Context
}

func InitBuyerservice(collection  *mongo.Collection, ctx context.Context)(interfaces.Ibuyers){
	return &Buyerservice{collection,ctx}
}

func (s *Buyerservice)CreateBuyer(buyers *models.Buyersmodel)(*models.DBrespone,error){
	res,err:=s.collection.InsertOne(s.ctx,buyers)
	if err != nil{
		fmt.Println("error inserting the data of buyer..!!!🐼🐼🐼")
	}
	fmt.Println("id 📌 : ",res.InsertedID)
	output := models.DBrespone{
		Name: buyers.Name,
	}
	return &output,nil
}