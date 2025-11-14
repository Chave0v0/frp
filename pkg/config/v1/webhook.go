package v1

import (
	"fmt"
)

var WebHookCfg WebHookConfig

// 定义webhook客户端接口
type WebHookClient interface {
	SendMessage(string)
}

// 定义常用webhook常量
const (
	DingTalk = iota
)

// 用于获取指定平台的webhook client
func NewWebHookClient(config WebHookConfig) WebHookClient {
	switch config.Platform {
	case DingTalk:
		return NewDingTalkWebhookClient(config)
	default:
		return nil // 其他平台等待扩展
	}
}

// 定义webhook配置基础结构体
type WebHookConfig struct {
	Platform int    `json:"platform,omitempty"`
	Url      string `json:"url,omitempty"`
}

// 定义dingtalk机器人配置
type DingTalkWebhookClient struct {
	WebHookConfig
}

func NewDingTalkWebhookClient(config WebHookConfig) *DingTalkWebhookClient {
	return &DingTalkWebhookClient{config}
}

func (client *DingTalkWebhookClient) SendMessage(content string) {
	fmt.Printf(content)
}
