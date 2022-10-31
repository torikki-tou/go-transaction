package service

import (
	"github.com/torikki-tou/go-transaction/dto"
	"github.com/torikki-tou/go-transaction/repo"
)

type UserService interface {
	ChangeBalance(request dto.ChangeBalance) error
}

type userService struct {
	userRepo repo.UserRepository
}

func NewUserService(userRepo repo.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

func (c *userService) ChangeBalance(request dto.ChangeBalance) error {
	_, err := c.userRepo.ChangeBalance(request.ClientID, request.Delta)
	if err != nil {
		return err
	}
	return nil
}
