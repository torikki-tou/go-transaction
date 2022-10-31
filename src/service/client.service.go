package service

import (
	"errors"
	"github.com/torikki-tou/go-transaction/common"
	"github.com/torikki-tou/go-transaction/dto"
	"github.com/torikki-tou/go-transaction/repo"
)

type ClientService interface {
	ChangeBalance(request dto.ChangeBalance) error
}

type clientService struct {
	clientRepo repo.ClientRepository
	queueRepo  repo.QueueRepository
}

func NewClientService(clientRepo repo.ClientRepository, queueRepo repo.QueueRepository) ClientService {
	return &clientService{
		clientRepo: clientRepo,
		queueRepo:  queueRepo,
	}
}

func (c *clientService) ChangeBalance(request dto.ChangeBalance) error {
	changedBalance, err := c.clientRepo.ChangeBalance(request.ClientID, request.Delta)
	if err != nil {
		if errors.Is(err, &common.LowBalanceError{}) {
			return err
		} else {
			return &common.InternalBDError{}
		}
	}
	err = c.queueRepo.ProduceNotification(request.ClientID, request.Delta, changedBalance)
	if err != nil {
		return &common.NotificationError{}
	}
	return nil
}
