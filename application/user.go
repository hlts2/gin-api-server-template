package application

import (
	"github.com/hlts2/gin-server-template/domain"
	"github.com/hlts2/gin-server-template/infrastructure"
	"github.com/hlts2/gin-server-template/infrastructure/persistence"
)

// GetUser returns user model object of given id
func GetUser(id string) (domain.User, error) {
	conn, err := infrastructure.NewDBConnection()
	if err != nil {
		return domain.User{}, err
	}
	defer conn.Close()

	repo := persistence.NewUserRepositoryImplWithRDB(conn)
	return repo.Get(id)
}

// GetUsers returns all user models
func GetUsers() (domain.Users, error) {
	conn, err := infrastructure.NewDBConnection()
	if err != nil {
		return domain.Users{}, err
	}

	defer conn.Close()

	rep := persistence.NewUserRepositoryImplWithRDB(conn)
	return rep.Gets()
}
