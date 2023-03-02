package gogpt_test

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"testing"

	. "github.com/wangjiancn/go-gpt3"
	"github.com/wangjiancn/go-gpt3/internal/test"
)

func TestCreateChatCompletion(t *testing.T) {
	// Client portion of the test
	config := DefaultConfig(test.GetTestToken())

	config.HTTPClient.Transport = &tokenRoundTripper{
		test.GetTestToken(),
		http.DefaultTransport,
	}

	client := NewClientWithConfig(config)
	ctx := context.Background()

	chatMessages := make([]ChatCompletionMessage, 0)
	chatMessages = append(chatMessages, ChatCompletionMessage{
		Role:    "user",
		Content: "Hello",
	})
	request := ChatCompletionRequest{
		Model:    GPT3Dot5Turbo,
		Messages: chatMessages,
	}

	res, err := client.CreateChatCompletion(ctx, request)
	if err != nil {
		t.Errorf("CreateCompletionStream returned error: %v", err)
	}
	bData, _ := json.Marshal(res)
	fmt.Println(string(bData))
}

func TestCreateChatCompletionStream(t *testing.T) {
	// Client portion of the test
	config := DefaultConfig(test.GetTestToken())

	config.HTTPClient.Transport = &tokenRoundTripper{
		test.GetTestToken(),
		http.DefaultTransport,
	}

	client := NewClientWithConfig(config)
	ctx := context.Background()

	chatMessages := make([]ChatCompletionMessage, 0)
	chatMessages = append(chatMessages, ChatCompletionMessage{
		Role:    "user",
		Content: "Hello",
	})
	request := ChatCompletionRequest{
		Model:    GPT3Dot5Turbo,
		Messages: chatMessages,
	}

	stream, err := client.CreateChatCompletionStream(ctx, request)
	if err != nil {
		t.Errorf("CreateCompletionStream returned error: %v", err)
	}
	defer stream.Close()

	for {
		receivedResponse, streamErr := stream.Recv()
		if errors.Is(streamErr, io.EOF) {
			fmt.Println("stream finished")
			return
		}

		if streamErr != nil {
			t.Errorf("stream.Recv() failed: %v", streamErr)
			return
		}
		bData, _ := json.Marshal(receivedResponse)
		fmt.Println(string(bData))
	}
}
