# MagfaGo
[![Go Report Card](https://goreportcard.com/badge/github.com/faridgh1991/MagfaGo)](https://goreportcard.com/report/github.com/faridgh1991/MagfaGo)
[![GoDoc](https://godoc.org/github.com/faridgh1991/MagfaGo?status.svg)](https://pkg.go.dev/github.com/faridgh1991/MagfaGo)
[![license](https://img.shields.io/badge/License-MIT-blue.svg)](https://github.com/faridgh1991/MagfaGo/blob/master/LICENSE)

Magfa SDK for Golang

- base on [Magfa http API v2](https://messaging.magfa.com/ui/?public/wiki/api/http_v2)


### Install
```bash
go get github.com/faridgh1991/MagfaGo
```

### Usage
```go
package main

import (
	"fmt"
	"github.com/faridgh1991/MagfaGo"
	"log"
	"time"
)

func main() {
	client, err := magfa.New("username", "domain", "password", time.Second)
	if err != nil {
		log.Fatal(err)
	}

	responses, err := client.Send(magfa.SendRequest{
		Senders:    []string{"3000XXXXX"},
		Recipients: []string{"98912XXXXXXX", "935XXXXXXX"},
		Messages:   []string{"Hi"},
		Encodings:  nil,
		Uids:       []int64{12345},
		Udhs:       nil,
	})
	fmt.Printf("%+v ,%v\n", responses, err)
}
```

