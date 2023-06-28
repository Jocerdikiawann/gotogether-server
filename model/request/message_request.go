package request

type MessageTelegram struct {
	Text                  string `json:"text"`
	DisableWebPagePreview bool   `json:"disable_web_page_preview"`
	DisableNotification   bool   `json:"disable_notification"`
	ReplyToMessageID      int64  `json:"reply_to_message_id,omitempty"`
	ChatID                int64  `json:"chat_id"`
}
