package magfa

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type messagesResult struct {
	Messages []Message
	Status   int
}

// Message is structure of received message
type Message struct {
	Body            string
	SenderNumber    string
	RecipientNumber string
	Date            CustomTime
}

// GetMessages that received to magfa servers
func (c *Client) GetMessages(count int) ([]Message, error) {
	if count < 1 || count > 100 {
		return nil, fmt.Errorf("message counts must be between 1 and 100")
	}
	path := fmt.Sprintf("messages/%d", count)
	raw, err := c.sendRequest(path, http.MethodGet, nil)
	if err != nil {
		return nil, err
	}

	var resp messagesResult
	err = json.Unmarshal(raw, &resp)
	if err != nil {
		return nil, err
	}
	if resp.Status != 0 {
		return nil, newError(resp.Status)
	}

	return resp.Messages, nil
}
