package repositories

import (
	"belajar_kafka/models"
	"belajar_kafka/pkg"
	"context"
	"database/sql"
)

type StudentRepository interface {
	Save(student models.ModelStudent)
}

type StudentRepositoryImpl struct {
	DB *sql.DB
}

func NewStudentRepositoryImpl(db *sql.DB) *StudentRepositoryImpl {
	return &StudentRepositoryImpl{DB: db}
}

func (repository *StudentRepositoryImpl) Save(student models.ModelStudent) {
	SQL := "INSERT INTO students(name, email) VALUES (?,?)"
	_, err := repository.DB.ExecContext(context.Background(), SQL, student.Name, student.Email)
	pkg.PanicIfError(err)
}
