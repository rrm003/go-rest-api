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
			name: "Successulf signup",
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
