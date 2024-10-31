package main

import (
	"database/sql"
	"reUnionBe/database"
	"reUnionBe/memberRes"
	"reUnionBe/scoringEp"
	"reUnionBe/user"

	"github.com/gin-gonic/gin"
)
var Db *sql.DB


func main() {
  
	route:=gin.Default()
	database.ConnectDatabase()

	//router setup 
	//users
	route.POST("/addUser", user.AddUser)
	route.GET("/getUser",user.GetUser)
	//member-res
	route.GET("/getUserRes",memberRes.GetUserRes)
	route.POST("/addUserRes",memberRes.AddMemberRes)
	//scoring-ep
	route.GET("getScoringEp",scoringEp.GetScoringEp)
	route.POST("addScoringEp",scoringEp.AddScoringEp)
	
	err:=route.Run(":8080")
	if err!= nil{
		panic(err)
	}
}
