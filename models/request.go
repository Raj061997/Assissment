package models

import "github.com/go-playground/validator/v10"

type CreateBlogRequest struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	Body        string `json:"body" validate:"required"`
}

type UpdateBlogRequest struct {
	Title       *string `json:"title" validate:"omitempty"`       // Optional
	Description *string `json:"description" validate:"omitempty"` // Optional
	Body        *string `json:"body" validate:"omitempty"`        // Optional
}

var Validate = validator.New()
