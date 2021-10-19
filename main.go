package main

import (
	"belajar_kafka/handlers"
	"belajar_kafka/pkg"
	"belajar_kafka/repositories"
	"belajar_kafka/services"
	"fmt"
	"github.com/Shopify/sarama"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
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
	studentHandler := handlers.NewStudentHandler(producer)

	go studentService.Subscribe()

	router := httprouter.New()
	router.GET("/", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		fmt.Fprintf(writer, "Hello World")
	})
	router.POST("/", studentHandler.Create)

	err = http.ListenAndServe(":3000", router)
	pkg.PanicIfError(err)
}
