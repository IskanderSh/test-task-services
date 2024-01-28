package services

import "log/slog"

type UserService struct {
	log         *slog.Logger
	usrProvider UserProvider
}

type UserProvider interface {
	Create()
	Get()
	Update()
	Delete()
}

func New(
	log *slog.Logger,
	usrProvider UserProvider,
) *UserService {
	return &UserService{log: log, usrProvider: usrProvider}
}

func (p *UserService) Create() {

}

func (p *UserService) Get() {

}

func (p *UserService) Update() {

}

func (p *UserService) Delete() {

}
