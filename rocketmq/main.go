package main

/*
import (
	"fmt"

	rocketmq "github.com/apache/rocketmq-client-go/core"
)

func main() {
	SendMessage()
}

func SendMessagge() {
	producer := rocketmq.NewProducer(config)
	producer.Start()
	defer producer.Shutdown()
	fmt.Printf("Producer: %s started... \n", producer)
	for i := 0; i < 100; i++ {
		msg := fmt.Sprintf("%s-*d", *body, i)
		result, err := producer.SendMessageSync(&rocketmq.Message{Topic: "test", Body: msg})
		if err != nil {
			fmt.Println("Error:", err)
		}
		fmt.Printf("send message: %s result: %s\n", msg, result)
	}
}
*/
