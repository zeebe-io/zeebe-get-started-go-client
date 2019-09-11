package main

import (
	"fmt"

	"github.com/zeebe-io/zeebe/clients/go/zbc"
)

const BrokerAddr = "0.0.0.0:26500"

func main() {
	zbClient, err := zbc.NewZBClientWithConfig(&zbc.ZBClientConfig{
		GatewayAddress: BrokerAddr,
		UsePlaintextConnection: true})
	if err != nil {
		panic(err)
	}

	// After the workflow is deployed.
	variables := make(map[string]interface{})
	variables["orderId"] = "31243"

	request, err := zbClient.NewCreateInstanceCommand().BPMNProcessId("order-process").LatestVersion().VariablesFromMap(variables)
	if err != nil {
		panic(err)
	}

	msg, err := request.Send()
	if err != nil {
		panic(err)
	}

	fmt.Println(msg.String())
}
