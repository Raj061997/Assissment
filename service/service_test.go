package service

import (
	"example/mocks"
	"example/models"
	"example/repo"
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/c2fo/testify/mock"
)

func Test_service_Create(t *testing.T) {
	type fields struct {
		repo repo.Repository
	}
	type args struct {
		req models.CreateBlogRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "positive",
			fields: fields{
				repo: func() repo.Repository {
					repo := new(mocks.Repository)
					repo.On("Create", mock.Anything).Return(nil)
					return repo
				}(),
			},
			args: args{
				req: models.CreateBlogRequest{Title: "title", Description: "Create", Body: "body"},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				repo: tt.fields.repo,
			}
			if err := s.Create(tt.args.req); (err != nil) != tt.wantErr {
				t.Errorf("service.Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_service_GetAll(t *testing.T) {
	type fields struct {
		repo repo.Repository
	}
	ti := time.Now()
	tests := []struct {
		name    string
		fields  fields
		want    []models.BlogPost
		wantErr bool
	}{
		{
			name: "positive",
			fields: fields{
				repo: func() repo.Repository {
					repo := new(mocks.Repository)
					repo.On("GetAll").Return([]models.BlogPost{{ID: 1, Title: "title", Description: "description", Body: "body", CreatedAt: ti, UpdatedAt: ti}}, nil)
					return repo
				}(),
			},
			want:    []models.BlogPost{{ID: 1, Title: "title", Description: "description", Body: "body", CreatedAt: ti, UpdatedAt: ti}},
			wantErr: false,
		},
		{
			name: "negative",
			fields: fields{
				repo: func() repo.Repository {
					repo := new(mocks.Repository)
					repo.On("GetAll").Return([]models.BlogPost{}, fmt.Errorf("unable to fetch post"))
					return repo
				}(),
			},
			want:    []models.BlogPost{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				repo: tt.fields.repo,
			}
			got, err := s.GetAll()
			if (err != nil) != tt.wantErr {
				t.Errorf("service.GetAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("service.GetAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_service_GetByID(t *testing.T) {
	type fields struct {
		repo repo.Repository
	}
	type args struct {
		id uint
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *models.BlogPost
		wantErr bool
	}{
		{
			name: "postive",
			fields: fields{
				repo: func() repo.Repository {
					repo := new(mocks.Repository)
					repo.On("GetByID", mock.Anything).Return(&models.BlogPost{ID: 1, Title: "title", Description: "description", Body: "body"}, nil)
					return repo
				}(),
			},
			want:    &models.BlogPost{ID: 1, Title: "title", Description: "description", Body: "body"},
			wantErr: false,
			args:    args{id: 1},
		},
		{
			name: "negative",
			fields: fields{
				repo: func() repo.Repository {
					repo := new(mocks.Repository)
					repo.On("GetByID", mock.Anything).Return(nil, fmt.Errorf("unable to fetch post"))
					return repo
				}(),
			},

			want:    nil,
			wantErr: true,
			args:    args{id: 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				repo: tt.fields.repo,
			}
			got, err := s.GetByID(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.GetByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("service.GetByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_service_Update(t *testing.T) {
	type fields struct {
		repo repo.Repository
	}
	ss := "mockString"
	type args struct {
		id  uint
		req *models.UpdateBlogRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *models.BlogPost
		wantErr bool
	}{
		{
			name: "name",
			fields: fields{
				repo: func() repo.Repository {
					repo := new(mocks.Repository)
					repo.On("GetByID", mock.Anything).Return(nil, fmt.Errorf("unable to fetch post"))
					return repo
				}(),
			},
			wantErr: true,
		},

		{
			name: "name",
			fields: fields{
				repo: func() repo.Repository {
					repo := new(mocks.Repository)
					repo.On("GetByID", mock.Anything).Return(&models.BlogPost{ID: 1, Title: "title", Description: "description", Body: "body"}, nil)
					repo.On("Update", mock.Anything, mock.Anything).Return(nil)

					return repo
				}(),
			},
			want:    &models.BlogPost{ID: 1, Title: ss, Description: ss, Body: ss},
			args:    args{id: 1, req: &models.UpdateBlogRequest{Title: &ss, Description: &ss, Body: &ss}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				repo: tt.fields.repo,
			}
			got, err := s.Update(tt.args.id, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("service.Update() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_service_Delete(t *testing.T) {
	type fields struct {
		repo repo.Repository
	}
	type args struct {
		id uint
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "positive",
			fields: fields{
				repo: func() repo.Repository {
					repo := new(mocks.Repository)
					repo.On("Delete", mock.Anything).Return(nil)

					return repo
				}(),
			},
			args: args{
				id: 1,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				repo: tt.fields.repo,
			}
			if err := s.Delete(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("service.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
