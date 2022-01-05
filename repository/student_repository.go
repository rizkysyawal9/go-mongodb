package repository

import (
	"fmt"
	"go-mongo/db"
	"go-mongo/model"
	"go-mongo/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	options2 "go.mongodb.org/mongo-driver/mongo/options"
)

type IStudentRepository interface {
	GetAll() ([]*model.Student, error)
	GetOneByUsername(name string) (*model.Student, error)
	CreateOne(student model.Student) (*model.Student, error)
	GetWithPage(skip int, limit int) ([]*model.Student, error)
}

type StudentRepository struct {
	repo *mongo.Collection
}

func (s *StudentRepository) GetWithPage(skip int, limit int) ([]*model.Student, error) {
	ctx, cancel := utils.InitContext()
	defer cancel()
	options := options2.Find()
	options.SetSkip(int64(skip))
	options.SetLimit(int64(limit))
	cursor, err := s.repo.Find(ctx, bson.D{}, options)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	students := []*model.Student{}
	for cursor.Next(ctx) {
		var student model.Student
		err = cursor.Decode(&student)
		if err != nil {
			return nil, err
		}
		students = append(students, &student)
	}
	return students, nil
}

func (s *StudentRepository) GetAll() ([]*model.Student, error) {
	students := []*model.Student{}
	ctx, cancel := utils.InitContext()
	defer cancel()

	cursor, err := s.repo.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	for cursor.Next(ctx) {
		var student model.Student
		err = cursor.Decode(&student)
		if err != nil {
			return nil, err
		}
		students = append(students, &student)
	}
	return students, nil
}

func (s *StudentRepository) GetOneByUsername(name string) (*model.Student, error) {
	filterUsername := bson.D{
		{"name", name},
	}
	ctx, cancel := utils.InitContext()
	defer cancel()
	student := model.Student{}
	err := s.repo.FindOne(ctx, filterUsername).Decode(&student)
	if err != nil {
		return nil, err
	}
	return &student, nil
}

func (s *StudentRepository) CreateOne(student model.Student) (*model.Student, error) {
	ctx, cancel := utils.InitContext()
	defer cancel()
	newId, err := s.repo.InsertOne(ctx, student)
	if err != nil {
		return nil, err
	}
	student.Id = newId.InsertedID.(primitive.ObjectID)
	return &student, nil
}

func NewStudentRepository(resources *db.Resource) IStudentRepository {
	studentCollection := resources.DB.Collection("students")
	studentRepository := &StudentRepository{repo: studentCollection}
	return studentRepository
}
