package magfa

import (
	"encoding/json"
	"net/http"
)

type balanceResult struct {
	Balance int
	Status  int
}

// GetBalance of account from magfa
func (c *Client) GetBalance() (int, error) {
	raw, err := c.sendRequest("balance", http.MethodGet, nil)
	if err != nil {
		return 0, err
	}

	var resp balanceResult
	err = json.Unmarshal(raw, &resp)
	if err != nil {
		return 0, err
	}
	if resp.Status != 0 {
		return 0, newError(resp.Status)
	}

	return resp.Balance, nil
}
