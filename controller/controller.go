package controller

import (
	"example/models"
	"example/service"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type BlogController struct {
	service service.Service
}

func NewController(service service.Service) BlogController {
	return BlogController{
		service: service,
	}
}

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
func (bc *BlogController) CreatePost(c *fiber.Ctx) error {
	var req models.CreateBlogRequest

	// Parse request body
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(models.ErrorResponse{Error: "Invalid request body"})
	}

	// Validate the struct
	if err := models.Validate.Struct(req); err != nil {
		return c.Status(400).JSON(models.ErrorResponse{Error: "All fields are required"})
	}
	id, err := bc.service.Create(req)
	if err != nil {
		return c.Status(500).JSON(models.ErrorResponse{Error: "unable to create blog"})

	}
	return c.Status(201).JSON(models.BlogPost{ID: id})
}

// Get all blog posts
// GetPosts retrieves all blog posts
// @Summary Get all blog posts
// @Description Retrieve a list of all blog posts
// @Tags Blog
// @Produce json
// @Success 200 {array} models.BlogPost
// @Router /blog-post [get]
func (bc *BlogController) GetPosts(c *fiber.Ctx) error {

	posts, err := bc.service.GetAll()

	if err != nil {
		return c.Status(404).JSON(models.ErrorResponse{Error: "unable to find blog"})
	}
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
func (bc *BlogController) GetPost(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil || id <= 0 {
		return c.Status(400).JSON(models.ErrorResponse{Error: "Invalid ID parameter"})
	}

	post, err := bc.service.GetByID(uint(id))
	if err != nil {
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
func (bc *BlogController) UpdatePost(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil || id <= 0 {
		return c.Status(400).JSON(models.ErrorResponse{Error: "Invalid ID parameter"})
	}
	var req models.UpdateBlogRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(models.ErrorResponse{Error: "invalid request body"})
	}

	post, err := bc.service.Update(uint(id), &req)
	if err != nil {
		return c.Status(500).JSON(models.ErrorResponse{Error: "unabel to update post"})
	}
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
func (bc *BlogController) DeletePost(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil || id <= 0 {
		return c.Status(400).JSON(models.ErrorResponse{Error: "Invalid ID parameter"})
	}

	_, err = bc.service.GetByID(uint(id))
	if err != nil {
		return c.Status(404).JSON(models.ErrorResponse{Error: "Post not found"})
	}
	if err := bc.service.Delete(uint(id)); err != nil {
		return c.Status(500).JSON(models.ErrorResponse{Error: "unable to delete error"})
	}
	return c.Status(204).Send(nil)

}
