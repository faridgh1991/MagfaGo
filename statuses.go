package magfa

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type StatusesResult struct {
	Status int
	Dlrs   []DeliveryStatus
}

type DeliveryStatus struct {
	Mid        int64
	Status     int
	StatusText string
	Date       CustomTime
}

var deliveryStatuses = map[int]string{
	-1: "Invalid MID",
	0:  "No Status",
	1:  "Received by destination",
	2:  "Not yet Received",

	8:  "Pending on ITC",
	16: "Forwarding to ITC",
}

type CustomTime struct {
	time.Time
}

const ctLayout = "2006-01-02 15:04:05"

func (ct *CustomTime) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), "\"")
	if s == "null" {
		ct.Time = time.Time{}
		return
	}
	ct.Time, err = time.Parse(ctLayout, s)
	return
}

func (c *Client) GetStatuses(mids []int64) ([]DeliveryStatus, error) {
	var valuesText []string
	for i := range mids {
		number := mids[i]
		text := strconv.Itoa(int(number))
		valuesText = append(valuesText, text)
	}

	midsString := strings.Join(valuesText, ",")
	path := fmt.Sprintf("statuses/%s", midsString)
	raw, err := c.sendRequest(path, http.MethodGet, nil)
	if err != nil {
		return nil, err
	}

	var resp StatusesResult
	err = json.Unmarshal(raw, &resp)
	if err != nil {
		return nil, err
	}
	if resp.Status != 0 {
		return nil, newError(resp.Status)
	}
	for i, status := range resp.Dlrs {
		resp.Dlrs[i].StatusText = deliveryStatuses[status.Status]
	}

	return resp.Dlrs, nil
}
