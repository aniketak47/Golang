package main

import "github.com/gin-gonic/gin"

func main() {
	ConnectDatabase()
	
	router := gin.Default()

	//Defining Group Route
	public := router.Group("/api")
	public.POST("/registerUser", CreateUser)

	router.Run(":8080")

}
