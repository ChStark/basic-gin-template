package middlewares

import (
	"github.com/gin-gonic/gin"
	"time"
)

//Here we can do whatever we cant with the context before and even after the request
func AuthChecker( context *gin.Context) {
	//Validations of the context
	println( "something before auth checker" )
	context.Next()
	println( "something after auth checker" )
}

func AnotherMiddleware( context *gin.Context) {
	println( "something before another middleware" )
	go func() {
		println( "starting side process that doesnt block the request" ) //this log can happen before or after the request ending, is up to go's goroutine management
		time.Sleep( time.Second )
		println( "ending side process that doesnt block the request" )
	}() //this is how you spawn a goroutine that doesn't block the goroutine that is handling the request
	context.Next()
	go func() {
		println( "starting another side process that doesnt block the request" )  //this log can happen before or after the request ending, is up to go's goroutine management
		time.Sleep( time.Second )
		println( "ending another side process that doesnt block the request" )
	}() //this is how you spawn a goroutine that doesn't block the goroutine that is handling the request
	println( "something after another middleware" )
}