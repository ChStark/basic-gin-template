package handlers

import (
	"gin-template/middlewares"
	"github.com/gin-gonic/gin"
)

func User(router *gin.Engine) {

	group := router.Group("/user")

	group.Use( middlewares.AuthChecker )
	group.Use( middlewares.AnotherMiddleware )

	//the full path is GET /user/profile since it prepends the group path
	group.GET("/profile", func(context *gin.Context) {
		context.JSON( 200 , gin.H{ "profile": "some profile" })
	})

	addressesHandler := func(context *gin.Context) {
		context.JSON( 200 , gin.H{ "profile": "some profile" })
	}

	//same handler for different paths
	group.GET("/addresses" , addressesHandler )
	group.GET("/my-addresses" , addressesHandler )

}
