package magfa

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// SendRequest is structure of send sms requests with senders, recipients, messages and etc.
type SendRequest struct {
	Senders    []string `json:"senders,omitempty"`
	Recipients []string `json:"recipients,omitempty"`
	Messages   []string `json:"messages,omitempty"`
	Encodings  []int    `json:"encodings,omitempty"`
	Uids       []int64  `json:"uids,omitempty"`
	Udhs       []string `json:"udhs,omitempty"`
}

type sendResult struct {
	Messages []SendResponse
	Status   int
}

// SendResponse is result of send Request for each message
type SendResponse struct {
	Status     int
	Id         int64
	UserId     int64
	Parts      int
	Tariff     float64
	Alphabet   string
	Recipient  string
	StatusText string
}

// Send a send request to magfa server
func (c *Client) Send(request SendRequest) ([]SendResponse, error) {

	body, err := json.Marshal(request)
	raw, err := c.sendRequest("send", http.MethodPost, body)
	if err != nil {
		return nil, fmt.Errorf("send failed: %s", err.Error())
	}

	var resp sendResult
	err = json.Unmarshal(raw, &resp)
	if err != nil {
		return nil, fmt.Errorf("unmarshal response: %s", err.Error())
	}
	if resp.Status != 0 {
		return nil, newError(resp.Status)
	}
	for i, msg := range resp.Messages {
		if m, ok := errorMap[msg.Status]; ok {
			resp.Messages[i].StatusText = m
		}
	}

	return resp.Messages, nil
}
