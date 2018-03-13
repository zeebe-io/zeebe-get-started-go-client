package main

import (
	"errors"
	"fmt"
	"github.com/zeebe-io/zbc-go/zbc"
	"github.com/zeebe-io/zbc-go/zbc/models/zbsubscriptions"
	"github.com/zeebe-io/zbc-go/zbc/services/zbsubscribe"
	"os"
	"os/signal"
)

const topicName = "default-topic"
const brokerAddr = "0.0.0.0:51015"

var errClientStartFailed = errors.New("cannot start client")

func handler(client zbsubscribe.ZeebeAPI, event *zbsubscriptions.SubscriptionEvent) error {
	fmt.Printf("Event: %v\n", event)
	return nil
}

func main() {
	zbClient, err := zbc.NewClient(brokerAddr)
	if err != nil {
		panic(errClientStartFailed)
	}

	subscription, err := zbClient.TopicSubscription(topicName, "subscrition-name", 128, 0, true, handler)

	if err != nil {
		panic("Failed to open subscription")
	}

	osCh := make(chan os.Signal, 1)
	signal.Notify(osCh, os.Interrupt)
	go func() {
		<-osCh
		err := subscription.Close()
		if err != nil {
			panic("Failed to close subscription")
		}
		fmt.Println("Subscription closed.")
		os.Exit(0)
	}()

	subscription.Start()
}
