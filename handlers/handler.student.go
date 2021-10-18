package handlers

import (
	"belajar_kafka/models"
	"belajar_kafka/pkg"
	"encoding/json"
	"github.com/Shopify/sarama"
	"net/http"
	"strings"
)

type StudentHandler struct {
	Producer sarama.SyncProducer
}

func NewStudentHandler(producer sarama.SyncProducer) *StudentHandler {
	return &StudentHandler{Producer: producer}
}

func (handler *StudentHandler) Create(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	pkg.PanicIfError(err)

	name := request.PostForm.Get("name")
	email := request.PostForm.Get("email")
	payload := models.PayloadStudentCreate{
		Name:  strings.Trim(name, " "),
		Email: strings.Trim(email, " "),
	}
	marshal, err := json.Marshal(payload)
	pkg.PanicIfError(err)

	message := &sarama.ProducerMessage{
		Topic: "create-student",
		Value: sarama.StringEncoder(marshal),
	}
	_, _, err = handler.Producer.SendMessage(message)
	pkg.PanicIfError(err)
}
