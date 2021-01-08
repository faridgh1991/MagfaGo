package magfa

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type MessagesResult struct {
	Messages []Message
	Status   int
}

type Message struct {
	Body            string
	SenderNumber    string
	RecipientNumber string
	Date            CustomTime
}

func (c *Client) GetMessages(count int) ([]Message, error) {
	if count < 1 || count > 100 {
		return nil, fmt.Errorf("message counts must be between 1 and 100")
	}
	path := fmt.Sprintf("messages/%d", count)
	raw, err := c.sendRequest(path, http.MethodGet, nil)
	if err != nil {
		return nil, err
	}

	var resp MessagesResult
	err = json.Unmarshal(raw, &resp)
	if err != nil {
		return nil, err
	}
	if resp.Status != 0 {
		return nil, newError(resp.Status)
	}

	return resp.Messages, nil
}
