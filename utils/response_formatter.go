package utils

func FormatSendVerificationResponse(parsedResponse map[string]interface{}) map[string]interface{} {
	return map[string]interface{}{
		"status":      parsedResponse["status"],
		"to":          parsedResponse["to"],
		"channel":     parsedResponse["channel"],
		"service_sid": parsedResponse["service_sid"],
		"created_at":  parsedResponse["date_created"],
	}
}

func FormatCheckVerificationResponse(parsedResponse map[string]interface{}) map[string]interface{} {
	return map[string]interface{}{
		"status":      parsedResponse["status"],
		"to":          parsedResponse["to"],
		"channel":     parsedResponse["channel"],
		"service_sid": parsedResponse["service_sid"],
		"created_at":  parsedResponse["date_created"],
		"updated_at":  parsedResponse["date_updated"],
		"valid":       parsedResponse["valid"],
	}
}
