package pub

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"rest-api/src/database/models"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
)

func Publish() {
	nc, err := nats.Connect("nats:4222")

	js, _ := jetstream.New(nc)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err = js.CreateStream(ctx, jetstream.StreamConfig{
		Name:     "test",
		Subjects: []string{"Order"},
	})

	if err != nil {
		log.Fatalln(err)
	}

	t := time.Now()

	order := &models.Order{
		OrderUid:     fmt.Sprint(rand.Int()),
		Track_number: "WBILMTESTTRACK",
		Entry:        "WBIL",
		Delivery: models.Delivery{
			Name:    "Test Testov",
			Phone:   "+9720000000",
			Zip:     "2639809",
			City:    "Kiryat Mozkin",
			Address: "Ploshad Mira 15",
			Region:  "Kraiot",
			Email:   "test@gmail.com",
		},
		Payment: models.Payment{
			Transaction:   fmt.Sprint(rand.Int()),
			Request_id:    "",
			Currency:      "USD",
			Provider:      "wbpay",
			Amount:        1817,
			Payment_dt:    1637907727,
			Bank:          "alpha",
			Delivery_cost: 1500,
			Goods_total:   317,
			Custom_fee:    0,
		},
		Items: []models.Item{{
			Chrt_id:      9934930,
			Track_number: "WBILMTESTTRACK",
			Price:        453,
			Rid:          "ab4219087a764ae0btest",
			Name:         "Mascaras",
			Sale:         30,
			Size:         "0",
			Total_price:  317,
			Nm_id:        2389212,
			Brand:        "Vivienne Sabo",
			Status:       202,
		}},
		Locale:             "en",
		Internal_signature: "",
		Customer_id:        "test",
		Delivery_service:   "meest",
		Shardkey:           "9",
		Sm_id:              99,
		Date_created:       t.Format(time.RFC3339),
		Oof_shard:          "1",
	}

	json, err := json.Marshal(order)

	if err != nil {
		fmt.Println(err)
	}

	js.Publish(ctx, "Order", json)

}
