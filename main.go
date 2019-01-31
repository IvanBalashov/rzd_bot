package main

import (
	"fmt"
	"log"
	"rzd_bot/reporting"
	"rzd_bot/server/rabbitmq"
	"encoding/json"
	"time"
)

func init() {
	log.SetFlags(log.LstdFlags)
}

type Configuration struct {
	RabbitMQUrl string
}

type Answer struct {
	TrainID   string  `json:"train_id"`
	MainRoute string  `json:"main_route"`
	Segment   string  `json:"segment"`
	StartDate string  `json:"start_date"`
	Seats     []Seats `json:"seats"`
}

type Seats struct {
	Name  string `json:"name"`
	Count int    `json:"count"`
	Price int    `json:"price"`
}

func main() {
	config := Configuration{
		RabbitMQUrl: "amqp://guest:guest@localhost:5672/",
	}

	logs := make(chan string)
	defer close(logs)

	logger := reporting.NewLogger(logs, "bot")
	logger.Start()

	time.Sleep(100 * time.Millisecond)

	server, err := rabbitmq.NewServer(config.RabbitMQUrl, nil, logs)
	if err != nil {
		logs <- fmt.Sprintf("Main: Can't connect to rabbitmq on addr - %s", config.RabbitMQUrl)
	} else {
		// TODO: Remove after complete rabbitmq files.
		// TODO: Think about call to another nodes about starting??
		logs <- fmt.Sprintf("Main: Success")
		trainsRequest := rabbitmq.NewRequestQueue(&server.Chanel,
			"Send_all_trains",
			"",
			false,
			false,
			false,
			false,
			nil)

		trainsResponse := rabbitmq.NewResponseQueue(&server.Chanel,
			"Get_all_trains",
			"",
			false,
			false,
			false,
			false,
			nil)
		go server.Serve(trainsRequest, trainsResponse)
		time.Sleep(20 * time.Second)
		msg := rabbitmq.MessageRabbitMQ{
			ID: 1,
			Event: "Save_one_train",
			Data: Answer{
				TrainID: "313030d0add0a95fd09cd09ed0a1d09ad092d09020d0afd0a05fd092d09bd090d094d098d092d09ed0a1d0a25f30312e30322e323031395f30312e30322e32303139d41d8cd98f00b204e9800998ecf8427e",
				MainRoute: "МОСКВА ЯР - МОСКВА ЯР",
				Segment: "МОСКВА ЯРОСЛАВСКАЯ (ЯРОСЛАВСКИЙ ВОКЗАЛ) - ЯРОСЛАВЛЬ (ЯРОСЛАВЛЬ-МОСКОВСКИЙ)",
				StartDate: "25.01.2019_00:35",
			},
		}
		time.Sleep(time.Second)
		data, _ := json.Marshal(msg)
		err := trainsResponse.Send(data)
		if err != nil {
			logs <- fmt.Sprintf("Main: Error in test send - %s", err.Error())
		}
	}
	time.Sleep(5 * time.Minute)
}
