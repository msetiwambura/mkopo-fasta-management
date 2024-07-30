package controllers

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strconv"
	"usrmanagement/configs"
	"usrmanagement/models"
	"usrmanagement/utils"
)

func Register(c *gin.Context) {
	var newUser models.User

	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password h"})
		return
	}
	newUser.Password = string(hashedPassword)

	if err := configs.DB.Create(&newUser).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}
	response := utils.CreateSuccessResponse("User successfully created", "User creation", []models.User{newUser})

	c.JSON(http.StatusCreated, response)
}

type ResUser struct {
	ID        uint   `json:"ID"`
	FirstName string `json:"FirstName"`
	LastName  string `json:"LastName"`
	Email     string `json:"Email"`
	Role      Role   `json:"Role"` // Detailed role object
}

type Role struct {
	RoleID uint   `json:"RoleID"`
	Name   string `json:"Name"`
}

func GetUsers(c *gin.Context) {
	var users []models.User
	var roles []models.Role

	// Fetch all users with preloaded roles
	if err := configs.DB.Preload("Role").Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, utils.CreateErrorResponse("Failed to retrieve users", err.Error()))
		return
	}

	// Fetch all roles and create a map for quick lookup
	if err := configs.DB.Find(&roles).Error; err != nil {
		c.JSON(http.StatusInternalServerError, utils.CreateErrorResponse("Failed to retrieve roles", err.Error()))
		return
	}

	// Create a map of roles for quick lookup
	roleMap := make(map[uint]models.Role)
	for _, role := range roles {
		roleMap[role.ID] = role
	}

	// Create response list with role details
	var res []ResUser
	for _, user := range users {
		role, ok := roleMap[user.RoleID] // Fetch the role details from the map
		if !ok {
			role = models.Role{} // Handle the case where role is not found
		}

		res = append(res, ResUser{
			ID:        user.ID,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Email:     user.Email,
			Role: Role{
				RoleID: role.ID,
				Name:   role.Name,
			},
		})
	}

	response := utils.CreateSuccessResponse("Success", "Request processed successfully", res)
	c.JSON(http.StatusOK, response)
}

func Login(c *gin.Context) {
	var input models.Login
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var user models.User
	//if err := configs.DB.Where("email = ?", input.Username).First(&user).Error; err != nil {
	//	c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
	//	return
	//}

	if err := configs.DB.Where("email = ?", input.Username).Preload("Role").First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}
	token, err := utils.GenerateToken(user.Email, strconv.Itoa(int(user.Role.ID)))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate JWT token"})
		return
	}
	resUser := models.ResUser{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		RoleID:    user.RoleID,
		Email:     user.Email,
	}

	res := utils.LoginResponse("Logged in successfully", "Ok", token, resUser)

	c.JSON(http.StatusOK, res)
}
