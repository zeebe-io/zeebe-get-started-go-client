package main

import (
	"errors"
	"fmt"
	"github.com/zeebe-io/zbc-go/zbc"
	"os"
	"os/signal"
	"log"
)

const topicName = "default-topic"
const brokerAddr = "0.0.0.0:51015"

var errClientStartFailed = errors.New("cannot start client")

func main() {
	zbClient, err := zbc.NewClient(brokerAddr)
	if err != nil {
		panic(errClientStartFailed)
	}

	payload := make(map[string]interface{})
	payload["a"] = "b"
	instance := zbc.NewWorkflowInstance("order-process", -1, payload)

	msg, err := zbClient.CreateWorkflowInstance(topicName, instance)
	fmt.Println(msg.String())
	if err != nil {
		panic(err)
	}

	subscriptionCh, subscription, err := zbClient.TaskConsumer(topicName, "sample-app", "payment-service")
	fmt.Println(subscription.String())

	osCh := make(chan os.Signal, 1)
	signal.Notify(osCh, os.Interrupt)
	go func() {
		<-osCh
		fmt.Println("Closing subscription.")
		_, err := zbClient.CloseTaskSubscription(subscription)
		if err != nil {
			fmt.Println("failed to close subscription: ", err)
		} else {
			fmt.Println("Subscription closed.")
		}
		os.Exit(0)
	}()

	for {
		log.Println("Waiting for events")
		message := <-subscriptionCh
		fmt.Println(message.String())


		response, _ := zbClient.CompleteTask(message)
		fmt.Println(response)
	}

}
