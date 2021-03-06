package user

import "github.com/zxhaaa6/gin-demo/model"

type Service struct {
	UserDao Dao
}

func InitService() Service {
	service := Service{}
	service.UserDao = InitDao()
	return service
}

func (r *Service) findUserById(id string) (model.User, error) {
	return r.UserDao.findUserById(id)
}

func (r *Service) getAllUsers() ([]model.User, error) {
	return r.UserDao.getAllUsers()
}
