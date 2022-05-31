package services

import (
	"app/repository"
	"log"
	"os"
)

type Stats struct {
	logger          *log.Logger
	statsRepository *repository.Stats
}

func NewStats() *Stats {
	logger := log.New(os.Stdout, "Stats Service: ", log.LstdFlags)
	stats := repository.NewStats()
	return &Stats{logger, stats}
}

func (service *Stats) AddStats(time int64) {

	service.statsRepository.AddStat(time)
}

func (service *Stats) GetStats() repository.StatResults {

	return service.statsRepository.GetStat()
}
