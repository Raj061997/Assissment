package controller

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"example/mocks"
	"example/models"

	"github.com/c2fo/testify/require"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetPosts(t *testing.T) {
	// Create a Fiber app
	app := fiber.New()

	// Create a mock service
	mockService := new(mocks.BlogService)

	// Create a BlogController with the mock service
	bc := &BlogController{service: mockService}

	// Register the handler
	app.Get("/blog-posts", bc.GetPosts)

	// Define test cases
	tests := []struct {
		description   string
		mockReturn    []models.BlogPost
		mockReturnErr error
		expectedCode  int
		mockCalled    bool
	}{
		{
			description: "success case - retrieved posts",
			mockReturn: []models.BlogPost{
				{ID: 1, Title: "First Blog", Description: "This is the first blog", Body: "Body content"},
				{ID: 2, Title: "Second Blog", Description: "This is the second blog", Body: "Body content"},
			},
			mockReturnErr: nil,
			expectedCode:  http.StatusOK,
			mockCalled:    true,
		},
		{
			description:   "failure case - unable to fetch posts",
			mockReturn:    nil,
			mockReturnErr: errors.New("unable to find blog"),
			expectedCode:  http.StatusNotFound,
			mockCalled:    true,
		},
	}

	// Iterate over test cases
	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			if test.mockCalled {
				// Mock the service method
				mockService.On("GetAll").Return(test.mockReturn, test.mockReturnErr).Once()
			}

			// Create test request
			req := httptest.NewRequest(http.MethodGet, "/blog-posts", nil)
			req.Header.Set("Content-Type", "application/json")

			// Execute the request and capture the response
			resp, err := app.Test(req)
			require.NoError(t, err) // Ensure no Fiber errors

			// Assert response status
			assert.Equalf(t, test.expectedCode, resp.StatusCode, test.description)

			// Assert the mock was called
			if test.mockCalled {
				mockService.AssertExpectations(t)
			}
		})
	}
}

func TestGetPost(t *testing.T) {
	// Create a Fiber app
	app := fiber.New()

	// Create a mock service
	mockService := new(mocks.BlogService)

	// Create a BlogController with the mock service
	bc := &BlogController{service: mockService}

	// Register the handler
	app.Get("/blog-post/:id", bc.GetPost)

	// Define test cases
	tests := []struct {
		description   string
		paramID       string
		mockReturn    *models.BlogPost
		mockReturnErr error
		expectedCode  int
		mockCalled    bool
	}{
		{
			description: "success case - retrieved post",
			paramID:     "1",
			mockReturn: &models.BlogPost{
				ID:          1,
				Title:       "Test Blog",
				Description: "This is a test blog",
				Body:        "Blog content",
			},
			mockReturnErr: nil,
			expectedCode:  http.StatusOK,
			mockCalled:    true,
		},
		{
			description:   "failure case - post not found",
			paramID:       "2",
			mockReturn:    nil,
			mockReturnErr: errors.New("Post not found"),
			expectedCode:  http.StatusNotFound,
			mockCalled:    true,
		},
		{
			description:   "failure case - invalid ID parameter",
			paramID:       "abc",
			mockReturn:    nil,
			mockReturnErr: nil, // Should not call the service
			expectedCode:  http.StatusBadRequest,
			mockCalled:    false,
		},
		{
			description:   "failure case - negative ID",
			paramID:       "-1",
			mockReturn:    nil,
			mockReturnErr: nil, // Should not call the service
			expectedCode:  http.StatusBadRequest,
			mockCalled:    false,
		},
	}

	// Iterate over test cases
	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			if test.mockCalled {
				// Mock only if the service is expected to be called
				mockService.On("GetByID", mock.AnythingOfType("uint")).
					Return(test.mockReturn, test.mockReturnErr).
					Once()
			}

			// Create test request
			req := httptest.NewRequest(http.MethodGet, "/blog-post/"+test.paramID, nil)
			req.Header.Set("Content-Type", "application/json")

			// Execute the request and capture the response
			resp, err := app.Test(req)
			require.NoError(t, err) // Ensure no Fiber errors

			// Assert response status
			assert.Equalf(t, test.expectedCode, resp.StatusCode, test.description)

			// Assert the mock was called only if expected
			if test.mockCalled {
				mockService.AssertExpectations(t)
			} else {
				mockService.AssertNotCalled(t, "GetByID")
			}
		})
	}
}

