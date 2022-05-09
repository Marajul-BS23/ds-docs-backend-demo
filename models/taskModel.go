package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Task struct {
	ID            primitive.ObjectID    `bson:"_id"`
	Title    	 *string            	`json:"title" validate:"required,min=1"`
	Content       *string            	`json:"content" validate:"required"`
	Type_ID       *string            	`json:"type_id" validate:"required"`
	User_ID       string            	`json:"user_id"`
	Project_ID    *string            	`json:"project_id" validate:"required"`
	Status        *string            	`json:"status" validate:"required"`
	Catagory_ID   *string            	`json:"category_id" validate:"required"`
	Date          *string			 	`json:"date" validate:"required"`
	Created_at    time.Time          	`json:"created_at"`
	Updated_at    time.Time          	`json:"updated_at"`
}

type UpdateTask struct{
	Title    	  	*string            	`json:"title,omitempty"`
	Content       *string            	`json:"content,omitempty"`
	Type_ID       *string            	`json:"type_id,omitempty"`
	Project_ID    *string            	`json:"project_id,omitempty"`
	Status        *string            	`json:"status,omitempty"`
	Catagory_ID   *string            	`json:"category_id,omitempty"`
	Date          *string							`json:"date,omitempty"`
}