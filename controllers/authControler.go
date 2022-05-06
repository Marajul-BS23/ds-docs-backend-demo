package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/BrainStation-23/dsdoc-backend/database"
	helper "github.com/BrainStation-23/dsdoc-backend/helpers"
	model "github.com/BrainStation-23/dsdoc-backend/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)
var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")
var validate = validator.New()


func HashPassword(password string) string{
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err!=nil{
		log.Panic(err)
	}
	return string(bytes)
}
func VerifyPassword(userPassword string, providedPassword string)(bool, string){
	err := bcrypt.CompareHashAndPassword([]byte(providedPassword), []byte(userPassword))
	check := true
	msg := ""

	if err!= nil {
		msg = fmt.Sprintf("email of password is incorrect")
		check=false
	}
	return check, msg
}

func Login(c *gin.Context) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	var user model.User
	var foundUser model.User

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
		return 
	}

	if user.Email==nil || user.Password==nil{
		c.JSON(http.StatusBadRequest, gin.H{"error":"Email/password required"})
		return
	}

	err := userCollection.FindOne(ctx, bson.M{"email":user.Email}).Decode(&foundUser)
	defer cancel()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error":"email or password is incorrect"})
		return
	}

	passwordIsValid, msg := VerifyPassword(*user.Password, *foundUser.Password)
	defer cancel()
	if passwordIsValid != true{
		c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
		return
	}

	if foundUser.Email == nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error":"user not found"})
	}
	token, refreshToken, _ := helper.GenerateAllTokens(*foundUser.Email, *foundUser.First_name, *foundUser.Last_name, *foundUser.User_type, foundUser.User_id)
	helper.UpdateAllTokens(token, refreshToken, foundUser.User_id)
	err = userCollection.FindOne(ctx, bson.M{"user_id":foundUser.User_id}).Decode(&foundUser)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.SetCookie("token",*foundUser.Token, 60*60*24, "/", "localhost:5000", false, false)
	
	c.JSON(http.StatusOK, foundUser)
	
	

	
	
		
}

func Signup(c *gin.Context) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	var user model.User

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	validationErr := validate.Struct(user)
	if validationErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":validationErr.Error()})
		for _, err := range validationErr.(validator.ValidationErrors) {

			// fmt.Println(err.Namespace())
			fmt.Println(err.Field())
			fmt.Println(err.StructNamespace())
			fmt.Println(err.StructField())
			fmt.Println(err.Tag())
			fmt.Println(err.ActualTag())
			fmt.Println(err.Kind())
			fmt.Println(err.Type())
			fmt.Println(err.Value())
			fmt.Println(err.Param())
			fmt.Println()
		}
		return
	}

	_, err := userCollection.CountDocuments(ctx, bson.M{"email":user.Email})
	defer cancel()
	if err != nil {
		log.Panic(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error":"error occured while checking for the email"})
	}

	password := HashPassword(*user.Password)
	user.Password = &password

	count, err := userCollection.CountDocuments(ctx, bson.M{"phone":user.Phone})
	defer cancel()
	if err!= nil {
		log.Panic(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error":"error occured while checking for the phone number"})
	}

	if count >0{
		c.JSON(http.StatusInternalServerError, gin.H{"error":"this email or phone number already exists"})
	}

	user_type := "USER"
	user.User_type = &user_type
	
	user.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	user.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	user.ID = primitive.NewObjectID()
	user.User_id = user.ID.Hex()
	token, refreshToken, _ := helper.GenerateAllTokens(*user.Email, *user.First_name, *user.Last_name, *user.User_type, *&user.User_id)
	user.Token = &token
	user.Refresh_token = &refreshToken

	resultInsertionNumber, insertErr := userCollection.InsertOne(ctx, user)
	if insertErr !=nil {
		msg := fmt.Sprintf("User item was not created")
		c.JSON(http.StatusInternalServerError, gin.H{"error":msg})
		return
	}
	defer cancel()
	c.JSON(http.StatusOK, resultInsertionNumber)
}