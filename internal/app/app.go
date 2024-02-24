package app

import (
	"database/sql"
	"fmt"
	"log"
	"math"
	"time"

	"github.com/gin-gonic/gin"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/zura-t/go_delivery_system-shops/config"
	v1 "github.com/zura-t/go_delivery_system-shops/internal/controller/http/v1"
	"github.com/zura-t/go_delivery_system-shops/internal/usecase"
	db "github.com/zura-t/go_delivery_system-shops/pkg/db/sqlc"
	"github.com/zura-t/go_delivery_system-shops/pkg/logger"
	"github.com/zura-t/go_delivery_system-shops/rmq"
	// "os"
)

func Run(config *config.Config) {
	dbconn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("can't connect to db:", err)
	}

	store := db.NewStore(dbconn)
	if err != nil {
		log.Fatal("can't create store:", err)
	}

	// rabbitConn, err := connectRabbitmq()
	// if err != nil {
	// 	log.Println(err)
	// 	os.Exit(1)
	// }
	// defer rabbitConn.Close()

	l := logger.New(config.LogLevel)

	usersUseCase := usecase.New(store, config)

	runGinServer(l, config, usersUseCase)

	// channel, consumer, err := setupRabbitmq(rabbitConn, server)
	// if err != nil {
	// 	panic(err)
	// }
	// defer channel.Close()

	// err = consumer.Listen([]string{})
	// if err != nil {
	// 	log.Println(err)
	// }
}

func runGinServer(l *logger.Logger, config *config.Config, shopUsecase *usecase.ShopUseCase) {
	handler := gin.New()
	v1.NewRouter(handler, l, shopUsecase)
	handler.Run(config.HttpPort)
}

func connectRabbitmq() (*amqp.Connection, error) {
	var counts int64
	var backOff = 1 * time.Second
	var connection *amqp.Connection
	for {
		c, err := amqp.Dial("amqp://admin:admin@rabbitmq:5672")

		if err != nil {
			fmt.Println("rabbitmq not yet ready")
			counts++
		} else {
			log.Println("Connected to rabbitmq.")

			connection = c
			break
		}

		if counts > 5 {
			fmt.Println(err)
			return nil, err
		}

		backOff = time.Duration(math.Pow(float64(counts), 2)) * time.Second
		log.Println("back off")
		time.Sleep(backOff)
		continue
	}

	return connection, nil
}

func setupRabbitmq(rabbitConn *amqp.Connection) (*amqp.Channel, *rmq.Consumer, error) {
	channel, err := rabbitConn.Channel()
	if err != nil {
		log.Fatal("can't create rabbitmq consumer", err)
		return nil, nil, err
	}

	consumer, err := rmq.NewConsumer(rabbitConn, channel)
	if err != nil {
		return nil, nil, err
	}

	return channel, &consumer, nil
}
