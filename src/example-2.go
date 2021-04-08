package main

import (
	"context"
	"fmt"
	"github.com/camunda-cloud/zeebe/clients/go/pkg/zbc"
)

const BrokerAddr = "0.0.0.0:26500"

func main() {
	zbClient, err := zbc.NewClient(&zbc.ClientConfig{
		GatewayAddress:         BrokerAddr,
		UsePlaintextConnection: true,
	})
	if err != nil {
		panic(err)
	}

	ctx := context.Background()

	response, err := zbClient.NewDeployProcessCommand().AddResourceFile("order-process-2.bpmn").Send(ctx)
	if err != nil {
		panic(err)
	}

	fmt.Println(response.String())
}
