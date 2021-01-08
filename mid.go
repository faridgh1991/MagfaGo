package magfa

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type MidResult struct {
	Mid    int64
	Status int
}

func (c *Client) GetMid(uid int) (int64, error) {
	path := fmt.Sprintf("mid/%d", uid)
	raw, err := c.sendRequest(path, http.MethodGet, nil)
	if err != nil {
		return 0, err
	}

	var resp MidResult
	err = json.Unmarshal(raw, &resp)
	if err != nil {
		return 0, err
	}
	if resp.Status != 0 {
		return 0, newError(resp.Status)
	}

	return resp.Mid, nil
}
