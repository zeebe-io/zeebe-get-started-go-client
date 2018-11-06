package main

import (
	"fmt"
	"github.com/zeebe-io/zeebe/clients/go"
	"github.com/zeebe-io/zeebe/clients/go/pb"
)

const BrokerAddr = "0.0.0.0:26500"

func main() {
	zbClient, err := zbc.NewZBClient(BrokerAddr)
	if err != nil {
		panic(err)
	}

	topology, err := zbClient.NewTopologyCommand().Send()
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
	case  pb.Partition_LEADER:
		return "Leader"
	case pb.Partition_FOLLOWER:
		return "Follower"
	default:
		return "Unknown"
	}
}
