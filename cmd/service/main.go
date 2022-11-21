package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"messenger/internal/handlers/messenger_test"
	"messenger/internal/repository"
	"messenger/internal/rpc"
)

func main() {
	fmt.Println("starting messenger up")
	f, err := os.OpenFile("./messenger.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	logger := log.New(f, "[messenger]", log.LstdFlags)

	rand.Seed(time.Now().UnixNano()) // инициируем Seed рандома для функции requestID
	flag.Parse()

	ctx := context.Background()

	mongoDB, err := mongo.NewClient(
		options.Client().ApplyURI("mongodb://127.0.0.1:27017").SetAuth(
			options.Credential{
				Username: "",
				Password: "",
			},
		),
	)
	err = mongoDB.Connect(ctx)
	if err != nil {
		logger.Fatal(err)
		return
	}

	repo := repository.New(mongoDB, logger)

	// стартуем rpc сервер
	rpcServer := rpc.NewServer()

	// проверка статуса работы с монгой
	err = rpcServer.Register("messenger_test", messenger_test.New(repo))

	// err = rpcServer.Register("chats", info.New(repository)) // список чатов
	// err = rpcServer.Register("chat/{chatId}/send", info.New(repository)) написать сообщение в чат
	// err = rpcServer.Register("chat/{chatId}", info.New(repository)) // прочитать последние 20 сообщений из чата

	http.Handle("/", rpcServer)

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", 8080),
		Handler:        http.DefaultServeMux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: http.DefaultMaxHeaderBytes,
	}

	go func() {
		logger.Println("Start server on port 8080")
		err := s.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			logger.Fatalf("failed to listen and serve: %v.", err.Error())
		}
	}()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	done := make(chan bool, 1)

	fmt.Println("awaiting signal")
	<-done
	fmt.Println("exiting")
}
