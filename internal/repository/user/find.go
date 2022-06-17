package user

import (
	"context"
	"gorm.io/gorm"
	"investidea.tech.test/internal/entities"
	"investidea.tech.test/pkg/database"
	info_log "investidea.tech.test/pkg/info-log"
)

func (i impl) Login(ctx context.Context, email, password string) (*entities.User, error) {
	if len(password) == 0 || len(email) == 0 {
		return nil, database.InvalidRequestError
	}

	var user = &entities.User{}
	query := i.db.GormDB().
		WithContext(ctx).
		Model(user).
		Where("email = ?", email).
		Where("password = ?", password)

	err := query.First(user).Error
	if err != nil {
		i.logger.Error("find user failed", info_log.ErrorToLogFields("details", err))
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			return nil, database.NotFoundError
		}
		return nil, err
	}

	return user, nil
}
