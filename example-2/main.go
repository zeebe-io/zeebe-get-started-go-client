package main

import (
	"fmt"
	"github.com/zeebe-io/zeebe/clients/go"
)

const brokerAddr = "0.0.0.0:26500"

func main() {
	zbClient, err := zbc.NewZBClient(brokerAddr)
	if err != nil {
		panic(err)
	}

	response, err := zbClient.NewDeployWorkflowCommand().AddResourceFile("order-process.bpmn").Send()
	if err != nil {
		panic(err)
	}

	fmt.Println(response.String())
}
