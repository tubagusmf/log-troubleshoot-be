package model

type WhatsAppWebhookRequest struct {
	Sender        string `json:"sender"`
	Message       string `json:"message"`
	QuotedMessage string `json:"quoted_message"`
	GroupName     string `json:"group_name"`
	Timestamp     string `json:"timestamp"`
}
