package dingtalk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type DingTalkRequestBody struct {
	MsgType string `json:"msgtype"`
	Text    struct {
		Content string `json:"content"`
	} `json:"text"`
}

func SendDingTalkMessage(WebhookURL string, content string) error {
	// 构建钉钉机器人请求体
	messageBody := DingTalkRequestBody{
		MsgType: "text",
		Text: struct {
			Content string `json:"content"`
		}{
			Content: content,
		},
	}
	// 将结构体转换为JSON
	messageJSON, err := json.Marshal(messageBody)
	if err != nil {
		return err
	}
	// 发送HTTP POST请求
	resp, err := http.Post(WebhookURL, "application/json", bytes.NewBuffer(messageJSON))
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)
	// 检查HTTP响应
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("HTTP请求失败，状态码: %d", resp.StatusCode)
	}
	fmt.Println("消息已成功发送到钉钉！")
	return nil
}
