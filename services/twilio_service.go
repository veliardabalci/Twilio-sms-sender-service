package services

import (
	"encoding/json"
	"github.com/go-resty/resty/v2"
)

type VerificationService interface {
	SendVerification(to, channel string) (map[string]interface{}, error)
	CheckVerification(to, code string) (map[string]interface{}, error)
}

type TwilioService struct {
	AccountSid string
	AuthToken  string
	ServiceSid string
	BaseURL    string
}

func NewTwilioService(accountSid, authToken, serviceSid string) VerificationService {
	return &TwilioService{
		AccountSid: accountSid,
		AuthToken:  authToken,
		ServiceSid: serviceSid,
		BaseURL:    "https://verify.twilio.com/v2/Services/",
	}
}

func (t *TwilioService) SendVerification(to, channel string) (map[string]interface{}, error) {
	client := resty.New()

	resp, err := client.R().
		SetBasicAuth(t.AccountSid, t.AuthToken).
		SetFormData(map[string]string{
			"To":      to,
			"Channel": channel,
		}).
		Post(t.BaseURL + t.ServiceSid + "/Verifications")

	if err != nil {
		return nil, err
	}

	var parsedResponse map[string]interface{}
	if err = json.Unmarshal([]byte(resp.String()), &parsedResponse); err != nil {
		return nil, err
	}

	return parsedResponse, nil
}

func (t *TwilioService) CheckVerification(to, code string) (map[string]interface{}, error) {
	client := resty.New()

	resp, err := client.R().
		SetBasicAuth(t.AccountSid, t.AuthToken).
		SetFormData(map[string]string{
			"To":   to,
			"Code": code,
		}).
		Post(t.BaseURL + t.ServiceSid + "/VerificationCheck")

	if err != nil {
		return nil, err
	}

	var parsedResponse map[string]interface{}
	if err = json.Unmarshal([]byte(resp.String()), &parsedResponse); err != nil {
		return nil, err
	}

	return parsedResponse, nil
}
