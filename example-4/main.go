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
        panic(err)
    }

    // deploy workflow
    response, err := zbClient.CreateWorkflowFromFile(topicName, zbc.BpmnXml, "order-process.bpmn")
    if err != nil {
        panic(err)
    }

    fmt.Println(response.String())

    // create a new workflow instance
    payload := make(map[string]interface{})
    payload["orderId"] = "31243"

    instance := zbc.NewWorkflowInstance("order-process", -1, payload)
    msg, err := zbClient.CreateWorkflowInstance(topicName, instance)

    if err != nil {
        panic(err)
    }

    fmt.Println(msg.String())

    // open a task subscription for the payment-service task
    subscriptionCh, subscription, err := zbClient.TaskConsumer(topicName, "sample-app", "payment-service")

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
        message := <-subscriptionCh
        fmt.Println(message.String())

        // complete task after processing
        response, _ := zbClient.CompleteTask(message)
        fmt.Println(response)
    }
}
