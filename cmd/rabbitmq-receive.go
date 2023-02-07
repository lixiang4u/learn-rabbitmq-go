package cmd

import (
	"fmt"
	"github.com/lixiang4u/learn-rabbitmq-go/utils"
	"github.com/rabbitmq/amqp091-go"
	"github.com/spf13/cobra"
	"log"
)

var mqReceiveCmd = &cobra.Command{
	Use:   "mqReceive",
	Short: "mqReceive",
	Run: func(cmd *cobra.Command, args []string) {
		runMqReceiveCmd()
	},
}

func init() {
	rootCmd.AddCommand(mqReceiveCmd)
}

func runMqReceiveCmd() {
	log.Println("runMqReceiveCmd")

	conn, err := amqp091.Dial("amqp://guest:guest@localhost:5672/")
	defer func() { _ = conn.Close() }()
	utils.PrintMqFailOnError(err, "Failed to connect to RabbitMQ")

	ch, err := conn.Channel()
	defer func() { _ = ch.Close() }()
	utils.PrintMqFailOnError(err, "Failed to open channel")

	// 重复的声明使从Queue取数据时Queue存在
	q, err := ch.QueueDeclare("queue_1", false, false, false, false, nil)
	utils.PrintMqFailOnError(err, "Failed to declare a queue")

	// 消费数据
	msgs, err := ch.Consume(q.Name, "", true, false, false, false, nil)
	utils.PrintMqFailOnError(err, "Failed to register a consumer")

	var forever chan struct{}
	go func() {
		for d := range msgs {
			log.Println(fmt.Sprintf("Received a message: %s", d.Body))
		}
	}()

	log.Println("[*] Waiting for message. To exit press CTRL+C")
	<-forever

}
