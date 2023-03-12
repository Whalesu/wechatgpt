package telegram

import (
	"strings"

	"wechatbot/openai"

	log "github.com/sirupsen/logrus"
)

func Handle(msg string, model string) *string {
	requestText := strings.TrimSpace(msg)
	reply, err := openai.Completions(requestText, model)
	if err != nil {
		log.Error(err)
	}
	return reply
}
