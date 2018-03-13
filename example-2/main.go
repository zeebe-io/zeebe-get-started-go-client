package main

import (
	"errors"
	"fmt"
	"github.com/zeebe-io/zbc-go/zbc"
	"github.com/zeebe-io/zbc-go/zbc/common"
)

const topicName = "default-topic"
const brokerAddr = "0.0.0.0:51015"

var errClientStartFailed = errors.New("cannot start client")
var errWorkflowDeploymentFailed = errors.New("creating new workflow deployment failed")

func main() {
	zbClient, err := zbc.NewClient(brokerAddr)
	if err != nil {
		panic(errClientStartFailed)
	}

	response, err := zbClient.CreateWorkflowFromFile(topicName, zbcommon.BpmnXml, "order-process.bpmn")
	if err != nil {
		panic(errWorkflowDeploymentFailed)
	}

	fmt.Println(response.String())
}
