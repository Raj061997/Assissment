package repo_test

import (
	"database/sql/driver"
	dbMock "example/database/mocks"
	"example/models"
	"example/repo"
	"reflect"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/gorm"
)

func Test_repo_Create(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	ti := time.Now()
	type args struct {
		post *models.BlogPost
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Negative",
			fields: fields{
				db: func() *gorm.DB {
					// Mock the GORM DB with expected transaction and query
					db, dbmock := dbMock.NewGormMock(t)

					// Expecting transaction begin
					dbmock.ExpectBegin()

					// Expecting the exact SQL query for insert (with placeholders)
					dbmock.ExpectExec(regexp.QuoteMeta("INSERT INTO \"blog_posts\" (\"title\",\"description\",\"body\",\"created_at\",\"updated_at\") VALUES ($1,$2,$3,$4,$5) RETURNING \"id\"")).
						// Here we use mock.AnythingOfType to match the time fields as dynamic time values

						WillReturnResult(driver.ResultNoRows) // Return no rows on successful insert

					// Expecting transaction commit
					dbmock.ExpectCommit()

					return db
				}(),
			},
			args: args{
				post: &models.BlogPost{Title: "title", Description: "description", Body: "body", CreatedAt: ti, UpdatedAt: ti},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := repo.NewRepo(tt.fields.db)
			if _, err := r.Create(tt.args.post); (err != nil) != tt.wantErr {
				t.Errorf("repo.Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_repo_GetAll(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	tests := []struct {
		name    string
		fields  fields
		want    []models.BlogPost
		wantErr bool
	}{
		{
			name: "positive",
			fields: fields{
				db: func() *gorm.DB {
					db, dbmock := dbMock.NewGormMock(t)
					selectRows := sqlmock.NewRows([]string{"id"}).AddRow("12345").AddRow("123456")
					dbmock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "blog_posts"`)).
						WillReturnRows(selectRows)
					return db
				}(),
			},
			want: []models.BlogPost{
				{ID: 12345},
				{ID: 123456},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := repo.NewRepo(tt.fields.db)
			got, err := r.GetAll()
			if (err != nil) != tt.wantErr {
				t.Errorf("repo.GetAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("repo.GetAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_repo_GetByID(t *testing.T) {
	type fields struct {
		db *gorm.DB
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
			name: "positive",
			args: args{id: 1},
			fields: fields{
				db: func() *gorm.DB {
					db, dbmock := dbMock.NewGormMock(t)
					selectRows := sqlmock.NewRows([]string{"id"}).AddRow("12345").AddRow("123456")
					dbmock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "blog_posts" WHERE "blog_posts"."id" = $1 ORDER BY "blog_posts"."id" LIMIT $2`)).
						WillReturnRows(selectRows)
					return db
				}(),
			},
			want: &models.BlogPost{ID: 12345},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := repo.NewRepo(tt.fields.db)
			got, err := r.GetByID(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("repo.GetByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("repo.GetByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_repo_Update(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		id   uint
		post *models.BlogPost
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "positive case",
			fields: fields{
				db: func() *gorm.DB {
					db, dbmock := dbMock.NewGormMock(t)
					dbmock.ExpectBegin()
					dbmock.ExpectExec(regexp.QuoteMeta("")).
						WillReturnResult(sqlmock.NewResult(1234, 1))
					dbmock.ExpectCommit()
					return db
				}(),
			},
			args: args{

				id:   1,
				post: &models.BlogPost{ID: 1},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := repo.NewRepo(tt.fields.db)
			if err := r.Update(tt.args.id, tt.args.post); (err != nil) != tt.wantErr {
				t.Errorf("repo.Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_repo_Delete(t *testing.T) {
	type fields struct {
		db *gorm.DB
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
				db: func() *gorm.DB {
					db, dbmock := dbMock.NewGormMock(t)
					dbmock.ExpectBegin()
					dbmock.ExpectExec(regexp.QuoteMeta("")).
						WillReturnResult(sqlmock.NewResult(1234, 1))
					dbmock.ExpectCommit()
					return db
				}(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := repo.NewRepo(tt.fields.db)

			if err := r.Delete(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("repo.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
