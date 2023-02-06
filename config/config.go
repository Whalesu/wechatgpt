package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"strings"
	"strconv"
)

var config *Config

type Config struct {
	ChatGpt ChatGptConfig `json:"chatgpt"`
}

type ChatGptConfig struct {
	Token         string  `json:"token,omitempty"`
	Wechat        *string `json:"wechat,omitempty"`
	WechatKeyword *string `json:"wechatkeyword"`
	Model        *string `json:"model,omitempty"`
	MaxLen        *int `json:"maxlen,omitempty"`
	Telegram      *string `json:"telegram"`
	TgWhitelist   *string `json:"tgWhitelist"`
	TgKeyword     *string `json:"tgKeyword"`
}

func LoadConfig() error {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./local")
	viper.AddConfigPath("./config")

	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	if err := viper.Unmarshal(&config); err != nil {
		return err
	}
	return nil
}

func GetWechat() *string {
	wechat := getEnv("wechat")

	if wechat != nil {
		return wechat
	}
	if config == nil {
		return nil
	}
	if wechat == nil {
		wechat = config.ChatGpt.Wechat
	}
	return wechat
}

func GetWechatKeyword() *string {
	keyword := getEnv("wechatkeyword")

	if keyword != nil {
		return keyword
	}
	if config == nil {
		return nil
	}
	if keyword == nil {
		keyword = config.ChatGpt.WechatKeyword
	}
	return keyword
}

func GetModelType() *string {
	keyword := getEnv("Model")

	if keyword != nil {
		return keyword
	}
	if config == nil {
		return nil
	}
	if keyword == nil {
		keyword = config.ChatGpt.Model
	}
	return keyword
}

func GetMaxLen() *int {
	keyword := getEnv("MaxLen")

	if keyword != nil {
		maxlen, _ := strconv.Atoi(*keyword)
		return &maxlen
	}
	if config == nil {
		return nil
	}
	if keyword == nil {
		return config.ChatGpt.MaxLen
	}
	return nil
}

func GetTelegram() *string {
	tg := getEnv("telegram")
	fmt.Println(tg)
	if tg != nil {
		return tg
	}
	if config == nil {
		return nil
	}
	if tg == nil {
		tg = config.ChatGpt.Telegram
	}
	return tg
}

func GetTelegramKeyword() *string {
	tgKeyword := getEnv("tg_keyword")

	if tgKeyword != nil {
		return tgKeyword
	}
	if config == nil {
		return nil
	}
	if tgKeyword == nil {
		tgKeyword = config.ChatGpt.TgKeyword
	}
	return tgKeyword
}

func GetTelegramWhitelist() *string {
	tgWhitelist := getEnv("tg_whitelist")

	if tgWhitelist != nil {
		return tgWhitelist
	}
	if config == nil {
		return nil
	}
	if tgWhitelist == nil {
		tgWhitelist = config.ChatGpt.TgWhitelist
	}
	return tgWhitelist
}

func GetOpenAiApiKey() *string {
	apiKey := getEnv("api_key")

	if apiKey != nil {
		return apiKey
	}

	if config == nil {
		return nil
	}
	if apiKey == nil {
		apiKey = &config.ChatGpt.Token
	}
	return apiKey
}

func getEnv(key string) *string {
	value := os.Getenv(key)
	if len(value) == 0 {
		value = os.Getenv(strings.ToUpper(key))
	}

	if len(value) > 0 {
		return &value
	}

	if config == nil {
		return nil
	}

	if len(value) > 0 {
		return &value
	} else if config.ChatGpt.WechatKeyword != nil {
		value = *config.ChatGpt.WechatKeyword
	}
	return nil
}
