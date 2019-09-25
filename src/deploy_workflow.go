package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/zeebe-io/zeebe/clients/go/zbc"
)

var (
	address = flag.String("address", "0.0.0.0", "Address of the Zeebe gateway")
	port    = flag.Int("port", 26500, "Destination port of the Zeebe gateway")
)

func main() {
	flag.Parse()

	client, err := zbc.NewZBClientWithConfig(&zbc.ZBClientConfig{
		GatewayAddress:         fmt.Sprintf("%s:%d", *address, *port),
		UsePlaintextConnection: true,
	})
	if err != nil {
		log.Fatal(err)
	}

	response, err := client.NewDeployWorkflowCommand().AddResourceFile("order-process.bpmn").Send()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(response)
}
