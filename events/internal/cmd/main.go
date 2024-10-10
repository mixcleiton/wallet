package cmd

import (
	"database/sql"
	"fmt"
	"log"

	"br.com.cleiton/events/internal/adapters/input/controller"
	"br.com.cleiton/events/internal/adapters/input/kafkamessage"
	"br.com.cleiton/events/internal/adapters/output/database"
	"br.com.cleiton/events/internal/config"
	"br.com.cleiton/events/internal/domain/usecases"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func StartEvents() {
	e := echo.New()
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	databaseConfig, err := config.LoadDatabaseConfig("database.yaml")
	if err != nil {
		panic(err)
	}

	kafkaConfig, err := config.LoadKafkaConfig("kafka.yaml")
	if err != nil {
		panic(err)
	}

	db, err := sql.Open("postgres", configDB(databaseConfig))
	if err != nil {
		panic(err)
	}

	kafkaProducer := kafkamessage.NewKafkaProducer(kafkaConfig)

	walletDatabase := database.NewWalletDatabase(db)
	eventDatabase := database.NewEventDatabase(db)
	createEventUC := usecases.NewCreateEventUC(&eventDatabase, kafkaProducer, &walletDatabase)
	eventController := controller.NewEventController(&createEventUC)

	processEventUC := usecases.NewProcessEventUC(&eventDatabase, &walletDatabase, kafkaProducer)

	kafkaConsumer := kafkamessage.NewKafkaConsumer(kafkaConfig, &processEventUC)
	kafkaConsumer.LoadReadMessages()

	g := e.Group("/api/v1")

	g.POST("/event", eventController.CreateEvent)

	e.Logger.Fatal(e.Start(":8081"))
}

func configDB(databaseConfig config.DatabaseConfig) string {
	urlDb := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		"localhost", 5432, "root", "wallettest", "postgres")

	log.Println(urlDb)
	return urlDb
}
