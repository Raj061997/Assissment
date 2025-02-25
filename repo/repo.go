package repo

import (
	"example/models"
	"gorm.io/gorm"
)

// BlogService defines methods for blog operations.
type Repository interface {
	Create(post *models.BlogPost) error
	GetAll() ([]models.BlogPost, error)
	GetByID(id uint) (*models.BlogPost, error)
	Update(id uint, post *models.BlogPost) error
	Delete(id uint) error
}

// BlogServiceImpl implements BlogService
type repo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) *repo {
	return &repo{db: db}
}

// Create a new blog post
func (r *repo) Create(post *models.BlogPost) error {
	return r.db.Create(post).Error
}

// Get all blog posts
func (r *repo) GetAll() ([]models.BlogPost, error) {
	var posts []models.BlogPost
	err := r.db.Find(&posts).Error
	return posts, err
}

// Get a single blog post by ID
func (r *repo) GetByID(id uint) (*models.BlogPost, error) {
	var post models.BlogPost
	err := r.db.First(&post, id).Error
	return &post, err
}

// Update a blog post
func (r *repo) Update(id uint, post *models.BlogPost) error {
	return r.db.Model(&models.BlogPost{}).Where("id = ?", id).Updates(post).Error
}

// Delete a blog post
func (r *repo) Delete(id uint) error {
	return r.db.Delete(&models.BlogPost{}, id).Error
}
