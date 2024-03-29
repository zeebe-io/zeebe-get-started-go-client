package main

import (
	"context"
	"fmt"
	"github.com/camunda/zeebe/clients/go/v8/pkg/pb"
	"github.com/camunda/zeebe/clients/go/v8/pkg/zbc"
	"os"
)

const ZeebeAddr = "0.0.0.0:26500"

/*
Sample application that connects to a cluster on Camunda Cloud, or a locally deployed cluster.

When connecting to a cluster in Camunda Cloud, this application assumes that the following
environment variables are set:

ZEEBE_ADDRESS
ZEEBE_CLIENT_ID
ZEEBE_CLIENT_SECRET
ZEEBE_AUTHORIZATION_SERVER_URL

Hint: When you create client credentials in Camunda Cloud you have the option
to download a file with the lines above filled out for you.

When connecting to a local cluster or node, this application assumes default port and no
authentication or encryption enabled.
*/
func main() {
	gatewayAddr := os.Getenv("ZEEBE_ADDRESS")
	var plainText bool

	if gatewayAddr == "" {
		gatewayAddr = ZeebeAddr
		plainText = true
	}

	zbClient, err := zbc.NewClient(&zbc.ClientConfig{
		GatewayAddress:         gatewayAddr,
		UsePlaintextConnection: plainText,
	})

	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	topology, err := zbClient.NewTopologyCommand().Send(ctx)
	if err != nil {
		panic(err)
	}

	for _, broker := range topology.Brokers {
		fmt.Println("Broker", broker.Host, ":", broker.Port)
		for _, partition := range broker.Partitions {
			fmt.Println("  Partition", partition.PartitionId, ":", roleToString(partition.Role))
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
