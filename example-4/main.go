package main

import (
	"fmt"
	"github.com/zeebe-io/zeebe/clients/go/entities"
	"github.com/zeebe-io/zeebe/clients/go/worker"
	"github.com/zeebe-io/zeebe/clients/go/zbc"
	"log"
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

	request, err := client.NewCreateInstanceCommand().BPMNProcessId("order-process").LatestVersion().VariablesFromMap(payload)
	if err != nil {
		panic(err)
	}

	result, err := request.Send()
	if err != nil {
		panic(err)
	}

	fmt.Println(result.String())

	jobWorker := client.NewJobWorker().JobType("payment-service").Handler(handleJob).Open()
	defer jobWorker.Close()

	jobWorker.AwaitClose()
}

func handleJob(client worker.JobClient, job entities.Job) {
	jobKey := job.GetKey()

	headers, err := job.GetCustomHeadersAsMap()
	if err != nil {
		// failed to handle job as we require the custom job headers
		failJob(client, job)
		return
	}

	payload, err := job.GetPayloadAsMap()
	if err != nil {
		// failed to handle job as we require the payload
		failJob(client, job)
		return
	}

	payload["totalPrice"] = 46.50;
	request, err := client.NewCompleteJobCommand().JobKey(jobKey).PayloadFromMap(payload)
	if err != nil {
		// failed to set the updated payload
		failJob(client, job)
		return
	}

	log.Println("Complete job", jobKey, "of type", job.Type)
	log.Println("Processing order:", payload["orderId"])
	log.Println("Collect money using payment method:", headers["method"])

	request.Send()
}

func failJob(client worker.JobClient, job entities.Job) {
	log.Println("Failed to complete job", job.GetKey())
	client.NewFailJobCommand().JobKey(job.GetKey()).Retries(job.Retries - 1).Send()
}
