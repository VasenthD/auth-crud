package controller

import (
	"context"
	"fmt"

	"github.com/VasenthD/auth-crud/interfaces"
	"github.com/VasenthD/auth-crud/models"
	pro "github.com/VasenthD/auth-crud/proto"
)

type Rpc struct {
	pro.UnimplementedBuyerServiceServer
}

var (
	Buycontroller interfaces.Ibuyers
)

func (r *Rpc) CreateBuyer(ctx context.Context, in *pro.Buyersmodel) (*pro.DBresponse, error) {
	input := models.Buyersmodel{
		Name:        in.Name,
		Mail:        in.Mail,
		Age:         int(in.Age),
		Phonenumber: int(in.Phonenumber),
		Password:    in.Password,
	}

	result, err := Buycontroller.CreateBuyer(&input)
	if err != nil {
		fmt.Println("error in controllers 🪼🪼🪼")
	}
	output := pro.DBresponse{
		Name: in.Name,
	}
	fmt.Println(result.Name)
	return &output, nil
}
