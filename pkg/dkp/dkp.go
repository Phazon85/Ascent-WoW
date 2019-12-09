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
	StartRaid(mc *discordgo.MessageCreate) error
	StopRaid(mc *discordgo.MessageCreate) error
	JoinRaid(mc *discordgo.MessageCreate) error
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

//StartRaid ...
func (s *Service) StartRaid(mc *discordgo.MessageCreate) error {
	return s.repo.StartRaid(mc)
}

//StopRaid ...
func (s *Service) StopRaid(mc *discordgo.MessageCreate) error {
	return s.repo.StopRaid(mc)
}

//JoinRaid ...
func (s *Service) JoinRaid(mc *discordgo.MessageCreate) error {
	return s.repo.JoinRaid(mc)
}
