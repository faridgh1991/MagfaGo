# MagfaGo
Magfa SDK for Golang


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

