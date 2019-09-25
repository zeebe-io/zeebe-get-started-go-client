package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/zeebe-io/zeebe/clients/go/pb"
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

	topology, err := client.NewTopologyCommand().Send()
	if err != nil {
		log.Fatal(err)
	}

	for _, broker := range topology.Brokers {
		fmt.Printf("Broker %s:%d\n", broker.Host, broker.Port)
		for _, partition := range broker.Partitions {
			fmt.Printf("  Partition %d: %s\n", partition.PartitionId, roleToString(partition.Role))
		}
	}
}

func roleToString(role pb.Partition_PartitionBrokerRole) string {
	switch role {
	case pb.Partition_LEADER:
		return "Leader"
	case pb.Partition_FOLLOWER:
		return "Follower"
	default:
		return "Unknown"
	}
}
