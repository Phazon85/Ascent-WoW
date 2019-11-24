package dkp

import (
	"github.com/bwmarrin/discordgo"
)

//Service holds a database service to query for results
type Service struct {
	repo Repository
}

//Repository holds functions for interacting with the repo database
type Repository interface {
	InitRaidGroup(mc *discordgo.MessageCreate) error
}

//New ...
func New(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}

//InitRaidGroup ...
func (s *Service) InitRaidGroup(mc *discordgo.MessageCreate) error {
	return s.repo.InitRaidGroup(mc)
}
