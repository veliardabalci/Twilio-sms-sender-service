package business

import (
	"Twilio-Sms-Sender/services"
	"Twilio-Sms-Sender/utils"
	"errors"
	"strings"
)

type VerificationBusiness struct {
	Service services.VerificationService
}

func NewVerificationBusiness(service services.VerificationService) *VerificationBusiness {
	return &VerificationBusiness{
		Service: service,
	}
}

func (vb *VerificationBusiness) SendVerification(to, channel string) (map[string]interface{}, error) {
	if to == "" || len(to) < 10 {
		return nil, errors.New("invalid phone number")
	}
	if channel != "sms" && channel != "email" {
		return nil, errors.New("invalid channel")
	}

	response, err := vb.Service.SendVerification(to, channel)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (vb *VerificationBusiness) CheckVerification(to, code string) (map[string]interface{}, error) {
	var validationErrors []string

	if len(to) < 10 || len(to) > 15 || !strings.HasPrefix(to, "+") {
		validationErrors = append(validationErrors, "invalid phone number")
	}

	if len(code) != 6 || !utils.IsNumeric(code) {
		validationErrors = append(validationErrors, "invalid verification code")
	}

	if len(validationErrors) > 0 {
		return nil, errors.New(strings.Join(validationErrors, ", "))
	}

	response, err := vb.Service.CheckVerification(to, code)
	if err != nil {
		return nil, err
	}

	return response, nil
}
