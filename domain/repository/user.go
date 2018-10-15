package repository

import "github.com/hlts2/gin-server-template/domain"

// UserRepository represetns repository interface of user
type UserRepository interface {
	Get(id string) (domain.User, error)
	Gets() (domain.Users, error)
}
