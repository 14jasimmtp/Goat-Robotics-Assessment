package main

import (
	"github.com/14jasimmtp/Goat-Robotics-Assessment/pkg/di"
	"github.com/14jasimmtp/Goat-Robotics-Assessment/pkg/routes"
	"github.com/gin-gonic/gin"
)

func main(){
	AuthHandler:=di.InjectDependencies()
	r:=gin.Default()
	routes.Routes(r,AuthHandler)

	r.Run(":3000")
}