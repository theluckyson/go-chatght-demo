package main

import (
	"github.com/gin-gonic/gin"
)

import "net/http"
import "context"
import "github.com/sashabaranov/go-openai"

func chat_gpt(c *gin.Context) {
	usercomment, ok1 := c.GetQuery("comment")
	userkey, ok2 := c.GetQuery("key")
	if ok1 && ok2 {

		client := openai.NewClient(userkey)
		resp, err := client.CreateChatCompletion(
			context.Background(),
			openai.ChatCompletionRequest{
				Model: openai.GPT3Dot5Turbo,
				Messages: []openai.ChatCompletionMessage{
					{
						Role:    openai.ChatMessageRoleUser,
						Content: usercomment,
					},
				},
			},
		)

		if err != nil {
			return
		}

		c.String(http.StatusOK, resp.Choices[0].Message.Content)

	}

}

func main() {

	// 初始化一个http服务对象
	r := gin.Default()

	// 设置一个GET请求的路由，url: '/ping'， 控制器函数： 闭包
	r.GET("/chat", chat_gpt)

	// 监听，并在 localhost:8080上启动服务
	r.Run(":8080")

}
