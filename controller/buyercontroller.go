package controller

import (
	"context"
	"fmt"

	"github.com/VasenthD/auth-crud/models"
	pro "github.com/VasenthD/auth-crud/proto"
)

type Rpc struct {
	pro.UnimplementedBuyerServiceServer
}

func (r *Rpc) CreateBuyer(ctx context.Context, in *pro.Buyersmodel) pro.DBresponse {

	input := models.Buyersmodel{
		Name:        in.Name,
		Mail:        in.Mail,
		Age:         int(in.Age),
		Phonenumber: int(in.Phonenumber),
		Password:    in.Password,
	}
	fmt.Println(input)
	output := pro.DBresponse{
		Name: in.Name,
	}
	return output
}
