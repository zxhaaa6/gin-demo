package user

import "github.com/zxhaaa6/gin-demo/model"

type Service struct {
}

var dao = Dao{}

func (r *Service) getAllUsers() ([]model.User, error) {
	return dao.getAllUsers()
}
