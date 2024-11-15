package mock

import (
	"errors"
)

type MockVerificationService struct {
	ShouldFail bool
	ValidTo    string
	ValidCode  string
}

func (m *MockVerificationService) SendVerification(to, channel string) (map[string]interface{}, error) {
	// Hata durumunu simüle et
	if m.ShouldFail {
		return nil, errors.New("failed to send verification")
	}

	// Geçerli girdiler için başarı yanıtı döndür
	return map[string]interface{}{
		"status":      "pending",
		"to":          to,
		"channel":     channel,
		"service_sid": "mock_service_sid",
		"created_at":  "2024-11-15T13:13:15Z",
	}, nil
}

func (m *MockVerificationService) CheckVerification(to, code string) (map[string]interface{}, error) {
	if to != m.ValidTo || code != m.ValidCode {
		return nil, errors.New("invalid phone number or code")
	}

	return map[string]interface{}{
		"status":      "approved",
		"to":          to,
		"valid":       true,
		"channel":     "sms",
		"service_sid": "mock_service_sid",
	}, nil
}
