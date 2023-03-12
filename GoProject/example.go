package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/sashabaranov/go-openai"
	"net/http"
)

func _chatGPT(c *gin.Context) {
	userComment, isCommentOK := c.GetQuery("comment")
	userKey, isKeyOK := c.GetQuery("key")
	if isKeyOK && isCommentOK {
		client := openai.NewClient(userKey)
		response, err := client.CreateChatCompletion(
			context.Background(),
			openai.ChatCompletionRequest{
				Model: openai.GPT3Dot5Turbo,
				Messages: []openai.ChatCompletionMessage{
					{
						Role:    openai.ChatMessageRoleUser,
						Content: userComment,
					},
				},
			},
		)
		if err != nil {
			return
		}
		c.String(http.StatusOK, response.Choices[0].Message.Content)
	}

}

func main() {
	r := gin.Default()
	r.GET("/chat", _chatGPT)
	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}
