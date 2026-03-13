package petitions

import (
	"log"

	"github.com/gin-gonic/gin"
)

func Register(g *gin.Engine) {
	v1 := g.Group("/api/v1")

	petitionsGroup := v1.Group("/petitions")
	petitionsGroup.GET("", petitionsGET)
	petitionsGroup.POST("", petitionsPOST)
}

func petitionsGET(c *gin.Context) {
	GetPetitionController().GetPetitions()
	c.JSON(200, gin.H{
		"status": "ok",
	})
}

func petitionsPOST(c *gin.Context) {
	var request CreatePetitionRequest
	c.BindJSON(&request)
	resp, err := GetPetitionController().CreatePetition(request.Name, request.Description)
	if err != nil {
		c.JSON(500, CreatePetitionResponse{
			Created: false,
		})
	}
	log.Println(resp)
	c.JSON(200, CreatePetitionResponse{
		Created: true,
	})
}
