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

	response, err := zbClient.NewDeployWorkflowCommand().AddResourceFile("order-process.bpmn").Send()
	if err != nil {
		panic(err)
	}

	fmt.Println(response.String())
}
