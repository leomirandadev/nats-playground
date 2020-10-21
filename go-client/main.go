package main

import (
	"fmt"
	"strconv"
	"time"

	nats "github.com/nats-io/nats.go"
)

type msg struct {
	Client  string `json:"client"`
	Message string `json:"message"`
}

func main() {

	var channel string = "cannel-test"
	var clientName string = generateClientName()

	c := connectNats()

	sendMessage(c, channel, msg{Client: clientName, Message: "The Dad is on"})

	subscribeChannel(c, channel)

	continueListen()

}

func generateClientName() string {
	milisecondsNow := time.Now().UnixNano() / int64(time.Millisecond)

	clientName := "go-client-" + strconv.FormatInt(milisecondsNow, 10)

	fmt.Println("Client name:", clientName)

	return clientName
}

func continueListen() {
	fmt.Println("Listen")
	select {}
}

func connectNats() *nats.EncodedConn {
	nc, err := nats.Connect(nats.DefaultURL)

	if err != nil {
		fmt.Println("Error on conection")
	}

	fmt.Println("Connected on", nats.DefaultURL)

	c, _ := nats.NewEncodedConn(nc, nats.JSON_ENCODER)

	return c
}

func sendMessage(c *nats.EncodedConn, channel string, message msg) {

	err1 := c.Publish(channel, message)

	if err1 != nil {
		fmt.Println("Error on publish message")
	}

	fmt.Println("Send message to:", channel)
}

func subscribeChannel(c *nats.EncodedConn, channel string) {

	_, err2 := c.Subscribe(channel, func(m *nats.Msg) {
		fmt.Printf("Received a message: %s\n", string(m.Data))
	})

	if err2 != nil {
		fmt.Println("Error on subcribe on a chanel")
	}

	fmt.Println("Subscribe successfuly on", channel)
}
