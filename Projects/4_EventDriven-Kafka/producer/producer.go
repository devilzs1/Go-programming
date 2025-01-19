package producer

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/IBM/sarama"
	"github.com/gofiber/fiber/v2"
)

type Comment struct {
	Text string `form:"text" json:"text"`
}

func Producer() {
	app := fiber.New()
	api := app.Group("/api/v1")
	api.Post("/comments", createComment)
	log.Fatal(app.Listen(":3000"))
}

func ConnectProducer(brokersUrl []string) (sarama.SyncProducer, error) {
	config := sarama.NewConfig()

	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 3

	conn, err := sarama.NewSyncProducer(brokersUrl, config)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func createComment(ctx *fiber.Ctx) error {
	cmt := new(Comment)
	if err := ctx.BodyParser(cmt); err != nil {
		log.Println(err)
		ctx.Status(400).JSON(&fiber.Map{
			"success": false,
			"message": err.Error(),
		})
		return err
	}

	cmtInBytes, err := json.Marshal(cmt)
	if err != nil {
		log.Println(err)
		ctx.Status(500).JSON(&fiber.Map{
			"success": false,
			"message": "Failed to marshal comment",
		})
		return err
	}

	PushCommentsToQueue("comments", cmtInBytes)

	err = ctx.JSON(&fiber.Map{
		"success": true,
		"message": "Comment pushed to queue successfully",
		"comment": cmt,
	})

	if err != nil {
		log.Println(err)
		ctx.Status(500).JSON(&fiber.Map{
			"success": false,
			"message": "Failed to map comment/ Error creating product",
		})
		return err
	}

	return nil
}

func PushCommentsToQueue(topic string, message []byte) error {
	brokersUrl := []string{"localhost:29092"}
	producer, err := ConnectProducer(brokersUrl)
	if err != nil {
		log.Fatal("Error connecting to producer : ", err)
		return err
	}
	defer producer.Close()

	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(message),
	}
	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		log.Fatal("Error sending message : ", err)
		return err
	}
	fmt.Printf("Message is stored in topic(%s)/partition(%d)/offset(%d)\n", topic, partition, offset)

	return nil
}
