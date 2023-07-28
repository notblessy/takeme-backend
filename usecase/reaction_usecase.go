package usecase

import (
	"github.com/notblessy/takeme-backend/model"
	"github.com/notblessy/takeme-backend/utils"
	"github.com/sirupsen/logrus"
)

type reactionUsecase struct {
	reactionRepo model.ReactionRepository
}

// NewReactionUsecase :nodoc:
func NewReactionUsecase(u model.ReactionRepository) model.ReactionUsecase {
	return &reactionUsecase{
		reactionRepo: u,
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

	err := u.reactionRepo.Create(&reaction)
	if err != nil {
		logger.Error(err.Error())
		return model.Reaction{}, err
	}

	return reaction, nil
}
