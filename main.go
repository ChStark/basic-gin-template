package main

import (
	"gin-template/handlers"
	"github.com/gin-gonic/gin"
)

func main(){
	r := gin.Default()

	handlers.User( r )
	handlers.Accounts( r )

	_ = r.Run("0.0.0.0:8080")
}
