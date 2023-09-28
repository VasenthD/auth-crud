package client

import (
	"fmt"
	"net/http"

	pro "github.com/VasenthD/auth-crud/proto"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func main() {
	route := gin.Default()
	con, err := grpc.Dial("locahost:2023", grpc.WithInsecure())
	if err != nil {
		fmt.Println("the error in client  😮‍💨 😮‍💨 😮‍💨")
	}
	clients := pro.NewBuyerServiceClient(con)

	route.POST("/create", func(c *gin.Context) {
		var request pro.Buyersmodel
		if err := c.ShouldBind(&request); err != nil {
			fmt.Println("err is Binding 😮‍💨 😮‍💨 😮‍💨")
			c.JSON(http.StatusBadRequest, gin.H{"spot : ": "err is Binding 😮‍💨 😮‍💨 😮‍💨"})
		}
		response, err := clients.CreateBuyer(c.Request.Context(), &request)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"value": response, "Welcome": "💐"})
	})
}
