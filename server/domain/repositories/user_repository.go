package repositories

import (
	"github.com/kobutomo/react-catchup-server/server/domain/models"
	"golang.org/x/net/context"
)

type UserRepository interface {
	GetByEmail(ctx context.Context, xrid string, email string) (*models.User, error)
}