func TestUpdatePost(t *testing.T) {
	// Create a Fiber app
	app := fiber.New()

	// Create a mock service
	mockService := new(mocks.BlogService)

	// Create a BlogController with the mock service
	bc := &BlogController{service: mockService}

	// Register the handler
	app.Put("/blog-post/:id", bc.UpdatePost)
	ss := ""
	// Define test cases
	tests := []struct {
		description   string
		paramID       string
		requestBody   models.UpdateBlogRequest
		mockReturn    *models.BlogPost
		mockReturnErr error
		expectedCode  int
		mockCalled    bool
	}{
		{
			description: "success case - updated post",
			paramID:     "1",
			requestBody: models.UpdateBlogRequest{
				Title:       &ss,
				Description: &ss,
				Body:        &ss,
			},
			mockReturn: &models.BlogPost{
				ID:          1,
				Title:       "Updated Title",
				Description: "Updated Description",
				Body:        "Updated Body",
			},
			mockReturnErr: nil,
			expectedCode:  http.StatusOK,
			mockCalled:    true,
		},
		{
			description:   "failure case - invalid ID parameter",
			paramID:       "abc",
			requestBody:   models.UpdateBlogRequest{},
			mockReturn:    nil,
			mockReturnErr: nil, // Should not call the service
			expectedCode:  http.StatusBadRequest,
			mockCalled:    false,
		},
		{
			description:   "failure case - negative ID",
			paramID:       "-1",
			requestBody:   models.UpdateBlogRequest{},
			mockReturn:    nil,
			mockReturnErr: nil, // Should not call the service
			expectedCode:  http.StatusBadRequest,
			mockCalled:    false,
		},
		{
			description:   "failure case - invalid request body",
			paramID:       "1",
			requestBody:   models.UpdateBlogRequest{}, // Empty request
			mockReturn:    nil,
			mockReturnErr: nil, // Should not call the service
			expectedCode:  http.StatusBadRequest,
			mockCalled:    false,
		},
		{
			description: "failure case - unable to update post",
			paramID:     "1",
			requestBody: models.UpdateBlogRequest{
				Title:       &ss,
				Description: &ss,
				Body:        &ss,
			},
			mockReturn:    nil,
			mockReturnErr: errors.New("unable to update post"),
			expectedCode:  http.StatusInternalServerError,
			mockCalled:    true,
		},
	}

	// Iterate over test cases
	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			// Convert request body to JSON
			reqBytes, _ := json.Marshal(test.requestBody)

			if test.mockCalled {
				// Mock only if the service is expected to be called
				mockService.On("Update", mock.AnythingOfType("uint"), mock.AnythingOfType("*models.UpdateBlogRequest")).
					Return(test.mockReturn, test.mockReturnErr).
					Once()
			}

			// Handle invalid request body case (send empty body)
			if test.description == "failure case - invalid request body" {
				reqBytes = []byte{}
			}

			// Create test request
			req := httptest.NewRequest(http.MethodPut, "/blog-post/"+test.paramID, bytes.NewBuffer(reqBytes))
			req.Header.Set("Content-Type", "application/json")

			// Execute the request and capture the response
			resp, err := app.Test(req)
			require.NoError(t, err) // Ensure no Fiber errors

			// Assert response status
			assert.Equalf(t, test.expectedCode, resp.StatusCode, test.description)

			// Assert the mock was called only if expected
			if test.mockCalled {
				mockService.AssertExpectations(t)
			} else {
				mockService.AssertNotCalled(t, "Update")
			}
		})
	}
}

func TestDeletePost(t *testing.T) {
	// Create a Fiber app
	app := fiber.New()

	// Create a mock service
	mockService := new(mocks.BlogService)

	// Create a BlogController with the mock service
	bc := &BlogController{service: mockService}

	// Register the handler
	app.Delete("/blog-post/:id", bc.DeletePost)

	// Define test cases
	tests := []struct {
		description   string
		paramID       string
		mockGetReturn *models.BlogPost
		mockGetErr    error
		mockDelErr    error
		expectedCode  int
		mockGetCalled bool
		mockDelCalled bool
	}{
		{
			description:   "success case - post deleted",
			paramID:       "1",
			mockGetReturn: &models.BlogPost{ID: 1, Title: "Test", Description: "Desc", Body: "Body"},
			mockGetErr:    nil,
			mockDelErr:    nil,
			expectedCode:  http.StatusNoContent, // 204 No Content
			mockGetCalled: true,
			mockDelCalled: true,
		},
		{
			description:   "failure case - invalid ID parameter",
			paramID:       "abc",
			mockGetReturn: nil,
			mockGetErr:    nil, // Should not call the service
			mockDelErr:    nil,
			expectedCode:  http.StatusBadRequest,
			mockGetCalled: false,
			mockDelCalled: false,
		},
		{
			description:   "failure case - negative ID",
			paramID:       "-1",
			mockGetReturn: nil,
			mockGetErr:    nil, // Should not call the service
			mockDelErr:    nil,
			expectedCode:  http.StatusBadRequest,
			mockGetCalled: false,
			mockDelCalled: false,
		},
		{
			description:   "failure case - post not found",
			paramID:       "1",
			mockGetReturn: nil,
			mockGetErr:    errors.New("post not found"),
			mockDelErr:    nil,
			expectedCode:  http.StatusNotFound,
			mockGetCalled: true,
			mockDelCalled: false,
		},
		{
			description:   "failure case - unable to delete post",
			paramID:       "1",
			mockGetReturn: &models.BlogPost{ID: 1, Title: "Test", Description: "Desc", Body: "Body"},
			mockGetErr:    nil,
			mockDelErr:    errors.New("unable to delete post"),
			expectedCode:  http.StatusInternalServerError,
			mockGetCalled: true,
			mockDelCalled: true,
		},
	}

	// Iterate over test cases
	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			if test.mockGetCalled {
				// Mock GetByID if expected
				mockService.On("GetByID", mock.AnythingOfType("uint")).
					Return(test.mockGetReturn, test.mockGetErr).
					Once()
			}
			if test.mockDelCalled {
				// Mock Delete if expected
				mockService.On("Delete", mock.AnythingOfType("uint")).
					Return(test.mockDelErr).
					Once()
			}

			// Create test request
			req := httptest.NewRequest(http.MethodDelete, "/blog-post/"+test.paramID, nil)

			// Execute the request and capture the response
			resp, err := app.Test(req)
			require.NoError(t, err) // Ensure no Fiber errors

			// Assert response status
			assert.Equalf(t, test.expectedCode, resp.StatusCode, test.description)

			// Assert the mock was called only if expected
			if test.mockGetCalled {
				mockService.AssertCalled(t, "GetByID", mock.AnythingOfType("uint"))
			} else {
				mockService.AssertNotCalled(t, "GetByID")
			}

			if test.mockDelCalled {
				mockService.AssertCalled(t, "Delete", mock.AnythingOfType("uint"))
			} else {
				mockService.AssertNotCalled(t, "Delete")
			}
		})
	}
}

