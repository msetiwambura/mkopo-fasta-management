package tests

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"usrmanagement/configs"
	"usrmanagement/models"
)

func TestRegister(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Setup a test router
	r := gin.Default()
	r.POST("/register")

	// Create a test user payload
	userPayload := models.User{
		FirstName: "Mseti",
		LastName:  "Nyasambo",
		Email:     "rafamceti@gmail.com",
		Password:  "Dmtml1",
		RoleID:    1,
	}

	// Marshal the user payload to JSON
	userJSON, _ := json.Marshal(userPayload)

	// Create a new HTTP request
	req, _ := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(userJSON))
	req.Header.Set("Content-Type", "application/json")

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()

	// Perform the request
	r.ServeHTTP(rr, req)

	// Check the status code
	assert.Equal(t, http.StatusCreated, rr.Code)

	// Check the response body
	var response map[string]interface{}
	err := json.Unmarshal(rr.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "User successfully created", response["ResponseBody"].(map[string]interface{})["Message"])
	assert.Equal(t, "User creation", response["ResponseBody"].(map[string]interface{})["Description"])
}

func init() {
	// Setup the test database connection
	models.MigrateDB()
	// Create the User table
	configs.DB.AutoMigrate(&models.User{})
}
