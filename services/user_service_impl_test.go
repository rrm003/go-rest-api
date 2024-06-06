package services

import (
	"fmt"
	"go-rest-api/database"
	mockDB "go-rest-api/database/mocks"
	"go-rest-api/models"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"gorm.io/gorm"
)

func Test_userService_SignUp(t *testing.T) {
	ctrl := gomock.NewController(t)
	mkdb := mockDB.NewMockDatabase(ctrl)

	type fields struct {
		db database.Database
	}
	type args struct {
		user models.User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		setup   func(*mockDB.MockDatabase)
		wantErr bool
	}{
		{
			name: "Successulf signup",
			fields: fields{
				db: mkdb,
			},
			args: args{
				user: models.User{
					Username: "rrm",
					Password: "roeeo",
					Country:  "india",
				},
			},
			setup: func(mockDB *mockDB.MockDatabase) {
				mockDB.EXPECT().Create(gomock.Any()).Return(&gorm.DB{}).Times(1)
			},
			wantErr: false,
		},
		{
			name: "Fail signup country empty",
			fields: fields{
				db: mkdb,
			},
			args: args{
				user: models.User{
					Username: "rrm",
					Password: "roeeo",
					Country:  "",
				},
			},
			setup: func(mockDB *mockDB.MockDatabase) {
				mockDB.EXPECT().Create(gomock.Any()).Return(&gorm.DB{Error: fmt.Errorf("country cannot be empty")}).Times(1)
			},
			wantErr: true,
		},
		{
			name: "Fail password too long",
			fields: fields{
				db: mkdb,
			},
			args: args{
				user: models.User{
					Username: "rrm",
					Password: "123456789123456789123456789123456789123456789123456789123456789123456789123456789",
					Country:  "",
				},
			},
			setup: func(mockDB *mockDB.MockDatabase) {
				mockDB.EXPECT().Create(gomock.Any()).Return(&gorm.DB{Error: fmt.Errorf("country cannot be empty")}).Times(0)
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &userService{
				db: tt.fields.db,
			}

			if tt.setup != nil {
				tt.setup(mkdb)
			}

			if err := s.SignUp(tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("userService.SignUp() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_userService_GetUsers(t *testing.T) {
	ctrl := gomock.NewController(t)
	mkdb := mockDB.NewMockDatabase(ctrl)

	type fields struct {
		db database.Database
	}

	tests := []struct {
		name    string
		fields  fields
		setup   func(*mockDB.MockDatabase)
		want    []models.User
		wantErr bool
	}{
		{
			name: "success",
			fields: fields{
				db: mkdb,
			},
			setup: func(md *mockDB.MockDatabase) {
				mkdb.EXPECT().Find(gomock.Any()).SetArg(0, []models.User{{Username: "user1", Password: "123", Country: "usa"}, {Username: "user2", Password: "12321", Country: "uae"}}).Return(&gorm.DB{}).Times(1)
			},
			want:    []models.User{{Username: "user1", Password: "123", Country: "usa"}, {Username: "user2", Password: "12321", Country: "uae"}},
			wantErr: false,
		},
		{
			name: "failure: invalid query",
			fields: fields{
				db: mkdb,
			},
			setup: func(md *mockDB.MockDatabase) {
				mkdb.EXPECT().Find(gomock.Any()).SetArg(0, []models.User{{Username: "user1", Password: "123", Country: "usa"}, {Username: "user2", Password: "12321", Country: "uae"}}).Return(&gorm.DB{Error: fmt.Errorf("invalid query syntax")}).Times(1)
			},

			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &userService{
				db: tt.fields.db,
			}

			tt.setup(mkdb)

			got, err := s.GetUsers()
			if (err != nil) != tt.wantErr {
				t.Errorf("userService.GetUsers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userService.GetUsers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userService_GetUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	mkdb := mockDB.NewMockDatabase(ctrl)

	type fields struct {
		db database.Database
	}
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		setup   func(*mockDB.MockDatabase)
		want    models.User
		wantErr bool
	}{
		{
			name: "success",
			fields: fields{
				db: mkdb,
			},
			args: args{
				id: "1",
			},
			setup: func(md *mockDB.MockDatabase) {
				mkdb.EXPECT().First(gomock.Any(), "1").SetArg(0, models.User{Username: "rrm", Password: "encrypted", Country: "usa"}).Return(&gorm.DB{}).Times(1)
			},
			want:    models.User{Username: "rrm", Country: "usa", Password: "encrypted"},
			wantErr: false,
		},
		{
			name: "failure",
			fields: fields{
				db: mkdb,
			},
			args: args{
				id: "1",
			},
			setup: func(md *mockDB.MockDatabase) {
				mkdb.EXPECT().First(gomock.Any(), "1").SetArg(0, models.User{}).Return(&gorm.DB{Error: gorm.ErrRecordNotFound}).Times(1)
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &userService{
				db: tt.fields.db,
			}

			tt.setup(mkdb)

			got, err := s.GetUser(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("userService.GetUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userService.GetUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userService_UpdateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	mkdb := mockDB.NewMockDatabase(ctrl)

	type fields struct {
		db database.Database
	}
	type args struct {
		id   string
		user models.User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		setup   func(*mockDB.MockDatabase)
		wantErr bool
	}{

		{
			name: "User not found",
			fields: fields{
				db: mkdb,
			},
			args: args{
				id: "1",
				user: models.User{
					Username: "rrm",
					Country:  "usa",
					Password: "encrypted",
				},
			},
			setup: func(md *mockDB.MockDatabase) {
				md.EXPECT().First(gomock.Any(), "1").Return(&gorm.DB{Error: gorm.ErrRecordNotFound}).Times(1)
			},
			wantErr: true,
		},
		{
			name: "Database error when fetching user",
			fields: fields{
				db: mkdb,
			},
			args: args{
				id: "1",
				user: models.User{
					Username: "rrm",
					Country:  "usa",
					Password: "encrypted",
				},
			},
			setup: func(md *mockDB.MockDatabase) {
				md.EXPECT().First(gomock.Any(), "1").Return(&gorm.DB{Error: fmt.Errorf("database error")}).Times(1)
			},
			wantErr: true,
		},
		{
			name: "Successful update",
			fields: fields{
				db: mkdb,
			},
			args: args{
				id: "1",
				user: models.User{
					Username: "rrm",
					Country:  "usa",
					Password: "encrypted",
				},
			},
			setup: func(md *mockDB.MockDatabase) {
				existingUser := models.User{Username: "rrm", Country: "india", Password: "oldpassword"}
				md.EXPECT().First(gomock.Any(), "1").SetArg(0, existingUser).Return(&gorm.DB{Error: nil}).Times(1)
				md.EXPECT().Save(gomock.Any()).Return(&gorm.DB{Error: nil}).Times(1)
			},
			wantErr: false,
		},
		{
			name: "Database error when saving user",
			fields: fields{
				db: mkdb,
			},
			args: args{
				id: "1",
				user: models.User{
					Username: "rrm",
					Country:  "usa",
					Password: "encrypted",
				},
			},
			setup: func(md *mockDB.MockDatabase) {
				existingUser := models.User{Username: "rrm", Country: "india", Password: "oldpassword"}
				md.EXPECT().First(gomock.Any(), "1").SetArg(0, existingUser).Return(&gorm.DB{Error: nil}).Times(1)
				md.EXPECT().Save(gomock.Any()).Return(&gorm.DB{Error: fmt.Errorf("save error")}).Times(1)
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &userService{
				db: tt.fields.db,
			}

			tt.setup(mkdb)

			if err := s.UpdateUser(tt.args.id, tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("userService.UpdateUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_userService_DeleteUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	mkdb := mockDB.NewMockDatabase(ctrl)

	type fields struct {
		db database.Database
	}
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		setup   func(*mockDB.MockDatabase)
		wantErr bool
	}{
		{
			name: "success in deletion",
			fields: fields{
				db: mkdb,
			},
			args: args{
				id: "1",
			},
			setup: func(md *mockDB.MockDatabase) {
				md.EXPECT().First(gomock.Any(), "1").Return(&gorm.DB{})
				md.EXPECT().Delete(gomock.Any()).Return(&gorm.DB{})
			},
			wantErr: false,
		},
		{
			name: "fail to find",
			fields: fields{
				db: mkdb,
			},
			args: args{
				id: "1",
			},
			setup: func(md *mockDB.MockDatabase) {
				md.EXPECT().First(gomock.Any(), "1").Return(&gorm.DB{Error: gorm.ErrRecordNotFound})
			},
			wantErr: true,
		},
		{
			name: "failure in deletion",
			fields: fields{
				db: mkdb,
			},
			args: args{
				id: "1",
			},
			setup: func(md *mockDB.MockDatabase) {
				md.EXPECT().First(gomock.Any(), "1").Return(&gorm.DB{})
				md.EXPECT().Delete(gomock.Any()).Return(&gorm.DB{Error: gorm.ErrCheckConstraintViolated})
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &userService{
				db: tt.fields.db,
			}

			tt.setup(mkdb)

			if err := s.DeleteUser(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("userService.DeleteUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