func TestCreatePost(t *testing.T) {
	// Create a Fiber app
	app := fiber.New()

	// Create a mock service
	mockService := new(mocks.BlogService)

	// Create a BlogController with the mock service
	bc := &BlogController{service: mockService}

	// Register the handler
	app.Post("/blog-post", bc.CreatePost)

	// Define test cases
	tests := []struct {
		description   string
		requestBody   interface{}
		mockReturnID  uint
		mockReturnErr error
		expectedCode  int
		mockCalled    bool // Whether Create() should be called
	}{
		{
			description: "success case - blog created",
			requestBody: models.CreateBlogRequest{
				Title:       "Test Blog",
				Description: "This is a test content",
				Body:        "body",
			},
			mockReturnID:  1,
			mockReturnErr: nil,
			expectedCode:  http.StatusCreated,
			mockCalled:    true,
		},
		{
			description: "failure case - missing title",
			requestBody: models.CreateBlogRequest{
				Title:       "",
				Description: "This is a test content",
				Body:        "body",
			},
			mockReturnID:  0,
			mockReturnErr: nil,
			expectedCode:  http.StatusBadRequest, // Expecting 400 due to validation
			mockCalled:    false,                 // Service should NOT be called
		},
		{
			description:   "failure case - invalid request body",
			requestBody:   "invalid-json", // Not a valid struct
			mockReturnID:  0,
			mockReturnErr: nil,
			expectedCode:  http.StatusBadRequest, // Expecting 400 due to bad JSON parsing
			mockCalled:    false,
		},
		{
			description: "failure case - service returns error",
			requestBody: models.CreateBlogRequest{
				Title:       "Test Blog",
				Description: "This is a test content",
				Body:        "body",
			},
			mockReturnID:  0,
			mockReturnErr: errors.New("unable to create blog"),
			expectedCode:  http.StatusInternalServerError, // Expecting 500
			mockCalled:    true,
		},
	}

	// Iterate over test cases
	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			var reqBytes []byte
			var err error

			// Convert request body to JSON if it’s a struct, else handle invalid case
			if body, ok := test.requestBody.(models.CreateBlogRequest); ok {
				reqBytes, err = json.Marshal(body)
				require.NoError(t, err)
			} else {
				reqBytes = []byte("invalid") // Simulating bad JSON
			}

			if test.mockCalled {
				// Mock only for cases where service should be called
				mockService.On("Create", mock.AnythingOfType("models.CreateBlogRequest")).
					Return(test.mockReturnID, test.mockReturnErr).
					Once()
			}

			// Create test request
			req := httptest.NewRequest(http.MethodPost, "/blog-post", bytes.NewBuffer(reqBytes))
			req.Header.Set("Content-Type", "application/json")

			// Execute the request and capture the response
			resp, err := app.Test(req)
			require.NoError(t, err) // Ensure no Fiber errors

			// Assert response status
			assert.Equalf(t, test.expectedCode, resp.StatusCode, test.description)

			// Assert the mock was called only if expected
			if test.mockCalled {
				mockService.AssertCalled(t, "Create", mock.AnythingOfType("models.CreateBlogRequest"))
			} else {
				mockService.AssertNotCalled(t, "Create")
			}
		})
	}
}
