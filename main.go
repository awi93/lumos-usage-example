package main

import (
	"dimall.id/standard-template/bootstrap"
	handler "dimall.id/standard-template/event"
	web "dimall.id/standard-template/http"
	"github.com/dimall-id/lumos/config"
	"github.com/dimall-id/lumos/event"
	"github.com/dimall-id/lumos/http"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
	"log"
)

func main() {
	err := config.InitConfig("config",".")
	if err != nil {log.Fatal(err)}
	err = bootstrap.InitDB()
	if err != nil {log.Fatal(err)}

	go func() {
		handler.SetupHandler()
		log.Fatal(event.StartConsumer(&kafka.ConfigMap{
			"bootstrap.servers" : config.GetString("kafka.servers"),
			"client.id" : config.GetString("service.name"),
			"group.id" : config.GetString("service.name"),
			"auto.offset.reset": config.GetString("kafka.auto_offset_reset"),
		}))
	}()

	go func() {
		prodKafkaConfig := kafka.ConfigMap{
			"bootstrap.servers" : config.GetString("kafka.servers"),
			"client.id" : config.GetString("service.name"),
		}

		datasourceConfig := event.DatasourceConfig{
			Host: config.GetString("db.host"),
			User: config.GetString("db.username"),
			Password: config.GetString("db.password"),
			Database: config.GetString("db.database"),
			Port: config.GetString("db.port"),
			SslMode: config.GetString("db.sslmode"),
		}

		log.Fatal(event.StartProducer(event.Config{
			KafkaConfig: prodKafkaConfig,
			DatasourceConfig: datasourceConfig,
			PoolDuration: config.GetDuration("producer.pool_duration"),
		}))
	}()

	web.Routes()
	log.Fatal(http.StartHttpServer(":8080"))
}