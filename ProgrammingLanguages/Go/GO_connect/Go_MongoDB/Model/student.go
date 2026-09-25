package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Student struct {
	ID primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name  string `json:"name" bson:"name"`
	Email string `json:"email" bson:"email"`
	PhoneNum string `json:"phoneNum" bson:"phoneNum"`
	GPA float64 `json:"gpa" bson:"gpa"`
	CreatedAt time.Time `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt" bson:"updatedAt"`
}

// Struct Tạo ra đối tượng
type CreateStudentInput struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required,email"`
	PhoneNum string `json:"phoneNum" binding:"required"`
	GPA float64 `json:"gpa" binding:"min=0,max=4"`
}

// Struct Cập nhật đối tượng
type UpdateStudentInput struct {
	Name  string `json:"name" bson:"name,omitempty"`
	Email string `json:"email" bson:"email,omitempty"`
	PhoneNum string `json:"phoneNum" bson:"phoneNum,omitempty"`
	GPA float64 `json:"gpa" bson:"gpa,omitempty"`
}