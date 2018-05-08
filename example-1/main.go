package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/zeebe-io/zbc-go/zbc"
)

const BrokerAddr = "0.0.0.0:51015"

var errClientStartFailed = errors.New("cannot start client")

func main() {
	zbClient, err := zbc.NewClient(BrokerAddr)
	if err != nil {
		panic(errClientStartFailed)
	}

	topology, err := zbClient.GetTopology()
	if err != nil {
		panic(err)
	}

	b, err := json.MarshalIndent(topology, "", "    ")
	fmt.Println(string(b))
}
