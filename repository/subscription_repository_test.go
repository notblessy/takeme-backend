package repository

import (
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/golang/mock/gomock"
	"github.com/notblessy/takeme-backend/model"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

type subscriptionRepo struct {
	db *gorm.DB
}

var subscriptionColumns = []string{"id", "user_id", "subscription_plan_id", "is_active", "created_at", "updated_at", "deleted_at"}

// TestSubscriptionRepo_Create :nodoc:
func TestSubscriptionRepo_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	dbMock, sqlMock := newMysqlMock()

	sr := subscriptionRepo{
		db: dbMock,
	}

	now := time.Now()

	subs := &model.Subscription{
		ID:                 1,
		UserID:             "42431469",
		SubscriptionPlanID: 2,
		IsActive:           true,
		CreatedAt:          now,
		UpdatedAt:          &now,
		DeletedAt:          nil,
	}

	t.Run("Success", func(t *testing.T) {
		sqlMock.ExpectBegin()

		sqlMock.ExpectExec(regexp.QuoteMeta("INSERT INTO `subscriptions` (`user_id`,`subscription_plan_id`,`is_active`,`created_at`,`updated_at`,`deleted_at`,`id`) VALUES (?,?,?,?,?,?,?)")).
			WithArgs(
				subs.UserID,
				subs.SubscriptionPlanID,
				subs.IsActive,
				subs.CreatedAt,
				subs.UpdatedAt,
				subs.DeletedAt,
				subs.ID,
			).
			WillReturnResult(sqlmock.NewResult(1, 1))
		sqlMock.ExpectCommit()

		err := sr.db.Create(subs).Error
		assert.NoError(t, err)
	})

	t.Run("Error", func(t *testing.T) {
		sqlMock.ExpectBegin()
		sqlMock.ExpectExec(regexp.QuoteMeta("INSERT INTO `subscriptions` (`user_id`,`subscription_plan_id`,`is_active`,`created_at`,`updated_at`,`deleted_at`,`id`) VALUES (?,?,?,?,?,?,?)")).
			WithArgs(
				subs.UserID,
				subs.SubscriptionPlanID,
				subs.IsActive,
				subs.CreatedAt,
				subs.UpdatedAt,
				subs.DeletedAt,
				nil,
			).
			WillReturnResult(sqlmock.NewResult(1, 1))
		sqlMock.ExpectCommit()
		sqlMock.ExpectCommit()

		err := sr.db.Create(subs).Error
		assert.Error(t, err)
	})
}

// TestSubscriptionRepo_FindByID :nodoc:
func TestSubscriptionRepo_FindByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	dbMock, sqlMock := newMysqlMock()

	sr := subscriptionRepo{
		db: dbMock,
	}

	now := time.Now()

	subs := &model.Subscription{
		ID:                 1,
		UserID:             "42431469",
		SubscriptionPlanID: 2,
		IsActive:           true,
		CreatedAt:          now,
		UpdatedAt:          &now,
		DeletedAt:          nil,
	}

	t.Run("Success", func(t *testing.T) {
		sqlMock.ExpectBegin()

		dbMock.Begin()
		rows := sqlmock.NewRows(subscriptionColumns).
			AddRow(subs.ID, subs.UserID, subs.SubscriptionPlanID, subs.IsActive, subs.CreatedAt, subs.UpdatedAt, subs.DeletedAt)

		sqlMock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `subscriptions`")).
			WillReturnRows(rows)

		err := sr.db.Find(&[]model.Subscription{}).Error
		assert.NoError(t, err)
	})

	t.Run("Error", func(t *testing.T) {
		sqlMock.ExpectBegin()

		dbMock.Begin()
		sqlMock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `subscriptions`")).
			WillReturnError(gorm.ErrRecordNotFound)

		err := sr.db.Find(&[]model.Subscription{}).Error
		assert.Error(t, err)
	})
}
