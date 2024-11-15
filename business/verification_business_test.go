package business

import (
	"Twilio-Sms-Sender/mock"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSendVerification_Success(t *testing.T) {
	mockService := &mock.MockVerificationService{ShouldFail: false}
	business := NewVerificationBusiness(mockService)

	response, err := business.SendVerification("+905555555555", "sms")

	assert.NoError(t, err)
	assert.Equal(t, "pending", response["status"])
	assert.Equal(t, "+905555555555", response["to"])
	assert.Equal(t, "sms", response["channel"])
}

func TestSendVerification_InvalidPhoneNumber(t *testing.T) {
	mockService := &mock.MockVerificationService{ShouldFail: false}
	business := NewVerificationBusiness(mockService)

	response, err := business.SendVerification("123", "sms")

	assert.Error(t, err)
	assert.Equal(t, "invalid phone number", err.Error()) // Beklenen hata mesajı değiştirildi
	assert.Nil(t, response)
}

func TestCheckVerification_Success(t *testing.T) {
	mockService := &mock.MockVerificationService{
		ShouldFail: false,
		ValidTo:    "+905555555555",
		ValidCode:  "123456",
	}
	business := NewVerificationBusiness(mockService)

	response, err := business.CheckVerification("+905555555555", "123456")

	assert.NoError(t, err)
	assert.Equal(t, "approved", response["status"])
	assert.Equal(t, true, response["valid"])
}

func TestCheckVerification_Failure_BothInvalid(t *testing.T) {
	mockService := &mock.MockVerificationService{
		ShouldFail: false,
		ValidTo:    "+905555555555",
		ValidCode:  "123456",
	}
	business := NewVerificationBusiness(mockService)

	// Yanlış telefon numarası ve kod
	response, err := business.CheckVerification("905555555", "1234")

	assert.Error(t, err)
	assert.Equal(t, "invalid phone number, invalid verification code", err.Error()) // Birleştirilmiş hata mesajı
	assert.Nil(t, response)
}

func TestCheckVerification_Failure_InvalidPhoneNumber(t *testing.T) {
	mockService := &mock.MockVerificationService{
		ShouldFail: false,
		ValidTo:    "+905555555555",
		ValidCode:  "123456",
	}
	business := NewVerificationBusiness(mockService)

	// Yanlış telefon numarası
	response, err := business.CheckVerification("905555555", "123456") // Doğru kod

	assert.Error(t, err)
	assert.Equal(t, "invalid phone number", err.Error())
	assert.Nil(t, response)
}

func TestCheckVerification_Failure_InvalidCode(t *testing.T) {
	mockService := &mock.MockVerificationService{
		ShouldFail: false,
		ValidTo:    "+905555555555",
		ValidCode:  "123456",
	}
	business := NewVerificationBusiness(mockService)

	// Yanlış kod
	response, err := business.CheckVerification("+905555555555", "") // Doğru telefon numarası

	assert.Error(t, err)
	assert.Equal(t, "invalid verification code", err.Error())
	assert.Nil(t, response)
}
