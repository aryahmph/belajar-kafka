package main

import (
	"belajar_kafka/handlers"
	"belajar_kafka/pkg"
	"belajar_kafka/repositories"
	"belajar_kafka/services"
	"fmt"
	"github.com/Shopify/sarama"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
)

func main() {

	producer, err := sarama.NewSyncProducer([]string{"localhost:9092"}, nil)
	pkg.PanicIfError(err)

	consumer, err := sarama.NewConsumer([]string{"localhost:9092"}, nil)
	pkg.PanicIfError(err)

	db := pkg.NewDatabase()

	studentRepository := repositories.NewStudentRepositoryImpl(db)
	studentService := services.NewStudentServiceImpl(consumer, studentRepository)
	_ = handlers.NewStudentHandler(producer)

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "Hello WOrld")
	})

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	studentService.Subscribe()
	err = server.ListenAndServe()
	pkg.PanicIfError(err)
}
