package cmd

import (
	"context"
	"fmt"
	"github.com/lixiang4u/learn-rabbitmq-go/utils"
	"github.com/rabbitmq/amqp091-go"
	"github.com/spf13/cobra"
	"log"
	"time"
)

var mqSendCmd = &cobra.Command{
	Use:   "mqSend",
	Short: "mqSend",
	Run: func(cmd *cobra.Command, args []string) {
		runMqSendCmd()
	},
}

func init() {
	rootCmd.AddCommand(mqSendCmd)
}

func runMqSendCmd() {
	log.Println("runMqSendCmd")

	conn, err := amqp091.Dial("amqp://guest:guest@localhost:5672/")
	defer func() { _ = conn.Close() }()
	utils.PrintMqFailOnError(err, "Failed to connect to RabbitMQ")

	ch, err := conn.Channel()
	defer func() { _ = ch.Close() }()
	utils.PrintMqFailOnError(err, "Failed to open channel")

	q, err := ch.QueueDeclare("queue_1", false, false, false, false, nil)
	utils.PrintMqFailOnError(err, "Failed to declare a queue")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	body := "hello world, " + time.Now().String()
	err = ch.PublishWithContext(ctx, "", q.Name, false, false, amqp091.Publishing{
		ContentType: "text/plain",
		Body:        []byte(body),
	})
	utils.PrintMqFailOnError(err, "Failed to publish a message")
	log.Println(fmt.Sprintf("[x] Sent %s", body))

}
