package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/zeebe-io/zeebe/clients/go/entities"
	"github.com/zeebe-io/zeebe/clients/go/worker"
	"github.com/zeebe-io/zeebe/clients/go/zbc"
)

var (
	address = flag.String("address", "0.0.0.0" , "")
	port = flag.Int("port", 26500, "")
)

func main() {
	client, err := zbc.NewZBClientWithConfig(&zbc.ZBClientConfig{
		GatewayAddress:         fmt.Sprintf("%s:%d", *address, *port),
		UsePlaintextConnection: true,
	})
	if err != nil {
		log.Fatal(err)
	}

	worker := client.NewJobWorker().JobType("inventory-service").Handler(fetchItems).Open()
	defer worker.Close()

	worker.AwaitClose()
}

func fetchItems(client worker.JobClient, job entities.Job) {
	var err error

	defer func() {
		if err != nil {
			log.Printf("%d Failed to complete job: %v", job.GetKey(), err)
			client.NewFailJobCommand().JobKey(job.GetKey()).Retries(job.Retries - 1).Send()
		}
	}()

	variables, err := job.GetVariablesAsMap()
	if err != nil {
		return
	}

	log.Printf("%d Processing order: %v", job.GetKey(), variables["orderId"])

	variables["itemsFetched"] = 3

	request, err := client.NewCompleteJobCommand().JobKey(job.GetKey()).VariablesFromMap(variables)
	if err != nil {
		return
	}

	log.Printf("%d Number of items fetched: %v", job.GetKey(), variables["itemsFetched"])

	request.Send()
}
