package usecase

import (
	"encoding/json"
	"time"

	"github.com/notblessy/takeme-backend/model"
	"github.com/notblessy/takeme-backend/utils"
	"github.com/sirupsen/logrus"
)

type reactionUsecase struct {
	reactionRepo     model.ReactionRepository
	notificationRepo model.NotificationRepository
}

// NewReactionUsecase :nodoc:
func NewReactionUsecase(r model.ReactionRepository, n model.NotificationRepository) model.ReactionUsecase {
	return &reactionUsecase{
		reactionRepo:     r,
		notificationRepo: n,
	}
}

// Create :nodoc:
func (u *reactionUsecase) Create(req model.ReactionRequest) (model.Reaction, error) {
	logger := logrus.WithFields(logrus.Fields{
		"req": utils.Dump(req),
	})

	reaction := model.NewReaction(req)
	if reaction.ID == "" {
		reaction.ID = utils.GenerateID()
	}

	// create only if swiped left
	if reaction.Type == model.ReactionTypePass {
		err := u.reactionRepo.Create(reaction)
		if err != nil {
			logger.Error(err.Error())
			return model.Reaction{}, err
		}

		return reaction, nil
	}

	// if swiped right
	match, err := u.reactionRepo.FindMatch(reaction.UserTo, reaction.UserBy)
	if err != nil {
		logger.Error(err.Error())
		return model.Reaction{}, err
	}

	// if not matched
	if !match.IsMatch() {
		err := u.reactionRepo.Create(reaction)
		if err != nil {
			logger.Error(err.Error())
			return model.Reaction{}, err
		}

		return reaction, nil
	}

	now := time.Now()
	reaction.MatchedAt = &now
	match.MatchedAt = &now

	err = u.reactionRepo.CreateMatched(reaction, match)
	if err != nil {
		logger.Error(err.Error())
		return model.Reaction{}, err
	}

	go u.sendMatchNotification(match)

	return reaction, nil
}

func (u *reactionUsecase) sendMatchNotification(reaction model.Reaction) {
	content := model.MatchMessage{
		Type:    model.ReactionTypeLike,
		UserID:  reaction.UserTo,
		Message: "Congratulations! You matched",
	}

	b, err := json.Marshal(content)
	if err != nil {
		logrus.WithField("reaction", reaction).Error(err.Error())
	}

	notif := model.NewMatchNotification(reaction.UserBy, string(b))

	err = u.notificationRepo.Create(notif)
	if err != nil {
		logrus.WithField("reaction", reaction).Error(err.Error())
	}
}
