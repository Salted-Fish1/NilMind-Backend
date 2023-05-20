package signupin

import (
	"context"
	"fmt"
	"golesson/model"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/gomail.v2"
)

var SecretKey = []byte("NilMind-Secret-Key")

type User struct {
	// Id               primitive.ObjectID `bson:"_id"`
	// User_id          primitive.ObjectID `bson:"user_id"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required"`
}

// 哈希密码
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedPassword), err
}

// 验证密码
func VerifyPassword(password string, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func generateRandomCode() int {
	min := 100000
	max := 999999
	randomCode := rand.Intn(max-min+1) + min
	return randomCode
}

var code int
var timer *time.Timer

var user User

func SignUp(c *gin.Context) {
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	hashedPassword, err := HashPassword(user.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user.Password = hashedPassword
	collection := model.DB.Database("NilMind-backend").Collection("users")
	insertResult, err := collection.InsertOne(context.Background(), user)
	fmt.Println(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	m := gomail.NewMessage()
	m.SetHeader("From", "just_for_register@qq.com")
	m.SetHeader("To", user.Email)
	m.SetHeader("Subject", "NilMind 注册通知")
	code = generateRandomCode()
	m.SetBody("text/plain", fmt.Sprintf("注册码: %v, 请在 5 分钟内验证注册码", code))

	// 发送邮件
	d := gomail.NewDialer("smtp.qq.com", 587, "just_for_register@qq.com", "otyahzrgdcyabdfh")
	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err.Error())
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Email sent successfully"})

	timer = time.AfterFunc(5*time.Minute, func() {
		fmt.Println(insertResult)
		filter := bson.M{"username": user.Username}
		collection.DeleteOne(context.Background(), filter)
		fmt.Println("5 Minute passed")
	})
}

type Code struct {
	Code int `json:"code" binding:"required"`
}

func VerifyCode(c *gin.Context) {
	var nCode Code
	if err := c.ShouldBindJSON(&nCode); err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(nCode.Code)
	fmt.Println(code)
	if nCode.Code != code {
		c.JSON(http.StatusBadRequest, gin.H{"error": "wrong code"})
		return
	}
	timer.Stop()
	c.JSON(http.StatusOK, gin.H{"message": "Signed Up"})
}

type SignInUser struct {
	// Id               primitive.ObjectID `bson:"_id"`
	// User_id          primitive.ObjectID `bson:"user_id"`
	// Username string `json:"username"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required"`
}

func SignIn(c *gin.Context) {
	var user SignInUser
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// c.JSON(http.StatusOK, gin.H{"message": "Signed In"})
	// filter := bson.M{"email": user.Email}
	collection := model.DB.Database("NilMind-backend").Collection("users")
	filter := bson.M{"email": user.Email}
	findResult := collection.FindOne(context.Background(), filter)
	fmt.Println(findResult)
	fmt.Println(user)
	fmt.Println("--------------")

	var doc User
	err := findResult.Decode(&doc)
	fmt.Println(doc)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := VerifyPassword(user.Password, doc.Password); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		// "foo": "bar",
		// "nbf": time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
		"email":     user.Email,
		"timeStamp": time.Now(),
	})
	fmt.Println(token)
	fmt.Println("---------")

	tokenString, err := token.SignedString(SecretKey)

	fmt.Println(tokenString, err)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"access_token": tokenString})
}

func SignTest(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"access_token": "ok"})

	var token struct {
		token string
	}
	// bearerToken := c.GetHeader("Authorization")
	if err := c.ShouldBindJSON(&token); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
}
