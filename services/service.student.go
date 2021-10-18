package services

import (
	"belajar_kafka/models"
	"belajar_kafka/pkg"
	"belajar_kafka/repositories"
	"encoding/json"
	"github.com/Shopify/sarama"
	"log"
	"os"
	"os/signal"
)

type StudentService interface {
	Subscribe()
	Create(payload []byte)
}

type StudentServiceImpl struct {
	Consumer          sarama.Consumer
	StudentRepository repositories.StudentRepository
}

func NewStudentServiceImpl(consumer sarama.Consumer, studentRepository repositories.StudentRepository) *StudentServiceImpl {
	return &StudentServiceImpl{Consumer: consumer, StudentRepository: studentRepository}
}

func (service *StudentServiceImpl) Subscribe() {
	partitionConsumer, err := service.Consumer.ConsumePartition("create-student", 0, sarama.OffsetNewest)
	pkg.PanicIfError(err)

	defer func() {
		err := partitionConsumer.Close()
		pkg.PanicIfError(err)
	}()

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

SubscribeLoop:
	for {
		select {
		case msg := <-partitionConsumer.Messages():
			service.Create(msg.Value)
			log.Printf("Consumed message offset %d\n", msg.Offset)
		case <-signals:
			break SubscribeLoop
		}
	}
}

func (service *StudentServiceImpl) Create(request []byte) {
	var payload models.PayloadStudentCreate
	err := json.Unmarshal(request, &payload)
	pkg.PanicIfError(err)

	service.StudentRepository.Save(models.ModelStudent{
		Name:  payload.Name,
		Email: payload.Email,
	})
}
