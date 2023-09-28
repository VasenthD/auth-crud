package interfaces

import "github.com/VasenthD/auth-crud/models"

type Ibuyers interface{
	CreateBuyer(input *models.Buyersmodel)(*models.DBrespone,error)
}