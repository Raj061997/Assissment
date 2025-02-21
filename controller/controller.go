package controller

import (
	"example/database"
	"example/models"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

// Create a blog post
// CreatePost creates a new blog post
// @Summary Create a new blog post
// @Description Create a new blog post with title, description, and body
// @Tags Blog
// @Accept  json
// @Produce  json
// @Param post body models.CreateBlogRequest true "Blog Post Data"
// @Success 201 {object} models.BlogPost
// @Failure 404 {object} models.ErrorResponse
// @Router /blog-post [post]
func CreatePost(c *fiber.Ctx) error {
	var req models.CreateBlogRequest

	// Parse request body
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(models.ErrorResponse{Error: "Invalid request body"})
	}

	// Validate the struct
	if err := models.Validate.Struct(req); err != nil {
		return c.Status(400).JSON(models.ErrorResponse{Error: "All fields are required"})
	}

	// Create a new BlogPost object
	post := models.BlogPost{
		Title:       req.Title,
		Description: req.Description,
		Body:        req.Body,
	}

	database.DB.Create(&post)
	return c.Status(201).JSON(post)
}

// Get all blog posts
// GetPosts retrieves all blog posts
// @Summary Get all blog posts
// @Description Retrieve a list of all blog posts
// @Tags Blog
// @Produce json
// @Success 200 {array} models.BlogPost
// @Router /blog-post [get]
func GetPosts(c *fiber.Ctx) error {
	var posts []models.BlogPost
	database.DB.Find(&posts)
	return c.JSON(posts)
}

// Get a single blog post
// GetPost retrieves a single blog post by ID
// @Summary Get a single blog post
// @Description Get details of a blog post by ID
// @Tags Blog
// @Produce json
// @Param id path int true "Blog Post ID"
// @Success 200 {object} models.BlogPost
// @Failure 404 {object} models.ErrorResponse
// @Router /blog-post/{id} [get]
func GetPost(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil || id <= 0 {
		return c.Status(400).JSON(models.ErrorResponse{Error: "Invalid ID parameter"})
	}
	var post models.BlogPost
	if err := database.DB.First(&post, id).Error; err != nil {
		return c.Status(404).JSON(models.ErrorResponse{Error: "Post not found"})
	}
	return c.JSON(post)
}

// Update a blog post
// UpdatePost updates a blog post by ID
// @Summary Update a blog post
// @Description Update a blog post's title, description, or body by ID
// @Tags Blog
// @Accept json
// @Produce json
// @Param id path int true "Blog Post ID"
// @Param post body models.UpdateBlogRequest  true "Updated Blog Post Data"
// @Success 200 {object} models.BlogPost
// @Failure 400 {object} models.ErrorResponse
// @Failure 404 {object} models.ErrorResponse
// @Router /blog-post/{id} [patch]
func UpdatePost(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil || id <= 0 {
		return c.Status(400).JSON(models.ErrorResponse{Error: "Invalid ID parameter"})
	}
	var post models.BlogPost
	if err := database.DB.First(&post, id).Error; err != nil {
		return c.Status(404).JSON(models.ErrorResponse{Error: err.Error()})
	}
	var req models.UpdateBlogRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(models.ErrorResponse{Error: "invalid request body"})
	}
	// Update only provided fields
	if req.Title != nil {
		post.Title = *req.Title
	}
	if req.Description != nil {
		post.Description = *req.Description
	}
	if req.Body != nil {
		post.Body = *req.Body
	}

	database.DB.Save(&post)
	return c.JSON(post)
}

// Delete a blog post
// DeletePost deletes a blog post by ID
// @Summary Delete a blog post
// @Description Delete a blog post by ID
// @Tags Blog
// @Param id path int true "Blog Post ID"
// @Success 204 "No Content"
// @Failure 404 {object} models.ErrorResponse
// @Router /blog-post/{id} [delete]
func DeletePost(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil || id <= 0 {
		return c.Status(400).JSON(models.ErrorResponse{Error: "Invalid ID parameter"})
	}
	var post models.BlogPost
	if err := database.DB.First(&post, id).Error; err != nil {
		return c.Status(404).JSON(models.ErrorResponse{Error: "Post not found"})
	}
	database.DB.Delete(&post)
	return c.Status(204).Send(nil)
}
