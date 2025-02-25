package service

import (
	"example/models"
	"example/repo"
	"fmt"
)

// BlogService defines methods for blog operations.
//
//go:generate mockery --name=Service --outpkg mocks
type Service interface {
	Create(req models.CreateBlogRequest) (uint, error)
	GetAll() ([]models.BlogPost, error)
	GetByID(id uint) (*models.BlogPost, error)
	Update(id uint, post *models.UpdateBlogRequest) (*models.BlogPost, error)
	Delete(id uint) error
}

// BlogServiceImpl implements BlogService
type service struct {
	repo repo.Repository
}

func NewService(repo repo.Repository) *service {
	return &service{repo: repo}
}

// Create a new blog post
func (s *service) Create(req models.CreateBlogRequest) (uint, error) {
	id, err := s.repo.Create(&models.BlogPost{
		Title:       req.Title,
		Description: req.Description,
		Body:        req.Body,
	})
	return id, err
}

// Get all blog posts
func (s *service) GetAll() ([]models.BlogPost, error) {
	posts, err := s.repo.GetAll()
	if err != nil {
		return []models.BlogPost{}, fmt.Errorf("unable to fetch posts: %w", err)
	}
	return posts, err
}

// Get a single blog post by ID
func (s *service) GetByID(id uint) (*models.BlogPost, error) {
	post, err := s.repo.GetByID(id)
	return post, err
}

// Update a blog post
func (s *service) Update(id uint, req *models.UpdateBlogRequest) (*models.BlogPost, error) {

	post, err := s.GetByID(id)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch post : %w", err)
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
	err = s.repo.Update(id, post)
	return post, err

}

// Delete a blog post
func (s *service) Delete(id uint) error {
	return s.repo.Delete(id)

}
