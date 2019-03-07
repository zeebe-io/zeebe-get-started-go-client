package main

import (
	"fmt"
	"github.com/zeebe-io/zeebe/clients/go/zbc"
)

const brokerAddr = "0.0.0.0:26500"

func main() {
	client, err := zbc.NewZBClient(brokerAddr)
	if err != nil {
		panic(err)
	}

	// After the workflow is deployed.
	payload := make(map[string]interface{})
	payload["orderId"] = "31243"

	request, err := client.NewCreateInstanceCommand().BPMNProcessId("order-process").LatestVersion().VariablesFromMap(payload)
	if err != nil {
		panic(err)
	}

	msg, err := request.Send()
	if err != nil {
		panic(err)
	}

	fmt.Println(msg.String())
}
