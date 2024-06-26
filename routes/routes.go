package routes

import (
	"mods/controller"
	"mods/middleware"
	"mods/service"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine, userController controller.UserController, diseaseController controller.DiseaseController, predictionController controller.PredictionController, chatroomController  controller.ChatroomController,jwtService service.JWTService) {
	inscurePublic := router.Group("/inscure")
	{
		// public can access
		inscurePublic.POST("/add", userController.AddUser)
		inscurePublic.POST("/login", userController.UserLoginToken)
		inscurePublic.POST("/adddoctor", userController.AddDoctor)
	}

	userPrivate := router.Group("/user").Use(middleware.Authenticate())
	{
		userPrivate.GET("/me", userController.Me)
		userPrivate.GET("", userController.GetAllUser)
		userPrivate.GET("/doctor", userController.GetAllDoctor)
		userPrivate.DELETE("/:id", userController.DeleteUser)
		userPrivate.PUT("/updateName", userController.UpdateUserName)
		userPrivate.PUT("/updateNotelp", userController.UpdateUserNotelp)
		userPrivate.PUT("/profilepic", userController.ProfilePicture)
	}

	diseasePublic := router.Group("/disease")
	{
		diseasePublic.POST("/add", diseaseController.AddDisease)
		diseasePublic.GET("", diseaseController.GetAllDisease)
		diseasePublic.DELETE("/:id", diseaseController.DeleteDisease)
		diseasePublic.GET("/:id", diseaseController.GetDiseaseByID)
	}

	predictionPublic := router.Group("/prediction").Use(middleware.Authenticate())
	{
		predictionPublic.POST("", predictionController.AddPrediction)
		predictionPublic.GET("/list", predictionController.GetPredictionByUserID)
		predictionPublic.GET("/:p_id", predictionController.GetPredictionByPredictionID)
		predictionPublic.DELETE("/del", predictionController.DeletePredictionbyId)
	}

	chatroomPublic := router.Group("/chatroom").Use(middleware.Authenticate())
	{
		chatroomPublic.POST("/add", chatroomController.AddChatroom)
		chatroomPublic.DELETE("/del/:id", chatroomController.RemoveChatroom)
		chatroomPublic.GET("/get", chatroomController.GetChatroom)
	}

}
