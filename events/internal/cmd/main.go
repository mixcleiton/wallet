package cmd

import (
	"database/sql"
	"fmt"

	"br.com.cleiton/events/internal/adapters/input/controller"
	"br.com.cleiton/events/internal/adapters/input/kafkamessage"
	"br.com.cleiton/events/internal/adapters/output/database"
	"br.com.cleiton/events/internal/config"
	"br.com.cleiton/events/internal/domain/usecases"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func StartEvents() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	databaseConfig, err := config.LoadDatabaseConfig("./config/database.yml")
	if err != nil {
		panic(err)
	}

	kafkaConfig, err := config.LoadKafkaConfig("./config/kafka.yml")
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
	createEventUC := usecases.NewCreateEventUC(&eventDatabase, kafkaProducer)
	eventController := controller.NewEventController(&createEventUC)

	processEventUC := usecases.NewProcessEventUC(&eventDatabase, &walletDatabase)

	kafkaConsumer := kafkamessage.NewKafkaConsumer(kafkaConfig, &processEventUC)
	kafkaConsumer.LoadReadMessages()

	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.POST("/event", eventController.CreateEvent)

	e.Logger.Fatal(e.Start(":8081"))
}

func configDB(databaseConfig config.DatabaseConfig) string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		databaseConfig.Host, databaseConfig.Port, databaseConfig.User, databaseConfig.Password, databaseConfig.DBName)
}
