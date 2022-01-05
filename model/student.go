package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Student struct {
	Id       primitive.ObjectID `bson:"_id"`
	Name     string
	Gender   string
	Age      int
	JoinDate time.Time `json:"joinDate"`
	IdCard   string    `bson:"idCard"`
	Senior   bool
}
