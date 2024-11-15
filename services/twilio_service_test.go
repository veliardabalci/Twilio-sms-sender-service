package services

import (
	"Twilio-Sms-Sender/mock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTwilioService_SendVerification_Success(t *testing.T) {
	mockService := &mock.MockVerificationService{
		ShouldFail: false,
	}

	response, err := mockService.SendVerification("+9053333381", "sms")

	assert.NoError(t, err)                         // Hata olmamalı
	assert.Equal(t, "pending", response["status"]) // Yanıt durumunu kontrol et
	assert.Equal(t, "+9053333381", response["to"]) // Telefon numarasını kontrol et
	assert.Equal(t, "sms", response["channel"])    // Kanalı kontrol et
}

func TestTwilioService_CheckVerification_Failure_InvalidCode(t *testing.T) {
	mockService := &mock.MockVerificationService{
		ShouldFail: false,
		ValidTo:    "+9053333381",
		ValidCode:  "123456",
	}

	response, err := mockService.CheckVerification("+9053333381", "654321")

	assert.Error(t, err)                                         // Hata beklenir
	assert.Nil(t, response)                                      // Yanıt boş olmalı
	assert.Equal(t, "invalid phone number or code", err.Error()) // Hata mesajını kontrol et
}
