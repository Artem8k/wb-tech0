package sub

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"rest-api/src/database"
	"rest-api/src/database/models"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
)

func Subscribe(repo *database.Repository) *nats.Conn {
	nc, err := nats.Connect("nats:4222")
	if err != nil {
		log.Fatalln(err)
	}

	js, err := jetstream.New(nc)

	if err != nil {
		fmt.Println(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	stream, err := js.CreateStream(ctx, jetstream.StreamConfig{
		Name:     "test",
		Subjects: []string{"Order"},
	})

	if err != nil {
		fmt.Println(err)
	}

	cons, err := stream.CreateOrUpdateConsumer(ctx, jetstream.ConsumerConfig{
		Durable: "testConsumer",
	})

	if err != nil {
		fmt.Println(err)
	}

	_, err = cons.Consume(func(msg jetstream.Msg) {
		var o *models.Order = &models.Order{}

		if err := json.Unmarshal(msg.Data(), &o); err != nil {
			log.Println(err)
		}

		if err != nil {
			log.Println(err)
		}

		repo.Create(*o)
		msg.Ack()
	})

	if err != nil {
		fmt.Println(err)
	}

	return nc
}
