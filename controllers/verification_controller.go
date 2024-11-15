package controllers

import (
	"Twilio-Sms-Sender/business"
	"Twilio-Sms-Sender/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type VerificationController struct {
	Business *business.VerificationBusiness
}

func NewVerificationController(business *business.VerificationBusiness) *VerificationController {
	return &VerificationController{
		Business: business,
	}
}

func (vc *VerificationController) SendVerification(c *gin.Context) {
	to := c.PostForm("To")
	channel := c.PostForm("Channel")

	response, err := vc.Business.SendVerification(to, channel)
	if err != nil {
		logrus.Errorf("Failed to send verification to %s via %s: %s", to, channel, err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if status, ok := response["status"].(float64); ok && status != 200 {
		logrus.Errorf("Verification failed for %s: %v", to, response)
		c.JSON(http.StatusBadRequest, gin.H{
			"code":      response["code"],
			"message":   response["message"],
			"more_info": response["more_info"],
			"status":    response["status"],
		})
		return
	}
	logrus.Infof("Verification message successfully sent to %s via %s", to, channel)

	formattedResponse := utils.FormatSendVerificationResponse(response)
	c.JSON(http.StatusOK, formattedResponse)
}

func (vc *VerificationController) CheckVerification(c *gin.Context) {
	to := c.PostForm("To")
	code := c.PostForm("Code")

	response, err := vc.Business.CheckVerification(to, code)
	if err != nil {
		logrus.Errorf("Failed to verify code for %s: %s", to, err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "unexpected error occurred"})
		return
	}

	if status, ok := response["status"].(float64); ok && status != 200 {
		logrus.Errorf("Verification failed for %s: %v", to, response)
		c.JSON(http.StatusBadRequest, gin.H{
			"code":      response["code"],
			"message":   response["message"],
			"more_info": response["more_info"],
			"status":    response["status"],
		})
		return
	}

	logrus.Infof("Verification successfully completed for %s", to)
	formattedResponse := utils.FormatCheckVerificationResponse(response)
	c.JSON(http.StatusOK, formattedResponse)
}
