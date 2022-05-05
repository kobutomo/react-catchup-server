package usecases

import (
	"context"

	"github.com/kobutomo/react-catchup-server/server/domain/repositories"
)

type SampleUsecase interface {
	Execute(ctx context.Context, xrid string, in string) (string, error)
}

type SampleInteractor struct {
	userRepository repositories.UserRepository
}

func NewSampleInteractor(userRepository repositories.UserRepository) SampleUsecase {
	return SampleInteractor{
		userRepository: userRepository,
	}
}

func (s SampleInteractor) Execute(ctx context.Context, xrid string, in string) (string, error) {
	user, err := s.userRepository.GetByEmail(ctx, xrid, in)
	if err != nil {
		return "", err
	}
	return user.Username, nil
}
