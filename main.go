package main

import (
	"GinTest/common"
	"GinTest/controller"
	"GinTest/middleware"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	InitConfig()
	common.InitDataBase()
	engine := gin.Default()
	CollectRoutes(engine)
	port := viper.GetString("server.port")
	if port != "" {
		panic(engine.Run(":" + port))
	}
	panic(engine.Run())
}

func CollectRoutes(r *gin.Engine) *gin.Engine {
	r.POST("/register", controller.Register)
	r.POST("/login", controller.Login)
	r.POST("/createquestion", controller.CreateQuestion)
	r.POST("/deletequestion", controller.DeleteQuestion)
	r.POST("/updatequestion", controller.UpdateQuestion)
	r.POST("/uploadexcel", controller.UploadExcel)
	r.POST("/createanswersheet", controller.CreateAnswerSheet)
	r.GET("/importstudents", controller.ImportStudents)
	r.GET("/userinfo", middleware.AuthMiddleware(), controller.UserInfo)
	return r
}

func InitConfig() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/config")
	viper.ReadInConfig()
}
