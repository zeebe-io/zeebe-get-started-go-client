package main

import (
	"fmt"
	"github.com/zeebe-io/zeebe/clients/go/zbc"
	"time"
)

const brokerAddr = "0.0.0.0:26500"

func main() {
	client, err := zbc.NewZBClient(brokerAddr)
	if err != nil {
		panic(err)
	}

	// deploy workflow
	response, err := client.NewDeployWorkflowCommand().AddResourceFile("order-process.bpmn").Send()
	if err != nil {
		panic(err)
	}

	fmt.Println(response.String())

	// create a new workflow instance
	payload := make(map[string]interface{})
	payload["orderId"] = "31243"

	request, err := client.NewCreateInstanceCommand().BPMNProcessId("order-process").LatestVersion().PayloadFromMap(payload)
	if err != nil {
		panic(err)
	}

	result, err := request.Send()
	if err != nil {
		panic(err)
	}

	fmt.Println(result.String())

	// sleep to allow job to be created
	time.Sleep(1 * time.Second)

	jobs, err := client.NewActivateJobsCommand().JobType("payment-service").Amount(1).WorkerName("sample-app").Timeout(1 * time.Minute).Send()
	if err != nil {
		panic(err)
	}

	for _, job := range jobs {
		client.NewCompleteJobCommand().JobKey(job.GetKey()).Send()
		fmt.Println("Completed job", job.String())
	}
}
