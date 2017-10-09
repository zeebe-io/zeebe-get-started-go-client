package main

import (
	"errors"
	"github.com/zeebe-io/zbc-go/zbc"
	"fmt"
)

const topicName = "default-topic"
const brokerAddr = "0.0.0.0:51015"

var errClientStartFailed = errors.New("cannot start client")

func main() {
	zbClient, err := zbc.NewClient(brokerAddr)
	if err != nil {
		panic(errClientStartFailed)
	}

	// After the workflow is deployed.

	instance := zbc.NewWorkflowInstance("order-process", -1, nil)
	msg, err := zbClient.CreateWorkflowInstance(topicName, instance)

	if err != nil {
		panic(err)
	}

	fmt.Println(msg.String())
}
