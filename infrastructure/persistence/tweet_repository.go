package persistence

import (
	"github.com/hlts2/gin-server-template/domain"
	"github.com/hlts2/gin-server-template/domain/repository"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

// userRepositoryImpl is implementation of UserRepository interfaces
type userRepositoryImpl struct {
	conn *gorm.DB
}

// NewUserRepositoryImplWithRDB returns userRepositoryImpl(repository.UserRepository) object
func NewUserRepositoryImplWithRDB(conn *gorm.DB) repository.UserRepository {
	return &userRepositoryImpl{conn: conn}
}

func (ur *userRepositoryImpl) Get(id string) (domain.User, error) {
	user := domain.User{}
	user.ID = id

	if err := ur.conn.First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return user, errors.Wrap(err, "tweet id#"+id)
		}
		return user, err
	}

	return user, nil
}

func (ur *userRepositoryImpl) Gets() (domain.Users, error) {
	users := domain.Users{}
	if err := ur.conn.Find(&users).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return users, errors.Wrap(err, "users ")
		}
		return users, err
	}

	if len(users) == 0 {
		return users, errors.Wrap(gorm.ErrRecordNotFound, "users ")
	}

	return users, nil
}
