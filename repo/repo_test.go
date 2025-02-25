package repo_test

func TestRepository_Add(t *testing.T) {
	t.Parallel()
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		ctx       context.Context
		log       log.Logger
		employees []model.Employee
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
				ctx: context.Background(),
				log: func() log.Logger {
					logger := new(logMock.Logger)
					logger.On("Errorf", mock.Anything, mock.Anything).Return(nil)
					return logger
				}(),
				employees: []model.Employee{{FirstName: "test"}},
			},
			wantErr: false,
		},
		{
			name: "negative case",
			fields: fields{
				db: func() *gorm.DB {
					db, dbmock := dbMock.NewGormMock(t)
					dbmock.ExpectBegin()
					dbmock.ExpectExec(regexp.QuoteMeta("")).
						WillReturnError(errRepo)
					dbmock.ExpectCommit()
					return db
				}(),
			},
			args: args{
				ctx: context.Background(),
				log: func() log.Logger {
					logger := new(logMock.Logger)
					logger.On("Errorf", mock.Anything, mock.Anything).Return(nil)
					return logger
				}(),
				employees: []model.Employee{{FirstName: "test"}},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			r := repo.NewRepository(tt.fields.db)
			if err := r.Add(tt.args.ctx, tt.args.log, tt.args.employees); (err != nil) != tt.wantErr {
				t.Errorf("Add() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
