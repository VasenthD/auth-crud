package controller

import (
	"context"
	"fmt"

	"github.com/VasenthD/auth-crud/models"
	pro "github.com/VasenthD/auth-crud/proto"
	"github.com/VasenthD/auth-crud/services"
)

type Rpc struct {
	pro.UnimplementedBuyerServiceServer
}

var (
	buycontroller services.Buyerservice	
)

func (r *Rpc) CreateBuyer(ctx context.Context, in *pro.Buyersmodel) pro.DBresponse {

	input := models.Buyersmodel{
		Name:        in.Name,
		Mail:        in.Mail,
		Age:         int(in.Age),
		Phonenumber: int(in.Phonenumber),
		Password:    in.Password,
	}
	
	result := buycontroller.CreateBuyer(&input)
	output := pro.DBresponse{
		Name: in.Name,
	}
	fmt.Println(result.Name)
	return output
}
