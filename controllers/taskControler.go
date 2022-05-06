package controllers

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/BrainStation-23/dsdoc-backend/database"
	model "github.com/BrainStation-23/dsdoc-backend/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var TaskCollection *mongo.Collection = database.OpenCollection(database.Client, "task")

func CreateTaks(c *gin.Context) {
	// user_name,_ := c.Get("email")
	// c.JSON(http.StatusOK, gin.H{"name":user_name})

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	var task model.Task

	if err := c.BindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
		return 
	}

	validationErr := validate.Struct(task)
	if validationErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":validationErr.Error()})
		return
	}


	user_id,_ := c.Get("uid")
	fmt.Println(reflect.TypeOf(user_id))
	user_id_OK, _ := user_id.(string)
	task.User_ID = user_id_OK
	task.ID = primitive.NewObjectID()
	task.Created_at,_ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	task.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	

	resultInsertionNumber, insertErr := TaskCollection.InsertOne(ctx, task)
	if insertErr !=nil {
		msg := fmt.Sprintf("Task item was not created")
		c.JSON(http.StatusInternalServerError, gin.H{"error":msg})
		return
	}

	defer cancel()
	c.JSON(http.StatusOK, resultInsertionNumber)
}

func UpdateTask(c *gin.Context){
	
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	var task model.Task
	
	Updated_id,_ := primitive.ObjectIDFromHex(c.Param("id"))
	User_Id_1,_ := c.Get("uid")
	User_id := User_Id_1.(string)
	
	fmt.Println(User_id)

	err := TaskCollection.FindOne(ctx, bson.M{"_id":Updated_id,"user_id":User_id}).Decode(&task)
	defer cancel()
	if  err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": task})
		return
	}

	var input model.UpdateTask
	if err := c.ShouldBindJSON(&input); err != nil {
	  c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	  return
	}


	editedTask, err := TaskCollection.UpdateOne(
		ctx,
		bson.M{"_id": Updated_id, "user_id":User_id},
		bson.D{
			{"$set", input},
		},
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{ "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, editedTask)

  
}

func GetTask(c *gin.Context) {

}