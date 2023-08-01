package repository

import (
	"fmt"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/golang/mock/gomock"
	"github.com/notblessy/takeme-backend/model"
	"github.com/notblessy/takeme-backend/utils"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type userRepo struct {
	db *gorm.DB
}

var userColumns = []string{"id", "name", "email", "password", "description", "gender", "preference", "age", "created_at", "updated_at", "deleted_at"}

func newMysqlMock() (db *gorm.DB, mock sqlmock.Sqlmock) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		panic(err)
	}

	dialector := mysql.New(mysql.Config{
		Conn:       mockDB,
		DriverName: "mysql",
	})

	columns := []string{"version"}
	mock.ExpectQuery("SELECT VERSION()").WithArgs().WillReturnRows(
		mock.NewRows(columns).FromCSVString("1"),
	)

	db, err = gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		logrus.Fatal(fmt.Sprintf("failed to connect: %s", err))
	}

	return
}

// TestUserRepo_Login :nodoc:
func TestUserRepo_Login(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	dbMock, sqlMock := newMysqlMock()

	ur := userRepo{
		db: dbMock,
	}

	userID := utils.GenerateID()
	now := time.Now()

	user := &model.User{
		ID:          userID,
		Name:        "Frederich Blessy",
		Email:       "frederichblessy@gmail.com",
		Password:    "asdzxc",
		Description: "Its me",
		Gender:      1,
		Preference:  2,
		Age:         26,
		CreatedAt:   now,
		UpdatedAt:   &now,
		DeletedAt:   nil,
	}

	t.Run("Success", func(t *testing.T) {
		sqlMock.ExpectBegin()

		sqlMock.ExpectExec(regexp.QuoteMeta("INSERT INTO `users` (`id`,`name`,`email`,`password`,`description`,`gender`,`preference`,`age`,`created_at`,`updated_at`,`deleted_at`) VALUES (?,?,?,?,?,?,?,?,?,?,?)")).
			WithArgs(
				user.ID,
				user.Name,
				user.Email,
				user.Password,
				user.Description,
				user.Gender,
				user.Preference,
				user.Age,
				user.CreatedAt,
				user.UpdatedAt,
				user.DeletedAt,
			).
			WillReturnResult(sqlmock.NewResult(1, 1))
		sqlMock.ExpectCommit()

		err := ur.db.Create(user).Error
		assert.NoError(t, err)
	})

	t.Run("Error", func(t *testing.T) {
		sqlMock.ExpectBegin()
		sqlMock.ExpectExec(regexp.QuoteMeta("INSERT INTO `users` (`id`,`name`,`email`,`password`,`description`,`gender`,`preference`,`age`,`created_at`,`updated_at`,`deleted_at`) VALUES (?,?,?,?,?,?,?,?,?,?,?)")).
			WithArgs(
				nil,
				user.Name,
				user.Email,
				user.Password,
				user.Description,
				user.Gender,
				user.Preference,
				user.Age,
				user.CreatedAt,
				user.UpdatedAt,
				user.DeletedAt,
			).
			WillReturnResult(sqlmock.NewResult(1, 1))
		sqlMock.ExpectCommit()
		sqlMock.ExpectCommit()

		err := ur.db.Create(user).Error
		assert.Error(t, err)
	})
}

// TestUserRepo_FindByEmail :nodoc:
func TestUserRepo_FindByEmail(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	dbMock, sqlMock := newMysqlMock()

	ur := userRepo{
		db: dbMock,
	}

	userID := utils.GenerateID()
	now := time.Now()

	user := &model.User{
		ID:          userID,
		Name:        "Frederich Blessy",
		Email:       "frederichblessy@gmail.com",
		Password:    "asdzxc",
		Description: "Its me",
		Gender:      1,
		Preference:  2,
		Age:         26,
		CreatedAt:   now,
		UpdatedAt:   &now,
		DeletedAt:   nil,
	}

	t.Run("Success", func(t *testing.T) {
		sqlMock.ExpectBegin()

		dbMock.Begin()
		rows := sqlmock.NewRows(userColumns).
			AddRow(user.ID, user.Name, user.Email, user.Password, user.Description, user.Gender, user.Preference, user.Age, user.CreatedAt, user.UpdatedAt, user.DeletedAt)

		sqlMock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users`")).
			WillReturnRows(rows)

		err := ur.db.Find(&[]model.User{}).Error
		assert.NoError(t, err)
	})

	t.Run("Error", func(t *testing.T) {
		sqlMock.ExpectBegin()

		dbMock.Begin()
		sqlMock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users`")).
			WillReturnError(gorm.ErrRecordNotFound)

		err := ur.db.Find(&[]model.User{}).Error
		assert.Error(t, err)
	})
}
