package magfa

import "fmt"

var errorMap = map[int]string{
	1: "Invalid Recipient",
	2: "Invalid Sender",
	3: "Invalid Encoding",
	4: "Invalid mClass",

	6: "Invalid UDH",

	13: "Empty Message",
	14: "Not Enough Credit left",
	15: "Remote Server Temporary Failure, Try again",
	16: "Account Disabled",
	17: "Account Expired",
	18: "Wrong Username/Password",
	19: "Authentication Failure",
	20: "Account Not Allowed To Use This Sender Number",

	22: "Account Not Allowed To Use This Service",
	23: "Remote Server Busy, Try Again Later",
	24: "Invalid Message ID",
	25: "Invalid Service Type",

	27: "recipient number is blocked by operator",
	28: "recipient number is blocked by Magfa",
	29: "IP Address Not Allowed To Use This Sender Number",
	30: "Too Many message parts(>256)",

	101: "message bodies and recipient numbers length Miss-Matched",
	102: "message class and recipient numbers length Miss-Matched",
	103: "sender numbers and recipient numbers length Miss-Matched",
	104: "udhs and recipient numbers length Miss-Matched",
	105: "priorities and recipient numbers length Miss-Matched",
	106: "Empty recipient numbers",
	107: "recipient numbers over-sized",
	108: "Empty sender numbers",
	109: "priorities and recipientNumbers length Miss-Matched",
	110: "checkingMessageIds and recipientNumbers length Miss-Matched",
}

func newError(code int) error {
	message := "Unknown Error"
	if msg, ok := errorMap[code]; ok {
		message = msg
	}
	return fmt.Errorf("%s, code:%d", message, code)
}
