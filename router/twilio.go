package router

import (
	"Twilio-Sms-Sender/controllers"

	"github.com/gin-gonic/gin"
)

func SetupTwilioRoutes(router *gin.RouterGroup, controller *controllers.VerificationController) {
	twilio := router.Group("/twilio")
	{
		twilio.POST("/send-verification", controller.SendVerification)
		twilio.POST("/check-verification", controller.CheckVerification)
	}
}
