package dkp

import (
	"github.com/bwmarrin/discordgo"
)

//Service ...
type Service struct {
	repo Repository
}

//Repository ...
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
