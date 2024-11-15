package router

import (
	"Twilio-Sms-Sender/business"
	"Twilio-Sms-Sender/controllers"
	"Twilio-Sms-Sender/services"
	"Twilio-Sms-Sender/utils"

	"github.com/gin-gonic/gin"
)

func SetupRouter(server *gin.Engine) {
	utils.LoadEnv()

	twilioService := services.NewTwilioService(
		utils.GetEnv("TWILIO_ACCOUNT_SID"),
		utils.GetEnv("TWILIO_AUTH_TOKEN"),
		utils.GetEnv("TWILIO_SERVICE_SID"),
	)

	verificationBusiness := business.NewVerificationBusiness(twilioService)
	verificationController := controllers.NewVerificationController(verificationBusiness)

	api := server.Group("/api")
	{
		SetupTwilioRoutes(api, verificationController)
	}
}
