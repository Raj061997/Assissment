package test

import (
	"bytes"
	"encoding/json"
	"example/controller"
	"example/mocks" // Import generated mocks

	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateBlogPost(t *testing.T) {
	// Create a new mock service
	mockService := new(mocks.BlogService)

	// Mock the Create method to return nil (success)
	mockService.On("Create", mock.Anything).Return(nil)

	// Initialize the controller with the mock service
	blogController := controller.BlogController{BlogService: mockService}

	app := fiber.New()
	app.Post("/api/blog-post", blogController.CreatePost)

	postData := map[string]string{
		"title":       "Test Blog",
		"description": "description",
		"body":        "This is the body of the test blog post",
	}
	jsonData, _ := json.Marshal(postData)

	req := httptest.NewRequest("POST", "/api/blog-post", bytes.NewReader(jsonData))
	req.Header.Set("Content-Type", "application/json")
	resp, _ := app.Test(req)

	assert.Equal(t, http.StatusCreated, resp.StatusCode)

	// Verify that the mock expectations were met
	mockService.AssertExpectations(t)
}
