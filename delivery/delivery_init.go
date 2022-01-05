package delivery

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-mongo/db"
	"go-mongo/utils"
	"log"
	"os"
)

type Routes struct {
}

func (app Routes) StartGin() {
	host := os.Getenv("API_HOST")
	port := os.Getenv("API_PORT")

	mbd, err := db.InitResource()
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := utils.InitContext()
	defer cancel()
	defer func() {
		if err = mbd.DB.Client().Disconnect(ctx); err != nil {
			log.Fatal(err)
			return
		}
	}()
	r := gin.Default()
	publicRoute := r.Group("/api")
	NewStudentApi(publicRoute, mbd)

	apiBaseUrl := fmt.Sprintf("%s:%s", host, port)
	err = r.Run(apiBaseUrl)
	if err != nil {
		log.Fatal(err)
	}
}
