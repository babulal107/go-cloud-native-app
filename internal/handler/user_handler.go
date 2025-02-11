package handler

import (
	"fmt"
	"github.com/babulal107/go-cloud-native-app/internal/model"
	"github.com/babulal107/go-cloud-native-app/internal/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

type UserHandler struct {
	UserSvc service.UserSvc
}

func NewUserHandler(userSvc service.UserSvc) *UserHandler {
	return &UserHandler{
		UserSvc: userSvc,
	}
}

// Register TODO : directly return gin HandlerFunc
func (u *UserHandler) Register() gin.HandlerFunc {
	return func(c *gin.Context) {

		fmt.Println("Register called")

		c.JSON(200, gin.H{
			"status":  http.StatusCreated,
			"message": "success",
		})
		return
	}
}

// PostUser : Add users
func (u *UserHandler) PostUser(c *gin.Context) {
	fmt.Println("PostUser")

	req := model.UserRequest{}

	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		log.Printf("invalid body request data error : %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "invalid body request data",
		})
		return
	}

	userId, err := u.UserSvc.AddUser(c.Request.Context(), req)
	if err != nil {
		fmt.Printf("add user error : %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "add user error",
		})
		return

	}
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "success",
		"userId":  userId,
	})
}

// GetUsers : Get list of users
func (u *UserHandler) GetUsers(c *gin.Context) {
	fmt.Println("GetUsers")

	data, err := u.UserSvc.GetUsers(c.Request.Context())
	if err != nil {
		log.Printf("get users error : %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "get users error",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "success",
		"data":    data,
	})
}

// GetUser : Get single user details
func (u *UserHandler) GetUser(c *gin.Context) {

	userId := c.Param("id")
	if userId == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "invalid user id",
		})
		return
	}

	id, err := strconv.Atoi(userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "invalid user id",
		})
		return
	}

	data, err := u.UserSvc.GetUser(c.Request.Context(), id)
	if err != nil {
		log.Printf("get users error : %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "get users error",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "success",
		"data":    data,
	})

}
