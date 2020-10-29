package handlers

import "github.com/gin-gonic/gin"

func Accounts( router *gin.Engine) {
	router.GET("/account/:id", func(context *gin.Context) {
		context.JSON( 200 , gin.H{ "id": context.Param("id") })
	})
}