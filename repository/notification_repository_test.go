package repository

import (
	"encoding/json"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/golang/mock/gomock"
	"github.com/notblessy/takeme-backend/model"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

type notificationRepo struct {
	db *gorm.DB
}

// TestNotificationRepo_Create :nodoc:
func TestNotificationRepo_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	dbMock, sqlMock := newMysqlMock()

	sr := notificationRepo{
		db: dbMock,
	}

	now := time.Now()

	msg := model.MatchMessage{
		Type:    "LIKE",
		UserID:  "4321",
		Message: "You have match",
	}

	byteMsg, _ := json.Marshal(msg)

	notif := &model.Notification{
		ID:        "1234",
		UserID:    "42431469",
		Content:   string(byteMsg),
		IsRead:    false,
		CreatedAt: now,
		UpdatedAt: &now,
		DeletedAt: nil,
	}

	t.Run("Success", func(t *testing.T) {
		sqlMock.ExpectBegin()

		sqlMock.ExpectExec(regexp.QuoteMeta("INSERT INTO `notifications` (`id`,`user_id`,`content`,`is_read`,`created_at`,`updated_at`,`deleted_at`) VALUES (?,?,?,?,?,?,?)")).
			WithArgs(
				notif.ID,
				notif.UserID,
				notif.Content,
				notif.IsRead,
				notif.CreatedAt,
				notif.UpdatedAt,
				notif.DeletedAt,
			).
			WillReturnResult(sqlmock.NewResult(1, 1))
		sqlMock.ExpectCommit()

		err := sr.db.Create(notif).Error
		assert.NoError(t, err)
	})

	t.Run("Error", func(t *testing.T) {
		sqlMock.ExpectBegin()
		sqlMock.ExpectExec(regexp.QuoteMeta("INSERT INTO `notifications` (`id`,`user_id`,`content`,`is_read`,`created_at`,`updated_at`,`deleted_at`) VALUES (?,?,?,?,?,?,?)")).
			WithArgs(
				nil,
				notif.UserID,
				notif.Content,
				notif.IsRead,
				notif.CreatedAt,
				notif.UpdatedAt,
				notif.DeletedAt,
			).
			WillReturnResult(sqlmock.NewResult(1, 1))
		sqlMock.ExpectCommit()
		sqlMock.ExpectCommit()

		err := sr.db.Create(notif).Error
		assert.Error(t, err)
	})
}
