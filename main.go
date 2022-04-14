package main

import (
	"github.com/gin-gonic/gin"
	"github.com/raycad/go-microservices/tree/master/src/movie-microservice/common"
)

type Main struct {
	router *gin.Engine
}

func (m *Main) initServer() error {
	var err error

	//load config file
	err = common.LoadConfig()
	if err != nil {
		return err
	}
}

func main() {

}
