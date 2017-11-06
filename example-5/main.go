package main

import (
	"errors"
	"fmt"
	"github.com/zeebe-io/zbc-go/zbc"
	"os"
	"os/signal"
)

const topicName = "default-topic"
const brokerAddr = "0.0.0.0:51015"

var errClientStartFailed = errors.New("cannot start client")

func main() {
	zbClient, err := zbc.NewClient(brokerAddr)
	if err != nil {
		panic(errClientStartFailed)
	}

	subscriptionCh, sub, err := zbClient.TopicConsumer(topicName, "subscription-name", 0)

	osCh := make(chan os.Signal, 1)
	signal.Notify(osCh, os.Interrupt)
	go func() {
		<-osCh
		fmt.Println("Closing subscription.")
		_, err := zbClient.CloseTopicSubscription(sub)
		if err != nil {
			fmt.Println("failed to close subscription: ", err)
		} else {
			fmt.Println("Subscription closed.")
		}
		os.Exit(0)
	}()

	for {
		message := <-subscriptionCh
		fmt.Println(message.String())
	}

}
