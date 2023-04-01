package controllers

import (
	"context"
	"jwt-gin/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type UserController struct {
	UserRepository *models.UserRepository
}

func (uc *UserController) Signup(c *gin.Context) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if ctx.Err() == context.DeadlineExceeded {
		c.JSON(http.StatusRequestTimeout, gin.H{"error": "Request timeout"})
		return
	}

	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if user.Name == "" || user.Email == "" || user.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Name, email, and password are required"})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}
	user.Password = string(hashedPassword)

	if err := uc.UserRepository.CreateUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, user)
}

func (uc *UserController) Login(c *gin.Context) {
	
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if ctx.Err() == context.DeadlineExceeded {
		c.JSON(http.StatusRequestTimeout, gin.H{"error": "Request timeout"})
		return
	}
	
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if user.Email == "" || user.Password == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Email and password are required"})
		return
	}

	existingUser, err := uc.UserRepository.GetUserByEmail(user.Email)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(user.Password)); err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	c.SetCookie("user_id", existingUser.Email, 3600, "/", "", false, true)
	c.String(http.StatusOK, "Login successful")
}

func (uc *UserController) Logout(c *gin.Context) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if ctx.Err() == context.DeadlineExceeded {
		c.JSON(http.StatusRequestTimeout, gin.H{"error": "Request timeout"})
		return
	}

	c.SetCookie("user_id", "", -1, "/", "", false, true)
	c.String(http.StatusOK, "Logout successful")
}

func (uc *UserController) GetUser(c *gin.Context) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if ctx.Err() == context.DeadlineExceeded {
		c.JSON(http.StatusRequestTimeout, gin.H{"error": "Request timeout"})
		return
	}

	userId := c.Param("id")
	user, err := uc.UserRepository.GetUserById(userId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to find user"})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (uc *UserController) GetUsers(c *gin.Context) {

	//  NÃO TA RETORNANDO TODOS OS USUERS CONFERIR DEPOIS
	
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if ctx.Err() == context.DeadlineExceeded {
		c.JSON(http.StatusRequestTimeout, gin.H{"error": "Request timeout"})
		return
	}

	users, err := uc.UserRepository.GetUsers()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to get users"})
		return
	}

	c.JSON(http.StatusOK, users)
}

func (uc *UserController) GetLoggedInUser(c *gin.Context) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if ctx.Err() == context.DeadlineExceeded {
		c.JSON(http.StatusRequestTimeout, gin.H{"error": "Request timeout"})
		return
	}

	userId, err := c.Cookie("user_id")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Failed to get user from session"})
		return
	}

	user, err := uc.UserRepository.GetUserByEmail(userId)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Failed to get user from session"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (uc *UserController) UpdateUser(c *gin.Context) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if ctx.Err() == context.DeadlineExceeded {
		c.JSON(http.StatusRequestTimeout, gin.H{"error": "Request timeout"})
		return
	}

	type updateUserInput struct {
		Name  string `json:"name"`
		Email string `json:"email"`
		Phone string `json:"phone"`
	}

	id := c.Param("id")

	var user models.User
	if err := uc.UserRepository.DB.Find(&user, id).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	var input updateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if input.Email != user.Email {
		if err := uc.UserRepository.DB.Where("email = ?", input.Email).First(&models.User{}).Error; err == nil {
			c.AbortWithStatus(http.StatusConflict)
			return
		}
	}

	user.Name = input.Name
	user.Email = input.Email
	user.Phone = input.Phone

	if err := uc.UserRepository.DB.Model(&user).Update("name", user.Name).Update("email", user.Email).Update("phone", user.Phone).Error; err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, user)
}

func (uc *UserController) DeleteUser(c *gin.Context) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()	
	if ctx.Err() == context.DeadlineExceeded {
		c.JSON(http.StatusRequestTimeout, gin.H{"error": "Request timeout"})
		return
	}

	userId := c.Param("id")

	user, err := uc.UserRepository.GetUserById(userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to find user"})
		return
	}

	if err := uc.UserRepository.DeleteUser(user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to delete user"})
		return
	}

	c.String(http.StatusOK, "User deleted successfully")
}
