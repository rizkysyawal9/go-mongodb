package usecase

import (
	"go-mongo/model"
	"go-mongo/repository"
)

type IStudentUseCase interface {
	FindAllStudents() ([]*model.Student, error)
	FindStudentByName(name string) (*model.Student, error)
	RegisterStudent(student model.Student) (*model.Student, error)
	FindStudentsWithPagination(skip int, limit int) ([]*model.Student, error)
}

type StudentUseCase struct {
	repo repository.IStudentRepository
}

func NewStudentUseCase(repo repository.IStudentRepository) *StudentUseCase {
	return &StudentUseCase{repo: repo}
}

func (s *StudentUseCase) FindAllStudents() ([]*model.Student, error) {
	return s.repo.GetAll()
}

func (s *StudentUseCase) FindStudentByName(name string) (*model.Student, error) {
	return s.repo.GetOneByUsername(name)
}

func (s *StudentUseCase) RegisterStudent(student model.Student) (*model.Student, error) {
	return s.repo.CreateOne(student)
}

func (s *StudentUseCase) FindStudentsWithPagination(skip int, limit int) ([]*model.Student, error) {
	return s.repo.GetWithPage(skip, limit)
}
