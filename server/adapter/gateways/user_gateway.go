package gateways

import (
	"context"
	"database/sql"
	"log"

	"github.com/kobutomo/react-catchup-server/server/domain/models"
	"github.com/kobutomo/react-catchup-server/server/domain/repositories"
	"github.com/kobutomo/react-catchup-server/server/infrastructure/dbmodels"
	"github.com/kobutomo/react-catchup-server/server/infrastructure/mysql"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type UserGateway struct {
	db *sql.DB
}

func NewUserGateway(m mysql.IMySQL) repositories.UserRepository {
	return UserGateway{
		db: m.GetConn(),
	}
}

func (u UserGateway) GetByEmail(ctx context.Context, xrid string, email string) (*models.User, error) {
	dm, err := dbmodels.Users(qm.Where("email = ?", email)).One(ctx, u.db)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return DBmodelToModel(dm), nil
}

func DBmodelToModel(dbmodel *dbmodels.User) *models.User {
	return &models.User{
		ID:       dbmodel.ID,
		Email:    dbmodel.Email,
		Username: dbmodel.Username,
		Password: dbmodel.Password,
	}
}
